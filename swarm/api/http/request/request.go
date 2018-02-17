package request

// Request wraps http.Request and also includes the parsed bzz URI

import (
	"net/http"

	"github.com/ethereum/go-ethereum/swarm/api/http/uri"
)

type Request struct {
	http.Request

	uri *api.URI
}
