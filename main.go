package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func ReadFileContents(path string) string {
    data, err := os.ReadFile(path)
    if err != nil {
        log.Fatal(err)
    }
    return string(data)
}

func Format(contents string, characterLimit int) string {
    contents = strings.Trim(contents, " ")
    lines := strings.Split(contents, "\n")
    for i, line := range lines {
        if len(line) > characterLimit {
            for j := characterLimit; j > 0; j-- {
                if unicode.IsSpace(rune(line[j])) {
                    lines[i] = line[:j] + "\n" + line[j+1:] 
                    break
                }
            }
        }
    }
   return strings.Join(lines, "\n")
}

func WriteContents(path, contents string) {
    err := os.WriteFile(path, []byte(contents), 0664)
    if err != nil {
        fmt.Println(err)
    }
}

func main() {
    file := flag.String("file", "", "file that we will run formatting on")
    output := flag.String("output", "", "name of file to output formatted content")
    charLimit := flag.Int("char-limit", 80, "lengths limit of each line")
    flag.Parse()
    contents := ReadFileContents(*file)
    contents = Format(contents, *charLimit)
    if *output == "" {
        WriteContents(*file, contents)
    } else {
        WriteContents(*output, contents)
    }
}
