package namespace

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Is", func() {
	It("should panic when blank was added", func() {
		nCtx := Set(ctx, "")
		for _, ns := range namespaces {
			Expect(func() { Is(nCtx, ns) }).Should(Panic())
		}
	})

	It("should return true when ns match", func() {
		for _, ns := range namespaces {
			nCtx := Set(ctx, ns)
			Expect(Is(nCtx, ns)).To(BeTrue())
		}
	})
})
