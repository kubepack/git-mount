package main

import (
	"fmt"
	"os"
	"path/filepath"
	"github.com/Masterminds/vcs"
	"io/ioutil"
	"log"
	"os/exec"
)

func main()  {
	for ; ; {
		dicision := false
		files, err := ioutil.ReadDir("/mypath")
		if err != nil {
			log.Println(err)
		}
		for _, value := range files {
			path := filepath.Join("/mypath", value.Name())
			fileInfo, err := os.Stat(path)
			if err != nil {
				log.Fatalln(err)
			}
			if fileInfo.IsDir() {
				outPath := filepath.Join(path, "_outlook")
				_, err := os.Stat(outPath)
				if err == nil {
					cmd := exec.Command("kubectl", "apply", "-R" ,"-f", outPath)
					cmd.Stdout = os.Stdout
					cmd.Stderr = os.Stderr

					dicision = true
					if err = cmd.Run(); err != nil {
						dicision = false
						fmt.Println(err)
						log.Println(err)
					}

					break
				}
			}
		}
		if dicision {
			break
		}
	}

	for ; ;  {
		fmt.Println("")
	}
}

func getRootDir(path string) (vcs.Repo, error) {
	var err error
	for ; ; {
		repo, err := vcs.NewRepo("", path)
		if err == nil {
			return repo, err
		}
		if os.Getenv("HOME") == path {
			break
		}
		path = filepath.Dir(path)
	}

	return nil, err
}
