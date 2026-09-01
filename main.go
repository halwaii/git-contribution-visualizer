package main

import (
	"fmt"

	"github.com/halwaii/git-contribution-visualizer/scan"
	"github.com/halwaii/git-contribution-visualizer/stats"
)

func main() {
	fmt.Println("hello\n")
	repos := scan.Scan("C:\\Users\\HP\\OneDrive\\Desktop\\Magic")

	fmt.Println("repositories found : ")
	for _, repo := range repos{
		fmt.Println(repo)
	}
	stats.Stats()
}