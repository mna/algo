//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	x, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		panic(err)
	}

	// Automatic optimization of divide-by-2 with a shift right:
	//
	// Divide by 3:
	// (div.go:31)        MOVQ    $-6148914691236517205, AX
	// (div.go:31)        IMULQ   CX
	// (div.go:31)        ADDQ    CX, DX
	// (div.go:31)        SARQ    $1, DX
	// (div.go:31)        SARQ    $63, CX
	// (div.go:31)        SUBQ    CX, DX
	//
	// Divide by 2:
	// (div.go:31)        MOVQ    AX, CX
	// (div.go:31)        SHRQ    $63, AX
	// (div.go:31)        ADDQ    CX, AX
	// (div.go:31)        SARQ    $1, AX

	y := x / 3
	fmt.Println(y)
}
