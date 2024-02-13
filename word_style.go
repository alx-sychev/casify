package main

import "strings"

type WordStyle interface {
    Style([]string) []string
}

type LowerWordStyle struct {}
func (ws LowerWordStyle) Style(src []string) []string {
    result := make([]string, len(src))
    for i, word := range src {
        result[i] = strings.ToLower(word)
    }
    return result
}

type CamelWordStyle struct {}
func (ws CamelWordStyle) Style(src []string) []string {
    if len(src) == 0 { return src }
    result := make([]string, len(src))
    result[0] = strings.ToLower(src[0])
    for i := 1; i < len(src); i++ {
        result[i] = strings.Title(src[i])
    }
    return result
}

type PascalWordStyle struct {}
func (ws PascalWordStyle) Style(src []string) []string {
    result := make([]string, len(src))
    for i, word := range src {
        result[i] = strings.Title(word)
    }
    return result
}

type UpperWordStyle struct {}
func (ws UpperWordStyle) Style(src []string) []string {
    result := make([]string, len(src))
    for i, word := range src {
        result[i] = strings.ToUpper(word)
    }
    return result
}

