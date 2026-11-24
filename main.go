package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func dockerList(kind string) ([]string, error) {
	format := "{{.ID}}\t{{.CreatedAt}}\t{{.Names}}"
	if kind == "images" {
		format = "{{.ID}}\t{{.CreatedAt}}\t{{.Repository}}:{{.Tag}}"
	}
	args := []string{"ps", "-a", "--format", format}
	if kind == "images" {
		args = []string{"images", "--format", format}
	}
	output, err := exec.Command("docker", args...).CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("docker %s: %w: %s", kind, err, strings.TrimSpace(string(output)))
	}
	return nonEmptyLines(string(output)), nil
}

func nonEmptyLines(input string) []string {
	var lines []string
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		if line := strings.TrimSpace(scanner.Text()); line != "" {
			lines = append(lines, line)
		}
	}
	return lines
}

func main() {
	containers, err := dockerList("containers")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	images, err := dockerList("images")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println("Containers")
	printRows(containers)
	fmt.Println("Images")
	printRows(images)
}

func printRows(rows []string) {
	if len(rows) == 0 {
		fmt.Println("(none)")
		return
	}
	for _, row := range rows {
		fmt.Println(row)
	}
}
