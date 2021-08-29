package main

import (
	"fmt"
	"strings"
)

func findFirstStringInBracket(str string) string {

	indexFirst := strings.Index(str, "(")
	if indexFirst>0 {
		indexLast:=strings.Index(str,")")
		if indexLast>0{
			return str[indexFirst+1 : indexLast]
		}
	}
	 
	return ""
}

func main() {
	param:="halo saya mbeng (aa gatot)"
	result:=findFirstStringInBracket(param)
	fmt.Println(result)
}
