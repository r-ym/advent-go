package util

import (
	"bufio"
	"net/http"
	"os"
)

func ReadURL(url string) ([]string, error) {
	// Read the session token from .session file
	file, err := os.Open("../.session")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	session_file := bufio.NewScanner(file)
	session_file.Scan()
	session := session_file.Text()
	if err := session_file.Err(); err != nil {
		return nil, err
	}
	

	// Create an HTTP client
	client := &http.Client{}
	// Create an HTTP GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Add session cookie to request
	
	cookie := &http.Cookie{Name: "session", Value: session}
	req.AddCookie(cookie)


	// Make the HTTP request
	response, err := client.Do(req)
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

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return inputArray, nil
}
