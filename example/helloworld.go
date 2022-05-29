package main

import (
	"fmt"
	"reflect"

	"github.com/kevinleestone/gotiny"
)

func main() {
	src1, src2 := "hello", []byte(" world!")
	ret1, ret2 := "", []byte{3, 4, 5}
	gotiny.Unmarshal(gotiny.Marshal(&src1, &src2), &ret1, &ret2)
	fmt.Println(ret1 + string(ret2)) // print "hello world!"

	enc := gotiny.NewEncoder(src1, src2)
	dec := gotiny.NewDecoder(ret1, ret2)

	ret1, ret2 = "", []byte{3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 5, 6, 7, 44, 7, 5, 6, 4, 7}
	dec.DecodeValue(enc.EncodeValue(reflect.ValueOf(src1), reflect.ValueOf(src2)),
		reflect.ValueOf(&ret1).Elem(), reflect.ValueOf(&ret2).Elem())
	fmt.Println(ret1 + string(ret2)) // print "hello world!"
}
