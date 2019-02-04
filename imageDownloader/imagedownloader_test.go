package main

import (
	"fmt"
	"os"
	"testing"
)

func TestDirectoryexists(t *testing.T) {

	var testcases = []struct {
		path     string
		expected bool
	}{
		{"", false},
		{".", true},
		{"../", true},
		{"/..", true},
		{".../", false},
		{"~/", false},
		{"/...", false},
	}
	for _, test := range testcases {
		if item := directoryexists(test.path); item != test.expected {
			t.Errorf("directoryexists(%s) = %v", test.path, item)
		}
	}
}

func TestImagedownloaderFilename(t *testing.T) {
	var testcases = []struct {
		url              string
		expectedfilename string
	}{
		{"https://raw.githubusercontent.com/assertgo/icon/master/assertgo_512.png", "assertgo_512.png"},
		{"https://cdn4.iconfinder.com/data/icons/new-google-logo-2015/400/new-google-favicon-512.png", "new-google-favicon-512.png"},
		{"https://www.seriouseats.com/images/20080606-drpepper.jpg", "20080606-drpepper.jpg"},
		{"https://vsoch.github.io/assets/images/posts/learning-go/gophercises_jumping.gif", "gophercises_jumping.gif"},
	}
	for _, test := range testcases {
		downloaddir := "."
		if item := imagedownloader(downloaddir, test.url); item != test.expectedfilename {
			t.Errorf("imagedownloader(%s) = %v", test.url, item)
		}
		fmt.Println(downloaddir + "/" + test.expectedfilename)
		// cleanup
		os.Remove(downloaddir + "/" + test.expectedfilename)
	}
}

func TestImagedownloaderURLS(t *testing.T) {
	var testcases = []struct {
		url              string
		expectedfilename string
	}{
		{"https://raw.githubusercontent.com/assertgo/icon/master/assertgo_512.png", "assertgo_512.png"},
		{"https://cdn4.iconfinder.com/data/icons/new-google-logo-2015/400/new-google-favicon-512.png", "new-google-favicon-512.png"},
		{"https://www.seriouseats.com/images/20080606-drpepper.jpg", "20080606-drpepper.jpg"},
		{"https://vsoch.github.io/assets/images/posts/learning-go/gophercises_jumping.gif", "gophercises_jumping.gif"},
		{"https://google.com", ""},
		{"https://facebook.com", ""},
		{"https://bbc.com", ""},
		{"nwquebqweiqwjebiqwbe", ""},
		{"https://tfl.gov.uk", ""},
		{"https://landing.google.com/sre/sre-book/toc/", ""},
		{"https://www.google.com/images/branding/googlelogo/1x/googlelogo_color_272x92dp.png", ""},
	}
	for _, test := range testcases {
		downloaddir := "."
		if item := imagedownloader(downloaddir, test.url); item != test.expectedfilename {
			t.Errorf("imagedownloader(%s) = %v", test.url, item)
		}
		fmt.Println(downloaddir + "/" + test.expectedfilename)
		// cleanup
		os.Remove(downloaddir + "/" + test.expectedfilename)
	}
}
