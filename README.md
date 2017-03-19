# datatypes
GoLang package which provides new iterable structs to easy managing and extends some other basic functionalities.

## collection
Helps to the iteration of non-sorted unique element and homogeneous lists 

Example:

```go
import "github.com/jaimelopez/datatypes/collection"

elementOne := "first element"
elementTwo := "second element"
elements := ElementList{elementOne, elementTwo}

col := NewCollection(elements)

for !col.IsEmpty() {
    element = col.Extract()

    fmt.Println(element)
}
```

In this case Extract() method will extract and return the first element of the collection. If you want you can iterate over the elements with Elements() method which just return all the elements without removing these elements.

Other example:

```go
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
    fmt.Println(elem)
}
```

\* [See tests for further more information about how play with it.](/collection/collection_test.go)


## dictionary
Provides an easy dictionary (key => value) struct management

This example fill the dictionary with two elements and then retrieves the first element:

```go
import "github.com/jaimelopez/datatypes/dictionary"

elementOne := KeyValueElement{"1Key", "1Value"}
elementTwo := KeyValueElement{"2Key", "2Value"}

dic := NewDictionary(KeyValueList{elementOne, elementTwo})

first := dic.First()

fmt.Println(first)
```

In this example we will see how iterate over dictionaries:

```go
import "github.com/jaimelopez/datatypes/dictionary"

elementOne := KeyValueElement{"1Key", "1Value"}
elementTwo := KeyValueElement{"2Key", "2Value"}

dic := NewEmptyDictionary()
dic.AddKeyValueElement(elementOne)
dic.AddKeyValueElement(elementTwo)

for key, value := range dic.Elements() {
    fmt.Println(key)
    fmt.Println(value)
}
```

Getting all stored keys:

```go
import "github.com/jaimelopez/datatypes/dictionary"

elementOne := KeyValueElement{"1Key", "1Value"}
elementTwo := KeyValueElement{"2Key", "2Value"}
elementList := KeyValueList{elementOne, elementTwo}

dic := NewEmptyDictionary()
dic.AddRange(elementList)

allKeys := dic.Keys()

fmt.Println(allKeys)

```

\* [See tests for further more information about how play with it.](/dictionary/dictionary_test.go)