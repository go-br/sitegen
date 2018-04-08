package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

func execHelper(path, name string, arg ...string) (err error) {
	cmd := exec.Command(name, arg...)
	cmd.Dir = path
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	return
}

func fileExists(path string) (ret bool) {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		fmt.Println(err)
		return
	}
	ret = true
	return
}

func visit(path string, f os.FileInfo, perr error) error {
	if perr != nil {
		return perr
	}
	if !f.IsDir() {
		return nil
	}
	fe := fileExists(path + "/README.md")
	if !fe {
		return nil
	}

	/*
		if f.Name() == "vendor" {
			return filepath.SkipDir
		}
	*/
	path, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	path = filepath.Join(path, "/README.md")
	fmt.Println(path)

	return nil
}

func main() {
	err := filepath.Walk("./", visit)
	if err == nil || err == io.EOF {
		return
	}
	fmt.Println(err)
}
