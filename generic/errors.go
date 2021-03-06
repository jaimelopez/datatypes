// Copyright (c) 2017 Jaime Lopez. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

// The datatypes/generic package includes some functionalities
// to treat in a simple way the 'generic' objects in Go

// This part of package contains the typped errors that the package uses

package generic

import "errors"

// ErrInvalidIterableElement represents an error for non-iterable elements
var ErrInvalidIterableElement = errors.New("Non-iterable type can not be converted to slice")
