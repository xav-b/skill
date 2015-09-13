package main

import (
  "runtime"
  "flag"
  "log"
  "os"
	"fmt"
)

var usage = `usage: %s [OPTIONS] <URL>

  skill -dest bin https://dl.bintray......package.zip

skill - teach your system a new trick.

It downloads the archive at the given url and unpack its content
where you decide.

OPTIONS:
`

type Options struct {
  URL string
  Dest string
  // TODO "github.com/jbenet/go-multihash"
  Checksum string
}

func bintrayURL(user, project, version string) string {
  return fmt.Sprintf("https://dl.bintray.com/%s/%s/%s_%s_%s_%s.zip", user, project, project, version, runtime.GOOS, runtime.GOARCH)
}

func githubURL(user, project, version, pkg string) string {
  return fmt.Sprintf("https://github.com/%s/%s/releases/download/%s/%s",
                     user, project, version, pkg)
}

func getOpts() *Options {
  flag.Usage = func() {
		fmt.Fprintf(os.Stderr, usage, os.Args[0])
		flag.PrintDefaults()
	}
  var opts = new(Options)
  flag.StringVar(&opts.Dest, "dest", ".", "Directory to unpack archive files")
  flag.StringVar(&opts.Checksum, "checksum", "", "Archive checksum; leave blank to skip this step")
  flag.Parse()

  if len(flag.Args()) == 0 {
    flag.Usage()
    os.Exit(1)
  }

  opts.URL = flag.Arg(0)
  //opts.URL = "https://dl.bintray.com/mitchellh/vault/vault_0.1.0_linux_amd64.zip"
  //opts.URL = bintrayURL("mitchellh", "vault", "0.1.0")
  //opts.Checksum := "f6a8674a54f5b6288ba705bd21843cb1c848107e9ff6e7c17b4cc82cdb46789a"

  return opts
}

func main() {
  opts := getOpts()

  if err := os.MkdirAll(opts.Dest, 0777); err != nil {
		log.Printf("Unable to create directory: %s", opts.Dest)
		os.Exit(1)
	}

  // TODO Spinner:
  //    - https://github.com/briandowns/spinner
  //    - https://github.com/tj/go-spin
  //    - https://github.com/tj/stack/blob/master/pkg/logger
  log.Printf("downloading archive at %s", opts.URL)
  pack, err := Download(opts.URL, opts.Checksum)
  if err != nil {
    log.Printf("error: %v\n", err)
    os.Exit(1)
  }

  for _, zf := range pack.File {
    log.Printf("saving %s into %s", zf.Name, opts.Dest)
    written, err := ToDisk(zf, opts.Dest)
    if err != nil {
      log.Printf("error: %v\n", err)
      os.Exit(1)
    }
    log.Printf("done: %v bytes written", written)
    // TODO Add execution permissions
  }
}
