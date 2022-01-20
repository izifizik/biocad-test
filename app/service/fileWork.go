package service

import (
	"errors"
	"io"
	"log"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
)

func UploadFile(filename string, file multipart.File) (string, error) {
	newFileName, err := CheckName(filename)
	if err != nil {
		return newFileName, err
	}
	out, err := os.Create("app/public/" + newFileName)
	if err != nil {
		return newFileName, err
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		return newFileName, err
	}
	return newFileName, nil
}

func CheckName(fileName string) (string, error) {
	if !strings.ContainsAny(fileName, ".txt") { // Если не тхт
		log.Println("Не тхт")
		return fileName, errors.New("File is not a txt file")
	}

	for {
		ok := statCheck(fileName)
		if ok {
			break
		}
		if !strings.ContainsAny(fileName, "()") { // если существует, но нет скобок с повторениями (1), (2)....
			split := strings.Split(fileName, ".")
			fileName = split[0] + "(1)" + ".txt"
			continue
		}

		s1 := strings.Index(fileName, "(")
		s2 := strings.Index(fileName, ")")

		number := fileName[s1+1 : s2]
		num, err := strconv.Atoi(number)
		if err != nil {
			return fileName, err
		}

		num += 1

		fileName = fileName[:s1] + "(" + strconv.Itoa(num) + ")" + fileName[s2+1:]
	}

	return fileName, nil
}

func statCheck(fileName string) bool {
	if _, err := os.Open("./app/public/" + fileName); err != nil {
		return true
	}
	return false
}


