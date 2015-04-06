package namespace

import "golang.org/x/net/context"

//Is determines if the namespace is the desired namespace panics when set namespace is blank and desired is not
func Is(ctx context.Context, namespace string) bool {
	ns := c(ctx)
	if namespace != "" && ns == "" {
		panic(ErrNoNamespaceContext)
	}
	return ns == namespace
}
