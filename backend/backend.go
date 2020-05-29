package backend

import "net/http"

type Backend interface {
	Router() http.Handler
	Run(addr ...string) error
}
