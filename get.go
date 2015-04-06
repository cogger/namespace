package namespace

import "golang.org/x/net/context"

//Get gets the namespace form the context
func Get(ctx context.Context) string {
	return c(ctx)
}
