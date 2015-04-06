package namespace

import "golang.org/x/net/context"

//Set sets the current namespace to the provide namespace for this context scope
func Set(ctx context.Context, namespace string) context.Context {
	return context.WithValue(ctx, &key, namespace)
}
