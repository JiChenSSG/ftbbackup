package service

import (
	"io/ioutil"
	"os"
	"sort"
)

func GetFileList(path, suffix string) ([]os.FileInfo, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var fileList []os.FileInfo
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if suffix != "" && file.Name()[len(file.Name())-len(suffix):] != suffix {
			continue
		}

		fileList = append(fileList, file)
	}

	// sort by mod time
	sort.Slice(fileList, func(i, j int) bool {
		return fileList[i].ModTime().Before(fileList[j].ModTime())
	})

	return fileList, nil
}
