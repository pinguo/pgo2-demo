package command

import (
	"flag"
	"fmt"

	"github.com/pinguo/pgo2"
)

type TestCommand struct {
	pgo2.Controller
}

// pgo2-demo --env=docker --cmd=test/index --name=cutomeName
func (t *TestCommand) ActionIndex() {
	name := flag.String("name", "", " --name=xxxx")
	flag.Parse()
	args := flag.Args()
	fmt.Println("call in command/Test.Index, name:", *name, "args:", args)
}
