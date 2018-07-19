 package main

 import (
 	"fmt"
 	"reflect"
 )

 type A struct {
 	A  fmt.ScanState "Tag A"                  // index 0
 	AB B             "Tag A of type B struct" // index 1
 	B  "Tag B"       // index 2 -embedded field
 }
 type B struct {
 	B0 string
 	B1 int
 }

 func main() {
 	var a A
 	var typeof reflect.Type = reflect.TypeOf(a)

 	var field reflect.StructField = typeof.Field(0) // index 0

 	fmt.Println("Name : ", field.Name)
 	fmt.Println("PkgPath : ", field.PkgPath) //<--- only applicable if the field name is small cap(unexported)
 	fmt.Println("Type : ", field.Type)
 	fmt.Println("Tag : ", field.Tag)
 	fmt.Println("Offset : ", field.Offset)
 	fmt.Println("Index : ", field.Index)
 	fmt.Println("Anonymous : ", field.Anonymous)

 	fmt.Println("// --------------------------------------------------")

 	field = typeof.Field(1) // index 1

 	fmt.Println("Name : ", field.Name)
 	fmt.Println("PkgPath : ", field.PkgPath) //<--- only applicable if the field name is small cap(unexported)
 	fmt.Println("Type : ", field.Type)
 	fmt.Println("Tag : ", field.Tag)
 	fmt.Println("Offset : ", field.Offset)
 	fmt.Println("Index : ", field.Index)
 	fmt.Println("Anonymous : ", field.Anonymous)

 	fmt.Println("// --------------------------------------------------")
 	field = typeof.Field(2) // index 2

 	fmt.Println("Name : ", field.Name)
 	fmt.Println("PkgPath : ", field.PkgPath) //<--- only applicable if the field name is small cap(unexported)
 	fmt.Println("Type : ", field.Type)
 	fmt.Println("Tag : ", field.Tag)
 	fmt.Println("Offset : ", field.Offset)
 	fmt.Println("Index : ", field.Index)
 	fmt.Println("Anonymous : ", field.Anonymous) // TRUE because it is embedded field

}