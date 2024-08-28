package pkg

import (
	"os"
	"strings"
)

var protectedNames = []string{"404", "static", "tag", "category"}

func CheckContentDir() error {
	folder, err := os.Open("content")
	if err != nil {
		return err
	}

	rootElements, err := folder.Readdir(-1)
	if err != nil {
		return err
	}

	var dirNames []string
	var fileNames []string

	for _, element := range rootElements {
		if element.IsDir() {
			dirNames = append(dirNames, element.Name())
		} else {
			fileNames = append(fileNames, element.Name())
		}
	}

	for _, dirName := range dirNames {
		for _, protectedName := range protectedNames {
			if dirName == protectedName {
				return &ProtectedNameError{dirName, "directory "}
			}
		}
	}

	for _, fileName := range fileNames {
		fileName = strings.Split(fileName, ".")[0]
		for _, protectedName := range protectedNames {
			if fileName == protectedName {
				return &ProtectedNameError{fileName, "file"}
			}
		}
	}

	return nil
}
