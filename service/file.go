package service

import (
	"os"
	"sort"
	"strings"
)

func GetLatestFile(path, suffix string) (os.DirEntry, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	fileList := make([]os.DirEntry, 0)
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if suffix != "" && !file.IsDir() && !strings.HasSuffix(file.Name(), suffix) {
			continue
		}

		fileList = append(fileList, file)
	}

	sort.Slice(fileList, func(i, j int) bool {
		info1, _ := fileList[i].Info()
		info2, _ := fileList[j].Info()
		return info1.ModTime().After(info2.ModTime())
	})

	if len(fileList) == 0 {
		return nil, nil
	}

	return fileList[0], nil
}
