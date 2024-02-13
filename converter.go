package main

import "fmt"

type Converter struct {
    ws WordStyle
    ss SeparatorStyle
}

func (c *Converter) Convert(src string) (string, error) {
    splitted, err := c.trySplit(src)
    if err != nil {
        return "", err
    }
    normalized := LowerWordStyle{}.Style(splitted)
    return c.ss.Join(c.ws.Style(normalized)), nil 
}

func (c *Converter) trySplit(src string) ([]string, error) {
    var ok bool
    var splitted []string

    ok, splitted = SnakeSeparatorStyle{}.TrySplit(src)
    if ok { return splitted, nil }

    ok, splitted = KebabSeparatorStyle{}.TrySplit(src)
    if ok { return splitted, nil }

    ok, splitted = EmptySeparatorStyle{}.TrySplit(src)
    if ok { return splitted, nil }

    return nil, fmt.Errorf("unable to detect separator style for word: %s", src) 
}

