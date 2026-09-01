package scan

import (
	"fmt"
	"os"
	"path/filepath"
)

// scan() -> find all repos -> store them -> return to main.go -> pass to stats()

// main function to return all repos path
func Scan(path string) []string{
	repos := []string{}

	scan(path, &repos)

	return repos
}

// helper function to find all paths
func scan(path string, repos *[]string){

	// os.ReadDir -> reads directory from file system
	// and return 2 values
	entries, err := os.ReadDir(path)

	if err != nil{
		fmt.Println("error : ",err)
		return
	}

	for _, entry := range entries{

		// we only need git repository
		if entry.IsDir() && entry.Name()==".git"{
			fmt.Println("git repository path found : ",path)

			// add paths to repos
			// using pointer to use real slice
			*repos = append(*repos, path)
			// repos = ["dsa","chess engine","lolcat","git contribution visulazer"]
			
			return
		}
		// IsDir -> checks if it is directory or not
		if entry.IsDir(){

			// filepath -> provides utility to parse, manipulate and construct file paths
			// filepath.join to join multiple strings using correct OS seperator
			fullpath := filepath.Join(path, entry.Name())

			// recursion until we find .git directories
			scan(fullpath, repos)
		}
	}
	// fmt.Println("scanning : ",path)
}

// we need to store these file paths and pass them to main and then pass
// them to stats to generate graph

// scan() -> find all repos -> store them -> return to main.go -> pass to stats()