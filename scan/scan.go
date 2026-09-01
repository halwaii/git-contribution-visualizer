package scan

import (
	"fmt"
	"os"
)

func Scan(path string) {

	entries, err := os.ReadDir(path)

	if err != nil{
		fmt.Println("error : ",err)
		return
	}

	for _, entry := range entries{
		fmt.Println(entry.Name())
	}
	// fmt.Println("scanning : ",path)
}