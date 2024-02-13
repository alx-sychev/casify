package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func findWordStyle(name string) (bool, WordStyle) {
    if name == "lower" { return true, LowerWordStyle{} }
    if name == "upper" { return true, UpperWordStyle{} }
    if name == "camel" { return true, CamelWordStyle{} }
    if name == "pascal" { return true, PascalWordStyle{} }

    return false, nil
}

func findSeparatorStyle(name string) (bool, SeparatorStyle) {
    if name == "snake" { return true, SnakeSeparatorStyle{} }
    if name == "kebab" { return true, KebabSeparatorStyle{} }
    if name == "empty" { return true, EmptySeparatorStyle{} }

    return false, nil
}

func createConverter(ssName, wsName string) (*Converter, error) {
    ok, ws := findWordStyle(wsName)
    if !ok {
        return nil, fmt.Errorf("unknown word style: %s", wsName)
    }

    ok, ss := findSeparatorStyle(ssName)
    if !ok {
        return nil, fmt.Errorf("unknown separator style: %s", ssName)
    }

    return &Converter{
        ws: ws,
        ss: ss,
    }, nil
}

func main() {
    log.SetFlags(0)
    args := os.Args

    if (len(args) != 3) {
        log.Fatal("error: invalid arguments count")
    }

    converter, err := createConverter(args[1], args[2])
    if err != nil {
        log.Fatalf("error: %v", err)
    }

    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanWords)
    for scanner.Scan() {
        input := scanner.Text()

        result, err := converter.Convert(input)
        if err != nil {
            log.Fatalf("error: %v", err)
        }

        fmt.Println(result)
    }

    if err := scanner.Err(); err != nil {
        log.Fatalf("error: %v", err)
    }
}

