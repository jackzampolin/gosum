package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	h := sha256.New()
	if _, err := io.Copy(h, reader); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%x  -\n", h.Sum(nil))
}
