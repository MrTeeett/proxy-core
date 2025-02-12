//go:build !with_quic

package main

import (
	"net/url"
	"os"

	box "proxy-core"
)

func initializeHTTP3Client(instance *box.Box) error {
	return os.ErrInvalid
}

func fetchHTTP3(parsedURL *url.URL) error {
	return os.ErrInvalid
}
