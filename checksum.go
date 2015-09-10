package main

import (
  "fmt"
  "io"
  "hash"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
)

// Stolen from https://github.com/gosexy/checksum/blob/master/checksum.go
func createHash(method string) hash.Hash {
	var h hash.Hash

	switch method {
	case "MD5":
		h = md5.New()
	case "SHA1":
		h = sha1.New()
	case "SHA224":
		h = sha256.New224()
	case "SHA256":
		h = sha256.New()
	case "SHA384":
		h = sha512.New384()
	case "SHA512":
		h = sha512.New()
	default:
    // NOTE return err, nil ? panic ? ? warning and default to 256 ?
		panic("Unknown hashing method.")
	}

	return h
}

// TODO More methods https://github.com/gosexy/checksum/blob/master/checksum.go
func Checksum(content []byte, method string) string {
    hash := createHash(method)
    io.WriteString(hash, string(content))
    return fmt.Sprintf("%x", hash.Sum(nil))
}
