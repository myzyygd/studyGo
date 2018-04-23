package main

import (
	"parseStruct/MyStruct"
	"log"
)

func main() {
	Field1 := int32(123)
	Field2 := "test"
	Field3 := []int{1, 2, 3}
	a := MyStruct.NewMyStruct(Field1, Field2, Field3)
	s:=a.Marshal()
	log.Println(s)
	a.UMarshal(s)
}
