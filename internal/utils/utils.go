package utils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// prompts for user input
func GetInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}
	return strings.TrimSpace(input)
}

// converts string y or n to boolean
func ParseBoolean(input string) (bool, error) {
	switch input {
	case "y":
		return true, nil
	case "n":
		return false, nil
	case "":
		return true, nil
	default:
		return false, fmt.Errorf("invalid input, should be boolean")
	}
}

// Execute executes shell commands with arguments
func Execute(workDir, script string, args ...string) error {
	cmd := exec.Command(script, args...)

	if workDir != "" {
		cmd.Dir = workDir
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
