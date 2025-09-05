package clip

import (
    "fmt"
    "github.com/atotto/clipboard"
)

func pollText() {
    text, err := clipboard.ReadAll()
    if err != nil {
        fmt.Printf("Error reading clipboard: %v\n", err)
        return
    }
    fmt.Println(text)
}