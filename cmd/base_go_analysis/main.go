package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"
	"golint-extend/pkg/analyzer"
)

func main() {
	singlechecker.Main(analyzer.Analyzer)
}
