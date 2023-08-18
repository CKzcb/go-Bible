/**
 * @Author: mrProtein
 * @Description:
 * @File:  compoundType
 * @Version: 1.0.0
 * @Date: 2023/08/18 7:43
 */
package structType

import "fmt"

func TestSlice() {
	var a = [3]int{}
	//var b = [...]int{}
	a[0] = 1
	a[1] = 2
	a[2] = 3
	var c []string
	c = append(c, "nihao")
	fmt.Println(&c, c[0])
	fmt.Println(a, c)
}
