package solution

import (
	"fmt"
	"github.com/kyokomi/emoji/v2"
)

func solution GetMessage(Message) string {

	Message := emoji.Sprint("Hello :world_map:!")
	fmt.Println("string:", Message)

	return "Message"
}
