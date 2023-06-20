package ecore

import (
	"database/sql"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
)

type sqlDecoderClassData struct {
	schema    *sqlClassSchema
	eClass    EClass
	eFactory  EFactory
	features  map[EStructuralFeature]*sqlDecoderFeatureData
	hierarchy []EClass
}

type sqlDecoderFeatureData struct {
	schema    *sqlFeatureSchema
	eFeature  EStructuralFeature
	eFactory  EFactory
	eDataType EDataType
}

type SQLDecoder struct {
	resource    EResource
	reader      io.Reader
	driver      string
	db          *sql.DB
	schema      *sqlSchema
	classesData map[int]*sqlDecoderClassData
}

func NewSQLDecoder(resource EResource, r io.Reader, options map[string]any) *SQLDecoder {
	// options
	schemaOptions := []sqlSchemaOption{}
	driver := "sqlite"
	if options != nil {
		if driver, isDriver := options[SQL_OPTION_DRIVER]; isDriver {
			driver = driver.(string)
		}

		idAttributeName, _ := options[JSON_OPTION_ID_ATTRIBUTE_NAME].(string)
		if resource.GetObjectIDManager() != nil && len(idAttributeName) > 0 {
			schemaOptions = append(schemaOptions, withIDAttributeName(idAttributeName))
		}
	}

	return &SQLDecoder{
		resource:    resource,
		reader:      r,
		driver:      driver,
		schema:      newSqlSchema(schemaOptions...),
		classesData: map[int]*sqlDecoderClassData{},
	}
}

func (d *SQLDecoder) createDB() (*sql.DB, error) {
	fileName := filepath.Base(d.resource.GetURI().Path())
	dbPath, err := sqlTmpDB(fileName)
	if err != nil {
		return nil, err
	}

	dbFile, err := os.Create(dbPath)
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(dbFile, d.reader)
	if err != nil {
		dbFile.Close()
		return nil, err
	}
	dbFile.Close()

	return sql.Open(d.driver, dbPath)
}

func (d *SQLDecoder) DecodeResource() {
	var err error
	d.db, err = d.createDB()
	if err != nil {
		d.addError(err)
		return
	}

	if err := d.decodeVersion(); err != nil {
		d.addError(err)
		return
	}

	if err := d.decodeClasses(); err != nil {
		d.addError(err)
		return
	}

	if err := d.decodeContents(); err != nil {
		d.addError(err)
		return
	}
}

func (d *SQLDecoder) DecodeObject() (EObject, error) {
	panic("SQLDecoder doesn't support object decoding")
}

func (d *SQLDecoder) addError(err error) {
	d.resource.GetErrors().Add(NewEDiagnosticImpl(err.Error(), d.resource.GetURI().String(), 0, 0))
}

func (d *SQLDecoder) decodeVersion() error {
	if row := d.db.QueryRow("PRAGMA user_version;"); row == nil {
		return fmt.Errorf("unable to retrieve user version")
	} else {
		var v int
		if err := row.Scan(&v); err != nil {
			return err
		}
		if v != sqlCodecVersion {
			return fmt.Errorf("history version %v is not supported", v)
		}
		return nil
	}
}

func (d *SQLDecoder) decodeContents() error {
	return nil
}

func (d *SQLDecoder) decodeObject() (EObject, error) {
	return nil, nil
}

func (d *SQLDecoder) decodeClasses() error {
	// read packages
	packagesData, err := d.decodePackages()
	if err != nil {
		return err
	}

	// read classes
	rows, err := d.db.Query(d.schema.classesTable.selectAllQuery())
	if err != nil {
		return err
	}
	defer rows.Close()

	classesData := map[int]*sqlDecoderClassData{}
	rawBuffer := make([]sql.RawBytes, 3)
	scanCallArgs := make([]any, len(rawBuffer))
	for i := range rawBuffer {
		scanCallArgs[i] = &rawBuffer[i]
	}

	for rows.Next() {
		// retrieve EClass
		if err := rows.Scan(scanCallArgs...); err != nil {
			return err
		}
		classID, _ := strconv.Atoi(string(rawBuffer[0]))
		packageID, _ := strconv.Atoi(string(rawBuffer[1]))
		className := string(rawBuffer[2])
		ePackage, _ := packagesData[packageID]
		if ePackage == nil {
			return fmt.Errorf("unable to find package with id '%d'", packageID)
		}
		eClass, _ := ePackage.GetEClassifier(className).(EClass)
		if eClass == nil {
			return fmt.Errorf("unable to find class '%s' in package '%s'", className, ePackage.GetNsURI())
		}

		// get class schema
		classSchema, err := d.schema.getClassSchema(eClass)
		if err != nil {
			return err
		}

		// compute class hierarchy
		classHierarchy := []EClass{eClass}
		for itClass := eClass.GetEAllSuperTypes().Iterator(); itClass.HasNext(); {
			classHierarchy = append(classHierarchy, itClass.Next().(EClass))
		}

		// compute class features
		classFeatures := map[EStructuralFeature]*sqlDecoderFeatureData{}
		for eFeature, featureSchema := range classSchema.features {
			eFeatureData := &sqlDecoderFeatureData{
				eFeature: eFeature,
				schema:   featureSchema,
			}
			if eAttribute, _ := eFeature.(EAttribute); eAttribute != nil {
				eFeatureData.eDataType = eAttribute.GetEAttributeType()
				eFeatureData.eFactory = eFeatureData.eDataType.GetEPackage().GetEFactoryInstance()
			}
			classFeatures[eFeature] = eFeatureData
		}

		// register class data
		classesData[classID] = &sqlDecoderClassData{
			eClass:    eClass,
			eFactory:  ePackage.GetEFactoryInstance(),
			schema:    classSchema,
			hierarchy: classHierarchy,
			features:  classFeatures,
		}

	}
	if err = rows.Err(); err != nil {
		return err
	}

	d.classesData = classesData
	return nil
}

func (d *SQLDecoder) decodePackages() (map[int]EPackage, error) {
	rows, err := d.db.Query(d.schema.packagesTable.selectAllQuery())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	packagesData := map[int]EPackage{}
	rawBuffer := make([]sql.RawBytes, 2)
	scanCallArgs := make([]any, len(rawBuffer))
	for i := range rawBuffer {
		scanCallArgs[i] = &rawBuffer[i]
	}

	for rows.Next() {
		if err := rows.Scan(scanCallArgs...); err != nil {
			return nil, err
		}
		packageID, _ := strconv.Atoi(string(rawBuffer[0]))
		packageURI := string(rawBuffer[1])
		packageRegistry := GetPackageRegistry()
		resourceSet := d.resource.GetResourceSet()
		if resourceSet != nil {
			packageRegistry = resourceSet.GetPackageRegistry()
		}
		ePackage := packageRegistry.GetPackage(packageURI)
		if ePackage == nil {
			return nil, fmt.Errorf("unable to find package '%s'", packageURI)
		}
		packagesData[packageID] = ePackage
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return packagesData, nil
}
