package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var result strings.Builder

	fmt.Println("Paste your configuration here and press CTRT+D when done:")

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		// Append the line to result only if it's not empty
		if line != "" {
			if result.Len() > 0 {
				// Add a space before appending of the builer is not empty
				result.WriteRune(' ')
			}
			result.WriteString(line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "reading standard input: %v", err)
		os.Exit(1)
	}

	// Output the single line configuration
	fmt.Println(result.String())
}
