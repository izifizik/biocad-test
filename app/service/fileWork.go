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
	if !strings.ContainsAny(fileName, ".txt") { // TODO: Если не тхт
		log.Println("Не тхт")
		return fileName, errors.New("File is not a txt file")
	}

	if _, err := os.Stat("/app/public" + fileName); err == nil { // TODO: если не существует то:
		log.Println("Не существует")
		return fileName, nil
	}
	log.Println("Существует")

	if !strings.ContainsAny(fileName, "()") { // TODO: если существует но нет скобок с повторениями (1), (2)....
		split := strings.Split(fileName, ".")
		newFileName, err := CheckName(split[0] + "(1)" + ".txt")
		if err != nil {
			return fileName, err
		}
		return newFileName, nil
	}

	s1 := strings.Index(fileName, "(")
	s2 := strings.Index(fileName, ")")
	number := fileName[s1+1 : s2]
	num, err := strconv.Atoi(number)
	if err != nil {
		return fileName, err
	}
	num++

	newFileName, err := CheckName(fileName[:s1] + "(" + strconv.Itoa(num) + ")" + fileName[s2+1:])
	if err != nil {
		return fileName, err
	}

	return newFileName, nil
}
