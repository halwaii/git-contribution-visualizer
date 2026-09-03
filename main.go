package main

import (
	"fmt"
	"os"

	"github.com/halwaii/git-contribution-visualizer/scan"
	"github.com/halwaii/git-contribution-visualizer/stats"
)

func main() {
	// fmt.Println("hello\n")
	if len(os.Args) < 2 {
		fmt.Println("Error : directory path is required")
		fmt.Println("usage : git-contribution-visulaizer <directory-path>")
		return
	}

	// get path from Command line
	path := os.Args[1]
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Printf("Error: The directory '%s' does not exist!\n", path)
		return
	}
	repos := scan.Scan(path)

	if len(repos)==0{
		fmt.Println("NO git repositories found in given path.")
		return
	}
	// fmt.Println("repositories found : ")
	// for _, repo := range repos{
	// 	fmt.Println(repo)
	// }

	stats.Stats(repos)
}
