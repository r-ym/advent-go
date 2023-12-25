package util

import (
	"bufio"
	"fmt"
	"net/http"
)

func ReadURL(url string) ([]string, error) {
	// Make an HTTP GET request to the specified URL
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var inputArray []string
	// Read the response body line by line and push to array
	scanner := bufio.NewScanner(response.Body)
	for scanner.Scan() {
		inputArray = append(inputArray, scanner.Text())
	}
	fmt.Println(inputArray)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return inputArray, nil
}
