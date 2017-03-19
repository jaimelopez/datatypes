// Package datatypes is a set of packages that provides many tools which helps you treating some complex structures and primitives types.
//
// The following packages are distributed under datatypes:
//  @datatypes/collection : Provides a new single-value list struct simplifying the iteration over it
//  @datatypes/dictionary : Provides a new object type which allows to treat them like <Key,Value> list
//  @datatypes/generic : Packages which includes some functionalities to treat 'generic' objects
//  @datatypes/string : Simple package which encapsulates some basic operations over strings
package datatypes

// blank imports help docs.
import (
	// collection package
	_ "github.com/jaimelopez/datatypes/collection"
	// dictionary package
	_ "github.com/jaimelopez/datatypes/dictionary"
	// generic package
	_ "github.com/jaimelopez/datatypes/generic"
	// string package
	_ "github.com/jaimelopez/datatypes/string"
)
