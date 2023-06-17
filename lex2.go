package main

import (
    "bufio"
    "fmt"
    "os"
    "unicode"
)

type Tokenisasi struct {
    baca string
    line string
    data *os.File
}

func (t *Tokenisasi) OpenFile() {
    file, err := os.Open("CFG.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        t.line = scanner.Text()
        t.baca += t.line + "\n"
    }

    if err := scanner.Err(); err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println(t.baca)
}

func (t *Tokenisasi) Token() {
    var spasi int
    for _, char := range t.baca {
        if unicode.IsSpace(char) {
            spasi++
            fmt.Println()
            fmt.Print("\t")
        }
        fmt.Print(string(char))
    }
    fmt.Scanln()
}

func main() {
    T := Tokenisasi{}
    fmt.Println("++++ISI FILE ++++\n\n")
    T.OpenFile()
    fmt.Println("++++HASIL TOKEN ++++\n\n")
    T.Token()
}
