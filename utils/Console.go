package utils

import (
	"bufio"
	"fmt"
	"os"
)

func WaitForExit() {
	fmt.Println()
	fmt.Print("Press Enter to close...")

	_, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		panic(err)
	}
}
