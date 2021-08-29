package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var dummy []string=[]string{
		"kita", "atik", "tika", "aku", "kia", "makan", "kua",
	}
	// fmt.Println(dummy)
	listAnagram := make(map[string][]string)

	for _, word := range dummy {
		key := sortString(word)
		// fmt.Println(key)
		listAnagram[key] = append(listAnagram[key], word)
	}
	// fmt.Println(len(listAnagram))
	for _, words := range listAnagram {
		fmt.Println(words)
	}
}

func sortString(param string) string {
	s := strings.Split(param, "")
	sort.Strings(s)

	return strings.Join(s, "")
}