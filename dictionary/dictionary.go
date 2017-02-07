package dictionary

//import "errors"
//
//type KeyElement interface{}
//type ValueElement interface{}
//type KeyRange interface{}
//
//type KeyValues map[KeyElement]ValueElement

//type Dictionary {
//     elements KeyValues
// }
//
//func (dicc Dictionary) Add(key KeyElement, value ValueElement) error {
//	if dicc.Contains(element) {
//		return NewDuplicatedKeyError()
//	}
//
//	dicc = append(dicc, element)
//
//	return nil
//}
//
//func (dicc *Dictionary) AddRange(elements []interface{}) error {
//	return nil
//}
//
//func (dicc *Dictionary) Delete(element interface{}) error {
//	if !dicc.Contains(element) {
//		return errors.New("Element not found")
//	}
//
//	return nil
//}
//
//func (dicc *Dictionary) DeleteRange(elements KeyRange) error {
//	return nil
//}
//
//func (dicc *Dictionary) Contains(element interface{}) bool {
//	for _, iterator := range dicc {
//		if iterator == element {
//			return true
//		}
//	}
//
//	return false
//}
//
//func (dicc *Dictionary) ContainsAny([]interface{}) bool {
//	return false
//}
//
//func (dicc *Dictionary) ContainsValue(element ValueElement) bool {
//	return false
//}
//
//func (dicc *Dictionary) Elements() {}
//
//func (dicc *Dictionary) GetKeys() {}
//
//func (dicc *Dictionary) GetValues() {}
//
//func (dicc *Dictionary) Count() int {
//	return len(dicc)
//}
//
//func NewEmptyDictionary() *Dictionary {
//	return new(Dictionary)
//}
//
// func NewDictionary(elements KeyValues) *Dictionary {
// 	dictionary := &Dictionary(elements)
// 	error := dictionary.AddRange(generic.ToSlice(elements))
//
// 	if error != nil {
// 		panic(error)
// 	}
//
// 	return collection
// }
