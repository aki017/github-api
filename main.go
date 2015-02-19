package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/octokit/go-octokit/octokit"
)

func main() {
	if len(os.Args) == 1 {
		return
	}
	client := octokit.NewClient(new(octokit.NetrcAuth))
	req, _ := client.NewRequest(os.Args[1])
	var output interface{}
	_, err := req.Get(&output)
	if err != nil {
		fmt.Printf("%s", err)
	}
	j, _ := json.Marshal(output)
	fmt.Println(string(j))
}
