package main

import "io"

type Request struct {
	method string
	url    string
	body   io.Reader
}
