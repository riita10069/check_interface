package main

import (
	"github.com/riita10069/check_interface"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(check_interface.Analyzer) }

