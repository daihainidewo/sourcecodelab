// Package strings strings
// file create by daihao, time is 2018/7/27 15:19
package main

import (
	"fmt"
	"strings"
)

// main main
func main() {
	s := "a bcd Ef    g h jkl op qrst UVw xyz "
	fmt.Println(strings.Contains(s, "abc"))
	fmt.Println(strings.ContainsAny(s, "abc"))
	fmt.Println(strings.ContainsRune(s, ' '))
	fmt.Println(strings.LastIndex(s, "u"))
	fmt.Println(strings.IndexRune(s, ' '))
	fmt.Println(strings.IndexAny(s, "gh"))
	fmt.Println(strings.LastIndexAny(s, " "))
	fmt.Println(strings.LastIndexByte(s, 'y'))
	fmt.Println(strings.SplitN(s, " ", 2))
	fmt.Println(strings.SplitAfterN(s, " ", 2))
	fmt.Println(strings.Split(s, " "))
	fmt.Println(strings.Fields(s))
	fmt.Println(strings.FieldsFunc(s, func(r rune) bool {
		if r == 'h' || r == 'o' {
			return true
		}
		return false
	}))
	a := []string{"123", "abc", "!@#"}
	fmt.Println(strings.Join(a, "|"))
	fmt.Println(strings.HasPrefix(s, "a "))
	fmt.Println(strings.HasSuffix(s, "xyz"))
	fmt.Println(strings.Map(func(r rune) rune {
		if r == ' ' {
			return '|'
		}
		return r
	}, s))
	fmt.Println(strings.Repeat(s, 2))
	fmt.Println(strings.ToUpper(s))
	fmt.Println(strings.ToLower(s))
	fmt.Println(strings.ToTitle(s))
	fmt.Println(strings.Title(s))
	fmt.Println(strings.IndexFunc(s, func(r rune) bool {
		if r == ' ' {
			return true
		}
		return false
	}))
	fmt.Println(strings.LastIndexFunc(s, func(r rune) bool {
		if r == 'a' || r == 'g' {
			return true
		}
		return false
	}))
	fmt.Println(strings.Trim(s, "axyz "))
	fmt.Println(strings.TrimLeft(s, "a"))
	fmt.Println(strings.TrimRight(s, " "))
	fmt.Println(strings.TrimSpace(s))
	fmt.Println(strings.TrimPrefix(s, "a bc"))
	fmt.Println(strings.TrimSuffix(s, "xyz"))
	fmt.Println(strings.Replace(s, " ", "|", 4))
	fmt.Println(strings.EqualFold("Abc", "aBC"))
}
