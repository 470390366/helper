package common

import "fmt"

func Help() {
	fmt.Println(`
Usage:
	helper <command> 
	
	`)
}

func Invalid_arg() {
	fmt.Println()
	fmt.Println("invalid agr please input right agrs .")
	fmt.Println("-------------------------------------")
	Help()
}

func Check_args(s []string) bool {
	var check bool
	check = true
	check_str := "--help"
	for _, v := range s {
		if v == check_str {
			check = false
			break
		}
	}
	return check
}
