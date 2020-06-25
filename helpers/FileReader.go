package helpers

import (
	"io/ioutil"
	"regexp"
	"strings"

	"gopkg.in/yaml.v2"
)

func getAllFilesPath(dir string, ext string, basePath string) []string {
	filesPath := []string{}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return filesPath
	}

	for _, file := range files {
		filename := file.Name()

		isExt, err := regexp.Match(ext, []byte(filename))

		if err != nil {
			continue
		}

		if file.IsDir() {
			filesPath = append(filesPath, getAllFilesPath(dir+"/"+filename, ext, basePath+filename+"/")...)
		} else if isExt {
			filesPath = append(filesPath, basePath+strings.Replace(filename, ext, "", 1))
		}

	}

	return filesPath
}

func ReadYML(filename string, data interface{}) (string, error) {

	file, err := ioutil.ReadFile(filename)

	if err != nil {
		return "File not found", err
	}

	err = yaml.Unmarshal(file, data)

	if err != nil {
		return "Failed to parse yaml file", err
	}

	return "", nil
}
