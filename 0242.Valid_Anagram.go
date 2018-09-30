package main

import (
	"fmt"
	"sort"
)

type runes []rune

func (p runes) Len() int           { return len(p) }
func (p runes) Less(i, j int) bool { return p[i] < p[j] }
func (p runes) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func isAnagram(s string, t string) bool {
	ss := runes(s)
	tt := runes(t)
	sort.Sort(ss)
	sort.Sort(tt)
	s = string(ss)
	t = string(tt)
	if s == t {
		return true
	}
	return false
}

func main() {
	s := "helloi"
	t := "ollhe"
	fmt.Println(isAnagram(s, t))
}
