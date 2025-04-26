package pkg

import (
	// "fmt"
	// "io/fs"
	"os"
	// "path/filepath"
)

	
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func Main(){
	err := os.Mkdir("subdir", 0755)
    check(err)
}