package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Download(url, checksum string) ([]byte, error) {
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

	return buf, nil
}
