package main

import (
	"fmt"
	"os"
	"path/filepath"

	"authorstyle/internal/analyzer"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run cmd/main.go <file1.txt> <file2.txt>")
		return
	}

	file1 := os.Args[1]
	file2 := os.Args[2]

	text1, err := os.ReadFile(file1)
	if err != nil {
		fmt.Println("Error reading file1:", err)
		return
	}

	text2, err := os.ReadFile(file2)
	if err != nil {
		fmt.Println("Error reading file2:", err)
		return
	}

	fmt.Println("Analyzing:", filepath.Base(file1))
	counts1 := analyzer.AnalyzeText(string(text1))
	analyzer.DisplayStemLeaf(counts1)

	fmt.Println("\nAnalyzing:", filepath.Base(file2))
	counts2 := analyzer.AnalyzeText(string(text2))
	analyzer.DisplayStemLeaf(counts2)
}
