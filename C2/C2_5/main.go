package main

import (
	"net/http"
)

func main() {

}

func FetchURL(url string) string {
	_, err := http.Get(url)
	if err != nil {
		return "Failed to fetch"
	}
	return "Successfully fetched"
}
