package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

// Функція для очищення консолі
func clearScreen() {
	cmd := exec.Command("clear")  // для UNIX/Linux/Mac
	if os.PathSeparator == '\\' { // Windows
		cmd = exec.Command("cmd", "/c", "cls")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// Функція для відображення світлофора
func displayLight(color string) {
	clearScreen()

	if color == "red" {
		fmt.Println("\x1b[41m     \x1b[0m\n\x1b[41m     \x1b[0m")
	} else {
		fmt.Println("\x1b[40m     \x1b[0m\n\x1b[40m     \x1b[0m")
	}

	if color == "yellow" {
		fmt.Println("\x1b[43m     \x1b[0m\n\x1b[43m     \x1b[0m")
	} else {
		fmt.Println("\x1b[40m     \x1b[0m\n\x1b[40m     \x1b[0m")
	}

	if color == "green" {
		fmt.Println("\x1b[42m     \x1b[0m\n\x1b[42m     \x1b[0m")
	} else {
		fmt.Println("\x1b[40m     \x1b[0m\n\x1b[40m     \x1b[0m")
	}
}

func main() {
	for {
		// Червоний світло
		displayLight("red")
		time.Sleep(5 * time.Second)

		// Жовтий світло
		displayLight("yellow")
		time.Sleep(2 * time.Second)

		// Зелений світло
		displayLight("green")
		time.Sleep(5 * time.Second)
	}
}
