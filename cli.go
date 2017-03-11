package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

const LOCAL_BIN_PATH = "/usr/local/bin"

const usage = `usage: %s [OPTIONS] <URL>

  skill -dest bin https://dl.bintray......package.zip

skill - teach your system a new trick.

It downloads (in-memory) the archive at the given url and unpack its content
where you decide.

OPTIONS:
`

type ShortParser func(string) string

func parseRawGithub(shortcut string) string {
	partials := strings.Split(shortcut, "@")
	user_project := partials[0]
	version := partials[1]

	// usage: -short rawgh "rupa/z@master/z.sh"
	return fmt.Sprintf("https://raw.githubusercontent.com/%s/%s", user_project, version)
}

func parseGithub(shortcut string) string {
	partials := strings.Split(shortcut, "@")
	user_project := partials[0]
	version := partials[1]

	return fmt.Sprintf("https://github.com/%s/releases/download/%s", user_project, version)
}

// TODO -rename ? or detect from -out ?
type Options struct {
	URL      string
	Shortcut string
	Out      string
	// TODO "github.com/jbenet/go-multihash"
	Checksum string
}

func getOpts() *Options {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, usage, os.Args[0])
		flag.PrintDefaults()
	}
	var opts = new(Options)
	flag.StringVar(&opts.Out, "out", LOCAL_BIN_PATH, "Directory to unpack archive files")
	flag.StringVar(&opts.Checksum, "checksum", "", "Archive checksum; leave blank to skip this step")
	flag.StringVar(&opts.Shortcut, "short", "", "extrapolate the given url (gh|rawgh)")
	flag.Parse()

	if len(flag.Args()) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	switch opts.Shortcut {
	case "gh":
		opts.URL = parseGithub(flag.Arg(0))
	case "rawgh":
		opts.URL = parseRawGithub(flag.Arg(0))
	default:
		log.Printf("no shortcut provided, using as is")
		opts.URL = flag.Arg(0)
	}

	return opts
}
