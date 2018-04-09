package main

// import (
// 	"fmt"
// )

// func main() {
// 	var1 := "Peter"
// 	fmt.Printf("Do more with less, %v\n", var1)
// }

import (
	"findlinks/server"
	"fmt"

	"github.com/golang/example/stringutil"
)

func main() {
	fmt.Println(stringutil.Reverse("!selpmaxe oG ,olleH"))
	server.CopyFile("./README.md2", "./README.md")
	a()
}

func a() {
	i := 0
	defer fmt.Println("- %d", i)
	i++
	fmt.Println(": %d", i)
	return
}
