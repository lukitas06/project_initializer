package go_init

import (
	// "fmt"
	// "io/fs"
	"io"
	"os"
	"os/exec"
	// "path/filepath"
)

	
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func CopiarArchivo(origen, destino string) error {
	srcFile, err := os.Open(origen)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	destFile, err := os.Create(destino)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	return err
}


func Main(){
	createEmptyFile := func(name string) {
        d := []byte("")
        check(os.WriteFile(name, d, 0644))
    }

	exec.Command("go", "mod", "init").Run()
	exec.Command("go", "mod", "tidy").Run()

	createEmptyFile("cmd/main.go")
	err := os.Mkdir("pkg", 0755)
	check(err)
	err = os.Mkdir("internal", 0755)
	check(err)
	err = os.Mkdir("/internal/api/handlers", 0755)
	check(err)
	err = os.Mkdir("/internal/api/auth", 0755)
	check(err)
	err = os.Mkdir("/internal/api/middleware", 0755)
	check(err)
	err = os.Mkdir("/internal/api/models", 0755)
	check(err)
	err = os.Mkdir("/internal/db", 0755)
	check(err)
	err = os.Mkdir("/internal/helpers", 0755)
	check(err)
	err = os.Mkdir("/internal/customerrors", 0755)
	check(err)

	err = CopiarArchivo("README.md", "README.md")
	check(err)

}