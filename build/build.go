package build

import (
	"os"

	"github.com/mochacoder/goserverless/build/basic"
	"github.com/mochacoder/goserverless/build/golang"
	"github.com/mochacoder/goserverless/function"
)

type buildFn func(f *function.Function) (function.FilesMap, function.Config, error)

type builder struct {
	build buildFn
}

func Build(f *function.Function) (function.FilesMap, function.Config, error) {
	var b builder

	if _, err := os.Stat(f.Path + "main.go"); os.IsNotExist(err) {
		b.setBuilder(basic.Build)
	} else {
		b.setBuilder(golang.Build)
	}

	return b.build(f)
}

func (b *builder) setBuilder(bfn buildFn) {
	b.build = bfn
}
