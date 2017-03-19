# datatypes
GoLang package which provides new iterable structs to easy managing and extends some other basic functionalities.

## collection
Helps to the iteration of non-sorted unique element and homogeneous lists 

Examples:
```
import "github.com/jaimelopez/datatypes/collection"

elementOne := "first element"
elementTwo := "second element"
elements := ElementList{elementOne, elementTwo}

col := NewCollection(elements)

for !col.IsEmpty() {
    element = col.Extract()

    fmt.PrintLn(element)
}
```

In this case Extract() method will extract and return the first element of the collection. If you want you can iterate over the elements with Elements() method which just return all the elements without removing these elements.

Other example:
```
import "github.com/jaimelopez/datatypes/collection"

elementOne := "first element"
elementTwo := "second element"

col := NewEmptyCollection()

if !col.Contains(elementOne) {
    col.Add(elementOne)
}

if !col.Contains(elementTwo) {
    col.Add(elementTwo)
}

for _, elem := range col.Elements() {
    fmt.PrintLn(elem)
}
```

\* [See test for further more information about how play with it.](/collection/collection_test.go)


## dictionary
Provides an easy dictionary (key => value) struct management
```
import "github.com/jaimelopez/datatypes/dictionary"
```

\* [See test for further more information about how play with it.](/dictionary/dictionary_test.go)