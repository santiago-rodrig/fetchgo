package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(prepareURL(url))

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		defer resp.Body.Close()
		fmt.Fprintf(os.Stdout, "HTTP Status: %s\n\n", resp.Status)
		_, err = io.Copy(os.Stdout, resp.Body)

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}

func prepareURL(url string) string {
	if strings.HasPrefix(url, "http://") {
		return url
	}

	return "http://" + url
}
