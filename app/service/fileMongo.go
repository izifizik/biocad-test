package service

import (
	"Biocad/app/database"
	"Biocad/app/model"
	"io/ioutil"
	"strings"
)

func InsertByFileName(filename string) error {
	var arr []string
	file, err := ioutil.ReadFile("app/public/" + filename)
	if err != nil {
		return err
	}

	split := strings.Split(string(file), "\n")

	for i:=1; i < len(split); i++ {
		arr = append(arr, split[i])
	}

	obj := model.File{
		Filename: filename,
		Header: split[0],
		Data: arr,
	}

	err = database.InsertFile(obj)
	if err != nil {
		return err
	}

	return nil
}