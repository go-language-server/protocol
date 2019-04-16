// Copyright 2019 The go-language-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protocol

const (
	fileSchema = "file://"
)

// ToDocumentURI returns the new DocumentURI from s.
func ToDocumentURI(s string) DocumentURI {
	return DocumentURI(fileSchema + s)
}
