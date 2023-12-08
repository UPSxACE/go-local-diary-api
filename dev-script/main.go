package main

import (
	"fmt"
	"log"
	"os/exec"
	"sync"
)

func execCommand(command [2]string, waitGroup sync.WaitGroup) {
	cmd := exec.Command("cd" , "../")
	cmd = exec.Command(command[0], command[1])

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error on command:", command[0], command[1])
		log.Fatal(err)
	}
	
	waitGroup.Done()
}

func main() {
	fmt.Println("Executing all 6 dev commands...")
	var wg sync.WaitGroup
	wg.Add(6)
	go execCommand([2]string{"make", "tailwind-watch"},wg)
	go execCommand([2]string{"make", "storybook"},wg)
	go execCommand([2]string{"make", "webpack-watch"},wg)
	go execCommand([2]string{"air", ""},wg)
	go execCommand([2]string{"make", "docs"},wg)
	go execCommand([2]string{"make", "swag"},wg)
	wg.Wait()
}