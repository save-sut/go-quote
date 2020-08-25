package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Quote struct {
	En     string `json:"en"`
	Author string `json:"author"`
}

func GetResponse(url string) (context, author string) {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
		return "", ""
	}
	defer response.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(response.Body)

	var quote Quote
	json.Unmarshal(bodyBytes, &quote)

	context = quote.En
	author = quote.Author
	return context, author
}

func WriteToReadme(context, author string) {
	// read the whole file at once
	message, err := ioutil.ReadFile("static_readme.md")
	if err != nil {
		panic(err)
	}

	message = append(message[:], []byte("\n**Quote of the day**\n")...)
	message = append(message[:], []byte(`> `+context+``)...)
	message = append(message[:], []byte("\n\n")...)
	message = append(message[:], []byte(`-`+author+``)...)

	// write the whole body at once
	err = ioutil.WriteFile("README.md", message, 0644)
	if err != nil {
		panic(err)
	}
}

func main() {
	url := "https://programming-quotes-api.herokuapp.com/quotes/random/lang/en/"
	context, author := GetResponse(url)
	WriteToReadme(context, author)
}
