/**
 * @Author: mrProtein
 * @Description:
 * @File:  intType
 * @Version: 1.0.0
 * @Date: 2023/08/16 7:41
 */
package structType

import (
	"fmt"
	"reflect"
)

func TestInt8() {
	var a = 0
	for i := 0; i < 260; i++ {
		a++
		fmt.Println(a, reflect.TypeOf(a))
	}
}

func TestIntLength() {
	var i1 uint8 = 0
	var i2 uint16 = 0
	var i3 uint32 = 0
	var i4 uint64 = 0
	var i5 uint = 0

	fmt.Println((^i1) >> 1)
	fmt.Println((^i2) >> 1)
	fmt.Println((^i3) >> 1)
	fmt.Println((^i4) >> 1)
	fmt.Println((^i5) >> 1)
}
