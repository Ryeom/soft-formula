package internal

import "os"

func CheckDirectoryPath(dirPath string) {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err = os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
}

func CheckFilePath(filePath string) {
	if _, statErr := os.Stat(filePath); os.IsNotExist(statErr) {
		file, createErr := os.Create(filePath)
		if createErr != nil {
			panic(createErr)
		} else {
			file.Close()
		}
	}
}
