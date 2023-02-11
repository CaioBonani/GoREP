package main

import (
	
	"example/hello/lib"
)

func main() {
	texto := lib.Run()
	lib.Search(texto, lib.Substring)
}
