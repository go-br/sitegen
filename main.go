package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/crgimenes/goconfig"
	"github.com/nuveo/log"
)

type config struct {
	InputFolder  string `json:"i" cfg:"i" cfgDefault:"./"`
	OutputFolder string `json:"o" cfg:"o"`
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

	cfg := config{}

	err := goconfig.Parse(&cfg)
	if err != nil {
		log.Errorln(err)
		return
	}

	if cfg.OutputFolder == "" {
		log.Errorln("error output folder not indicated")
		return
	}

	err = filepath.Walk(cfg.InputFolder, visit)
	if err == nil || err == io.EOF {
		return
	}
	log.Errorln(err)
}
