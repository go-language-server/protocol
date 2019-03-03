// Copyright 2019 The go-language-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protocol

import (
	"log"
)

// unhandledError is used in places where an error may occur that cannot be handled.
// This occurs in things like rpc handlers that are a notify, where we cannot
// reply to the caller, or in a call when we are actually attempting to reply.
// In these cases, there is nothing we can do with the error except log it, so
// we do that in this function, and the presence of this function acts as a
// useful reminder of why we are effectively dropping the error and also a
// good place to hook in when debugging those kinds of errors.
func unhandledError(err error) {
	if err == nil {
		return
	}
	log.Printf("%v", err)
}
