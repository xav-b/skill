package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/c4milo/unzipit"
)

const RWX_FILE = 0755
const FULL_FILE = 0777

var EXTENSIONS = []string{"zip", "tar.gz"}

func main() {
	opts := getOpts()

	if err := os.MkdirAll(opts.Out, FULL_FILE); err != nil {
		log.Printf("Unable to create directory %s: %v\n", opts.Out, err)
		os.Exit(1)
	}

	// TODO Spinner:
	//    - https://github.com/briandowns/spinner
	//    - https://github.com/tj/go-spin
	//    - https://github.com/tj/stack/blob/master/pkg/logger
	log.Printf("downloading archive at %s", opts.URL)
	// TODO Bring back checksum check
	//stream, err := Download(opts.URL, opts.Checksum)
	res, err := http.Get(opts.URL)
	if err != nil {
		log.Printf("error: %v\n", err)
		os.Exit(1)
	}

	// TODO can force extention from the cli
	for _, ext := range EXTENSIONS {
		if strings.HasSuffix(opts.URL, ext) {
			_, err := unzipit.UnpackStream(res.Body, opts.Out)
			if err != nil {
				log.Printf("error: %v\n", err)
				os.Exit(1)
			}
			log.Printf("wrote file to %s\n", opts.Out)
			return
		}
	}

	// else
	buf, _ := ioutil.ReadAll(res.Body)
	partials := strings.Split(opts.URL, "/")
	filename := partials[len(partials)-1]
	if err := ioutil.WriteFile(path.Join(opts.Out, filename), buf, RWX_FILE); err != nil {
		log.Printf("error: %v\n", err)
		os.Exit(1)
	}

	log.Printf("wrote file to %s\n", opts.Out)
}
