package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		cwd, err := getWorkingDir()

		if err == nil {
			fmt.Print(cwd + "> ")
		} else {
			fmt.Println("Error while reading cwd")
			fmt.Print("> ")
		}

		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Error while reading input:", err)
			continue
		}

		if err = execCommand(input); err != nil {
			fmt.Println(err)
			continue
		}
	}
}

func getWorkingDir() (string, error) {
	cwd, err := os.Getwd()

	if err != nil {
		fmt.Println("Error while reading cwd", err)
		return "", err
	}

	path := strings.Split(cwd, "/")

	if len(path) > 4 {
		return "~/" + strings.Join(path[len(path)-3:], "/"), nil
	}

	return cwd, nil
}

func execCommand(input string) error {
	input = strings.TrimSpace(input)
	args := strings.Split(input, " ")

	switch args[0] {
	case "cd":
		if len(args) < 2 {
			return os.Chdir("/")
		}

		return os.Chdir(args[1])
	case "exit":
		os.Exit(0)
	}

	cmd := exec.Command(args[0], args[1:]...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
