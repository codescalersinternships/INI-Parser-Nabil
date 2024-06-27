package main

import (
	"fmt"

	iniparser "github.com/codescalersinternships/INI-Parser-Nabil/pkg"
)

func main() {
	parser := iniparser.NewIniParser()
	parser.LoadFromFile("./input.ini")
	fmt.Println(parser.SaveToFile("Hello.ini"))
}
