package main

import (
	_ "pgo2-demo/pkg/command"
	_ "pgo2-demo/pkg/controller"

	"github.com/pinguo/pgo2"
)

func main() {
	pgo2.Run()
}
