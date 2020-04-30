package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	file, err := ioutil.ReadFile("websites.txt")

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	w := string(file)
	websites := []string(strings.Fields(w))

	fmt.Printf("%-10s %-10s %-10s\n", "Website", "Status", "Status Text")
	fmt.Println(strings.Repeat("-", 30))

	for _, website := range websites {

		req, err := http.Get(website)
		if err != nil {
			fmt.Println("Error: ", err)
		}

		fmt.Printf("%s %d %s\n", website, req.StatusCode, http.StatusText(req.StatusCode))
	}

}
