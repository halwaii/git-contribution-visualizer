package main

import (
	"fmt"

	"github.com/halwaii/git-contribution-visualizer/scan"
	"github.com/halwaii/git-contribution-visualizer/stats"
)

func main() {
	fmt.Println("hello")
	scan.Scan("C:\\Users\\HP\\OneDrive\\Desktop\\Magic")
	stats.Stats()
}