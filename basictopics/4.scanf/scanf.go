package scanf

import (
	"bufio"
	"fmt"
	"os"
)

func Scanf() {
	fmt.Println("Hey!, What's your name")
	// var name string
	// fmt.Scan(&name) // it take only a word if there is space it ignore after all ther words after first space
	// fmt.Println("Hello", name)

	// We use Reader to take input or taking pragrapgh from user with space
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString(
		'\n',
	)
	fmt.Println("Hello", name)

}
