package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
)

func main() {

	downloaddir := os.Args[1]
	urls := os.Args[2:]

	if len(urls) == 0 {
		panic("No download url specified")
	}

	if !directoryexists(downloaddir) {
		os.Exit(1)
	}
	for i := 0; i < len(urls); i++ {
		imagedownloader(downloaddir, urls[i])
	}
}

func imagedownloader(downloaddir string, url string) string {

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("The following url is not valid " + url)
		return ""
	}

	if !(strings.Contains(resp.Header.Get("Content-Type"), "image/")) {
		fmt.Println("URL does not contain image format:" + url)
		return ""
	}

	defer resp.Body.Close()

	out, err := os.Create(downloaddir + "/" + path.Base(url))
	if err != nil {
		fmt.Println("Error Downloading file, skipping")
		return ""
	}

	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Println("Error Downloading file, skipping")
		return ""
	}
	return path.Base(url)
}

func directoryexists(downloadir string) bool {
	errormessage := ("Directry to store downloaded file does not exist")
	if _, err := os.Stat(downloadir); os.IsNotExist(err) {
		fmt.Println(errormessage)
		return false
	}
	return true
}
