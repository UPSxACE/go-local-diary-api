package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"sync"
)

func execCommand(command [2]string, waitGroup sync.WaitGroup) {
	cmd := exec.Command("cd", "../")
	cmd = exec.Command(command[0], command[1])

	var stdBuffer bytes.Buffer

	mw := io.MultiWriter(os.Stdout, &stdBuffer)
	cmd.Stdout = mw
	cmd.Stderr = mw

	// Execute the command
	if err := cmd.Run(); err != nil {
		log.Panic(err)
	}

	waitGroup.Done()
}

func main() {
	fmt.Println("Executing all 6 dev commands...")
	var wg sync.WaitGroup
	wg.Add(6)
	go execCommand([2]string{"make", "tailwind-watch"}, wg)
	go execCommand([2]string{"make", "storybook"}, wg)
	go execCommand([2]string{"make", "webpack-watch"}, wg)
	go execCommand([2]string{"air", ""}, wg)
	go execCommand([2]string{"make", "docs"}, wg)
	go execCommand([2]string{"make", "swag"}, wg)
	wg.Wait()
}
