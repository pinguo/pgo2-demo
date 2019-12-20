package command

import "github.com/pinguo/pgo2"

func init()  {
	pgo2.App().Container().Bind(&TestCommand{})
}
