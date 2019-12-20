package command

import (
	"flag"
	"fmt"

	"github.com/pinguo/pgo2"
)

type TestCommand struct {
	pgo2.Controller
}

func (t *TestCommand) ActionIndex() {
	name := flag.String("name", "", " --name=xxxx")
	flag.Parse()
	args := flag.Args()
	fmt.Println("call in command/Test.Index, name:", name, "args:", args)
}
