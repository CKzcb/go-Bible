/**
 * @Author: mrProtein
 * @Description:
 * @File:  defineInit
 * @Version: 1.0.0
 * @Date: 2023/08/14 21:41
 */
package DeclarationsAndDefine

import "fmt"

var s1 string
var s2 string = "init"

var i1 int
var i2 int = 30

func DefineInitTest1() {
	fmt.Println("this is s1 ... ", s1)
	fmt.Println("this is s2 ... ", s2)
	fmt.Println("this is i1 ... ", i1)
	fmt.Println("this is i2 ... ", i2)
}
