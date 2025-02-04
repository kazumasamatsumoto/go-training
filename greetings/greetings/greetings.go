package greetings

import "fmt"

// Hello は、名前を受け取って挨拶のメッセージを返します。
func Hello(name string) string {
	// 名前をメッセージに埋め込み、挨拶のメッセージを生成します。
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
