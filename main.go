package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var usage = `Usage of gilo
 example: gilo Tyler Wray
 output:
  _______    _            __          __
 |__   __|  | |           \ \        / /
    | |_   _| | ___ _ __   \ \  /\  / / __ __ _ _   _
    | | | | | |/ _ \ /__|   \ \/  \/ / /__/ _\ | | | |
    | | |_| | |  __/ |       \  /\  /| | | (_| | |_| |
    |_|\__, |_|\___|_|        \/  \/ |_|  \__,_|\__, |
        __/ |                                    __/ |
       |___/                                    |___/
`

func main() {
	args := os.Args[1:]

	if len(args) == 0 || contains(args, "-help") || contains(args, "--help") || contains(args, "-h") {
		fmt.Println(usage)
		os.Exit(0)
	}

	art := ascii(args...)
	fmt.Printf("\n%s\n", art)
}

func contains(slice []string, finder string) bool {
	for _, value := range slice {
		if finder == value {
			return true
		}
	}
	return false
}

func ascii(words ...string) string {
	url := fmt.Sprintf("http://artii.herokuapp.com/make?text=%s", strings.Join(words, "+"))

	resp, err := http.Get(url)

	if err != nil {
		fmt.Print(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	return string(body)
}
