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

	entries, err := os.ReadDir(path)

	if err != nil{
		fmt.Println("error : ",err)
		return
	}

	for _, entry := range entries{

		// we only need git repository
		if entry.IsDir() && entry.Name()==".git"{

			*repos = append(*repos, path)
			
			return
		}
		if entry.IsDir(){

			fullpath := filepath.Join(path, entry.Name())

			// recursion until we find .git directories
			scan(fullpath, repos)
		}
	}
}