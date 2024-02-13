package main

import (
	"strings"
	"unicode"
)

type SeparatorStyle interface {
    TrySplit(string) (bool, []string)
    Join([]string) string
}

type SnakeSeparatorStyle struct {}
func (ss SnakeSeparatorStyle) TrySplit(src string) (bool, []string) {
    splitted := strings.Split(src, "_")
    return len(splitted) > 1, splitted
}
func (ss SnakeSeparatorStyle) Join(src []string) string {
    return strings.Join(src, "_")
}

type KebabSeparatorStyle struct {}
func (ss KebabSeparatorStyle) TrySplit(src string) (bool, []string) {
    splitted := strings.Split(src, "-")
    return len(splitted) > 1, splitted
}
func (ss KebabSeparatorStyle) Join(src []string) string {
    return strings.Join(src, "-")
}

type EmptySeparatorStyle struct {}
func (ss EmptySeparatorStyle) TrySplit(src string) (bool, []string) {
    var res = make([]string, 0)
    var buf = make([]rune, 0) 

    for i, r := range src {
        buf = append(buf, r)
        if len(src) - 1 == i || unicode.IsUpper(rune(src[i + 1])) {
            res = append(res, string(buf))
            buf = make([]rune, 0)
        }
    }

    return true, res
}
func (ss EmptySeparatorStyle) Join(src []string) string {
    return strings.Join(src, "")
}

