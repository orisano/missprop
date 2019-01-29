package main

import (
	"github.com/orisano/missprop"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(missprop.Analyzer)
}
