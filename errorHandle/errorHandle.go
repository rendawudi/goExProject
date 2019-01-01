package errorHandle

import (
	"log"
	"net/http"
	"os"
)

type appHandler func(writer http.ResponseWriter,
	r *http.Request) error

type httpHandler func(writer http.ResponseWriter,
	r *http.Request)

func ErrorHandler(handler appHandler) func(writer http.ResponseWriter,
	r *http.Request) {
	return func(writer http.ResponseWriter,
		r *http.Request){
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic: %v \n", r)
				http.Error(writer,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()
		err := handler(writer, r)

		if err != nil {
			log.Printf("error occured: inHandle -- %s", err.Error())

			if ueserErro, ok := err.(userError); ok {
				http.Error(writer,
					ueserErro.Message(),
					http.StatusBadRequest)
				return
			}
		}

		code := http.StatusOK

		switch {
		case os.IsExist(err):
			code = http.StatusNotFound
		case os.IsPermission(err):
			code = http.StatusForbidden
		default:
            code = http.StatusInternalServerError
		}

		http.Error(writer,
			http.StatusText(code),
			code)
	}


}

type userError interface {
	error
	Message() string
}