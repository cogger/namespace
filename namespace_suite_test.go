package namespace_test

import (
	"crypto/rand"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"golang.org/x/net/context"

	"testing"
)

var namespaces = []string{}

const dictionary = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
const count = 100

func TestNamespace(t *testing.T) {
	RegisterFailHandler(Fail)

	namespaces = make([]string, count)
	for i := 0; i < count; i++ {
		var bytes = make([]byte, i+1)
		rand.Read(bytes)
		for k, v := range bytes {
			bytes[k] = dictionary[v%byte(len(dictionary))]
		}

		namespaces[i] = string(bytes)
	}
	RunSpecs(t, "Namespace Suite")
}

var ctx context.Context
var _ = BeforeSuite(func() {
	ctx = context.Background()
	Expect(ctx).NotTo(BeNil())
})
