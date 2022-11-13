package main

import (
	"canti/app/codecs"
	"fmt"
)

func main() {
	a := "{6786578;l,78;l,67;8l,;l,5;67l,;98l,dfgsdfgsdfgsdfg}"
	b := "3c6d08d667d0ee0ccad77c55b19d3e4ab2552f7163ec40a9389095a18f86c398"

	fmt.Println(codecs.Encode(a, b))
}
