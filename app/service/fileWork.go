package service

import (
	"errors"
	"io"
	"mime/multipart"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func UploadFile(filename string, file multipart.File ) error {
	newFileName, err := CheckName(filename)
	if err != nil {
		return err
	}
	out, err := os.Create("app/public/" + newFileName)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		return err
	}
	return nil
}

func CheckName(fullFileName string) (string, error) {
	if _, err := os.Stat("./app/public/" + fullFileName); err != nil { // TODO: если уже есть то:
		fileName := strings.Split(fullFileName, ".")

		if fileName[1] != "txt" { // TODO: проверка на вход тхт файла
			return fullFileName, errors.New("Bad request")
		}

		pattern := `\((\d){0,}\)`

		fileNameArr := []string { // TODO: если итерироваться по строчкам то каждая буква будет иметь (*)
			fileName[0],
		}

		for _, a := range fileNameArr {
			ok, err := regexp.Match(pattern, []byte(a))
			if err != nil {
				return fullFileName, errors.New("Error with mach in regular")
			}
			if ok { // TODO: если есть соответствие
				split := strings.Split(fileName[0], "(")
				split2 := strings.Split(split[1], ")")
				num, err := strconv.Atoi(split2[0])
				if err != nil {
					return fullFileName, errors.New("Error with atoi")
				}
				num++
				return split[0] + "(" + strconv.Itoa(num) + ").txt", nil
			}
			return fileName[0] + "(1).txt", nil
		}

	} else if os.IsNotExist(err) { // TODO: если нет то:
		s := strings.Split(fullFileName, ".")
		if s[1] != "txt" {
			return fullFileName, errors.New("Bad extension")
		}
		return fullFileName, nil
	} else {
		return "", err //файл может существовать, а может и не существовать.
	}
	return fullFileName, errors.New("Hz kak on mojet byt'")
}