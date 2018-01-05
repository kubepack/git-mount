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
		fmt.Println("Hello World!!!")
		files, err := ioutil.ReadDir("/mypath")
		if err != nil {
			log.Println(err)
		}
		for key, value := range files {
			fmt.Println("hello key----", key)
			fmt.Println("hello value--", value)
			path := filepath.Join("/mypath", value.Name())
			fileInfo, err := os.Stat(path)
			if err != nil {
				log.Fatalln(err)
			}
			if fileInfo.IsDir() {
				outPath := filepath.Join(path, "_outlook")
				outFileInfo, err := os.Stat(outPath)
				if err == nil {
					fmt.Println("-------------", outFileInfo)
					cmd := exec.Command("kubectl", "apply", "-f", outPath)
					cmd.Stdout = os.Stdout
					cmd.Stderr = os.Stderr
					if err = cmd.Run(); err != nil {
						log.Println(err)
					}
					break
				}
			}
		}
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
