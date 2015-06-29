package namespace

import (
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Get", func() {
	It("should return blank when nothing was added", func() {
		Expect(Get(ctx)).To(BeEmpty())
	})

	It("should return what was added to it", func() {
		for _, ns := range namespaces {
			nCtx := Add(ns)(ctx, &http.Request{})
			Expect(Get(nCtx)).To(Equal(ns))
		}
	})
})
