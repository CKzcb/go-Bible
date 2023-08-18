/**
 * @Author: 金金
 * @Description:
 * @File:  testCommandParams
 * @Version: 1.0.0
 * @Date: 2023/08/14 21:30
 */
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)
	for i := 0; i < len(os.Args); i++ {
		fmt.Println(os.Args[i])
	}
}
