package proxy

import "io"

type Tunnel interface {
	Proxy(io.Reader, io.Writer)
}
