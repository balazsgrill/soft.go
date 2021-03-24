// *****************************************************************************
// Copyright(c) 2021 MASA Group
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// *****************************************************************************

package ecore

import "net/url"

type XMLResourceFactory struct {
}

func (f *XMLResourceFactory) CreateResource(uri *url.URL) EResource {
	r := newXMLResourceImpl()
	r.SetURI(uri)
	return r
}
