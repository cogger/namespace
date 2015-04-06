package namespace

import (
	"errors"
	"net/http"

	"golang.org/x/net/context"
)

//ErrNoNamespaceContext is called when no namespace is provided
var ErrNoNamespaceContext = errors.New("no namespace context")

var key = "namespace"

//Add adds the namespace context to a context
func Add(namespace string) func(context.Context, *http.Request) context.Context {
	return func(ctx context.Context, req *http.Request) context.Context {
		return context.WithValue(ctx, &key, namespace)
	}
}

func c(ctx context.Context) string {
	namespace, _ := ctx.Value(&key).(string)
	return namespace
}
