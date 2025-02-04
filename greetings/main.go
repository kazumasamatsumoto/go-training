package main

import (
    "fmt"
    "example.com/greetings/greetings" // サブディレクトリ内の greetings パッケージをインポート
)

func main() {
    message := greetings.Hello("World")
    fmt.Println(message)
}
