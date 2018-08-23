package funky

import (
	"fmt"
)

var (
	// FunctionName is the functions name set at build time.
	FunctionName = ""
)

func init() {
	fmt.Println("Func")
	fmt.Println(FunctionName)
}
