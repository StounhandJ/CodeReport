package utils

import (
	"bufio"
	"fmt"
	"os"
)

func WaitForExit() {
	fmt.Println()
	fmt.Print("Press Enter to close...")
	_, _ = bufio.NewReader(os.Stdin).ReadString('\n')
}
