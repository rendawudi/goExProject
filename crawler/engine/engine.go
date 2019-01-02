package engine

import (
	"fmt"
	"github.com/gpmgo/gopm/modules/log"
	"goExProject/crawler/fetcher"
)

func Run(seeds ...Request) {
	var requestQueue []Request

	requestQueue = append(requestQueue, seeds...)

	for len(requestQueue) > 0 {
		request := requestQueue[0]
		requestQueue = requestQueue[1:]

		body, err := fetcher.Fetch(request.Url)

		if err != nil {
			log.Error("Url : %s ; Fetch request Error : %s", request.Url, err)
		}

		newRequests := request.ParserFunc(body)
		requestQueue = append(requestQueue, newRequests.Requests...)

		for _, m := range newRequests.Items {
			fmt.Printf("City: %s \n", m)
		}
	}
}
