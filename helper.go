package main

import (
	"encoding/json"
	"fmt"
	"helper/common"
	"os"
	"strings"
)

type translate struct {
	Type            string                  `json:"type"`
	ErrorCode       int                     `json:"errorCode"`
	ElapsedTime     int                     `json:"elapsedTime"`
	TranslateResult [1][1]map[string]string `json:"translateResult"`
}

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
		arg = strings.Join(args, " ")
		out, status := common.Exec_shell(arg)
		if status != nil {
			common.Invalid_arg()
		}
		countSplit := strings.Split(out, "\n")
		for _, v := range countSplit {
			if v == "" {
				continue
			}
			var result string
			fmt.Println(v)
			// if k == 30 {
			// 	break
			// }
			result = common.Gettranslate(v)
			// fmt.Println(result)
			var translate_res translate
			err := json.Unmarshal([]byte(result), &translate_res)

			if err != nil {

				panic(err)
			}
			// fmt.Println(translate_res.TranslateResult)
			for _, v := range translate_res.TranslateResult {
				fmt.Println(v[0]["tgt"])
				// fmt.Println(strings.Replace(v[0]["tgt"], "\n", "", -1))

			}
			fmt.Println("------------------------------------------")
		}
	}

}
