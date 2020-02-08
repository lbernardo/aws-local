package storage

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func StorageContent(file, content string) error {
	fmt.Println("File",file)
	fop,err := os.Create(file)

	defer fop.Close()
	if err != nil {
		return err
	}
	_,err = fop.Write([]byte(content))

	return err
}


func CreateVolume(path string) error{
	return os.MkdirAll(path,0755)
}

func ReadFile(path string) ([]byte, string, error) {
	content,err :=  ioutil.ReadFile(path)
	if err != nil {
		return nil,"",err
	}
	typeFile, err := GetFileContentType(path)
	return content, typeFile, err
}

func GetFileContentType(path string) (string, error) {
	// Open File
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)

	_, err = f.Read(buffer)
	if err != nil {
		return "", err
	}

	// Use the net/http package's handy DectectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)

	return contentType, nil
}