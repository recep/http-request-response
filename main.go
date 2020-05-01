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

	fmt.Printf("%-30s %-10s %-10s\n", "Website", "Status", "Status Text")
	fmt.Println(strings.Repeat("-", 55))

	for _, website := range websites {

		resp, err := http.Get(website)
		if err != nil {
			fmt.Println("Error: ", err)
		}

		fmt.Printf("%-30s %-10d %-10s\n", website, resp.StatusCode, http.StatusText(resp.StatusCode))
	}

}
