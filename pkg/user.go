package pkg

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func ReadInput() string {
	fmt.Print("Enter a color: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	err := scanner.Err()

	if err != nil {
		log.Fatal(err)
	}

	return scanner.Text()
}
