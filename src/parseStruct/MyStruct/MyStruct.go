package MyStruct

import (
	"encoding/binary"
	"log"
)

type MyStruct struct {
	Field1 int32
	Field2 string
	Field3 []int
}

func (s *MyStruct) BirySize() int {
	n := 4 + 2 + len(s.Field2) + 2 + 2*len(s.Field3)
	return n
}

func NewMyStruct(Field1 int32, Field2 string, Field3 []int) *MyStruct {
	s := &MyStruct{
		Field1: Field1,
		Field2: Field2,
		Field3: Field3,
	}
	return s
}
//打包
func (s *MyStruct) Marshal() []byte {
	b := make([]byte, s.BirySize())
	n := 0
	binary.BigEndian.PutUint32(b[n:], uint32(s.Field1))
	n += 4
	binary.BigEndian.PutUint16(b[n:], uint16(len(s.Field2)))
	n += 2
	copy(b[n:], s.Field2)
	n += len(s.Field2)
	binary.BigEndian.PutUint16(b[n:], uint16(len(s.Field3)))
	n += 2
	for i := 0; i < len(s.Field3); i++ {
		binary.BigEndian.PutUint16(b[n:], uint16(s.Field3[i]))
		n += 2
	}
	return b
}
//解包
func (s *MyStruct) UMarshal(b []byte) {
	n := 0
	s.Field1 = int32(binary.BigEndian.Uint32(b[n:]))
	n += 4
	x := int(binary.BigEndian.Uint16(b[n:]))
	n += 2
	Field2 := string(b[n:n+x])
	log.Println(Field2)
	s.Field2 = Field2
	n += x
	y := binary.BigEndian.Uint16(b[n:])
	n += 2
	log.Println(y)
	Field3 := make([]int, y)
	for i := 0; i < len(s.Field3); i++ {
		Field3[i] = int(binary.BigEndian.Uint16(b[n:]))
		n += 2
	}
	log.Println(Field3)
}
