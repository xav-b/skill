package main

import (
  "fmt"
  "path"
  "os"
  "archive/zip"
	"bytes"
  "io"
	"io/ioutil"
	"net/http"
)

// TODO Support more format
// an interesting simplification / resource ? https://github.com/c4milo/unzipit
//func Unpack(res *http.Response) (*zip.Reader, error) {
func Unpack(b []byte) (*zip.Reader, error) {
  	r := bytes.NewReader(b)
  	return zip.NewReader(r, int64(r.Len()))
}

func Download(url, checksum string) (*zip.Reader, error) {
  	res, err := http.Get(url)
  	if err != nil {
  		return nil, err
    }
  	defer res.Body.Close()

  	buf, err := ioutil.ReadAll(res.Body)
    if err != nil {
      return nil, err
    }

    sum := Checksum(buf, "SHA256")
    if checksum != sum && checksum != "" {
      return nil, fmt.Errorf("invalid checksum %s / %s\n", sum, checksum)
    }

  	return Unpack(buf)
}

func ToDisk(zf *zip.File, dest string) (int64, error) {
    dst, err := os.Create(path.Join(dest, zf.Name))
    if err != nil {
      return 0, err
    }
    defer dst.Close()

    // TODO Detect and make executable ?

    src, err := zf.Open()
    if err != nil {
      return 0, err
    }
    defer src.Close()

    return io.Copy(dst, src)
}
