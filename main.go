package main

import (
	"bufio"
	"flag"
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

func filterRows(rows []string, match string) []string {
	if match == "" {
		return rows
	}
	filtered := make([]string, 0, len(rows))
	for _, row := range rows {
		if strings.Contains(strings.ToLower(row), strings.ToLower(match)) {
			filtered = append(filtered, row)
		}
	}
	return filtered
}

func main() {
	match := flag.String("match", "", "only show rows containing this text")
	dryRun := flag.Bool("dry-run", false, "show the selected inventory without changing Docker")
	flag.Parse()
	if *dryRun {
		fmt.Println("Dry run: no Docker changes will be made.")
	}
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
	printRows(filterRows(containers, *match))
	fmt.Println("Images")
	printRows(filterRows(images, *match))
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
