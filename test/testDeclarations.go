/**
 * @Author: mrProtein
 * @Description:
 * @File:  testDeclarations
 * @Version: 1.0.0
 * @Date: 2023/08/14 21:37
 */
package main

import (
	"fmt"
	dad "github.com/ckzcb/go-Bible/DeclarationsAndDefine"
	"reflect"
)

func fToc(f float64) float64 {
	return (f - 32) * 5 / 9
}

func main() {
	//const freezingF, boilingF = 32.0, 212.0
	//fmt.Printf("%g°F = %g°C\n", freezingF, fToc(freezingF))
	//fmt.Printf("%g°F = %g°C\n", boilingF, fToc(boilingF))

	dad.DefineInitTest1()

	a := 200
	fmt.Println(reflect.TypeOf(a));

	var b rune
	b = '好'
	fmt.Println(reflect.TypeOf(b));
}
