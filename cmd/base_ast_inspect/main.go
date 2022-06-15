package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"
	"golint-extend/pkg/analyzer_inspect"
)

func main() {
	singlechecker.Main(analyzer_inspect.Analyzer)
}
