package main

import (
	"fmt"
	"helper/common"
	"os"
)

func main() {
	args := os.Args[1:]
	helper_help := "--help"
	if len(args) == 0 {
		common.Help()
	} else if common.Check_args(args) {
		common.Invalid_arg()
	} else if len(args) == 1 && args[0] == helper_help {
		common.Help()
	} else {
		var arg string
		for _, v := range args {
			arg += " " + v
		}
		out, status := common.Exec_shell(arg)
		if status != nil {
			common.Invalid_arg()
		}
		fmt.Println(out)

	}

}
