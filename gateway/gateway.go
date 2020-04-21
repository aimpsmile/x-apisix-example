package gateway

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/micro/go-micro/v2/server"
)

// Endpoint is a mapping between an RPC method and HTTP endpoint
type Endpoint struct {
	// RPC Method e.g. Greeter.Hello
	Name string
	// API Handler e.g rpc, proxy
	Handler string
	// HTTP Methods e.g GET, POST
	Method []string
	// HTTP Path e.g /greeter. Expect POSIX regex
	Path string
}

func strip(s string) string {
	return strings.TrimSpace(s)
}

func slice(s string) []string {
	var sl []string

	for _, p := range strings.Split(s, ",") {
		if str := strip(p); len(str) > 0 {
			sl = append(sl, strip(p))
		}
	}

	return sl
}

// Encode encodes an endpoint to endpoint metadata
func Encode(e *Endpoint) map[string]string {
	if e == nil {
		return nil
	}

	// endpoint map
	ep := make(map[string]string)

	// set vals only if they exist
	set := func(k, v string) {
		if v == "" {
			return
		}
		ep[k] = v
	}

	endpoint := strings.Split(e.Name, ".")
	set("grpc_service", endpoint[0])
	set("grpc_method", endpoint[1])
	set("handler", e.Handler)
	set("method", strings.Join(e.Method, ","))
	set("path", e.Path)

	return ep
}

// Decode decodes endpoint metadata into an endpoint
func Decode(e map[string]string) *Endpoint {
	if e == nil {
		return nil
	}

	return &Endpoint{
		Name:    fmt.Sprintf("%s.%s", e["grpc_service"], e["grpc_method"]),
		Method:  slice(e["method"]),
		Path:    e["path"],
		Handler: e["handler"],
	}
}

// Validate validates an endpoint to guarantee it won't blow up when being served
func Validate(e *Endpoint) error {
	if e == nil {
		return errors.New("endpoint is nil")
	}

	if e.Name == "" {
		return errors.New("name required")
	}
	_, err := regexp.CompilePOSIX(e.Path)
	if err != nil {
		return err
	}

	if e.Handler == "" {
		return errors.New("invalid handler")
	}

	return nil
}

// Usage:
//
// 	proto.RegisterHandler(service.Server(), new(Handler), api.WithEndpoint(
//		&api.Endpoint{
//			Name: "Greeter.Hello",
//			Path: "/greeter",
//		},
//	))
func WithEndpoint(e *Endpoint) server.HandlerOption {
	return server.EndpointMetadata(e.Name, Encode(e))
}
