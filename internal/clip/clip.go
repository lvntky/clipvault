package clip

import (
       "fmt"
       "github.com/atotto/clipboard"
)

func pollText() {

     text, _ := clipboard.readAll()
     fmt.Println(text);
}