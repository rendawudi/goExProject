package appHandle

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type userError string

func (err userError) Error() string {
	return err.Message()
}

func (err userError) Message() string {
	return string(err)
}

const Prefix = "/list/"

func FileHandle(w http.ResponseWriter,
				r *http.Request) error  {
	if strings.Index(r.URL.Path, Prefix) != 0 {
		return userError(fmt.Sprint("FileHandle -- Path error"))
	}

	filePath := r.URL.Path[len(Prefix):]

	file, err := os.Open(filePath)
	if err != nil {
		return err
	}

	defer file.Close()

	fileData, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	w.Write(fileData)

	return nil
}
