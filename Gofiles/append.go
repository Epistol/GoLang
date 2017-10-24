package Gofiles

import (
	"fmt"
	/*"math"
	"strings"*/
)


-func Concat(integers []int){
	integers = append(integers,1)
	fmt.Println("In concact : ", integers)
}

func ConcatByAddr(integers *[]int){
	*integers = append(*integers,1)
}

func ConcatbyVal(integers []int) []int{
	return append(integers,1)
}

func main() {

	slice0 := []int{-1,0}
	Concat(slice0)
	ConcatByAddr(&slice0)
	fmt.Println("After : ", slice0)
	slice0 = ConcatbyVal(slice0)
	fmt.Println("After : ", slice0)

}