package namespace

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Set", func() {
	It("should return blank when blank was added", func() {
		nCtx := Set(ctx, "")
		Expect(Get(nCtx)).To(BeEmpty())
	})

	It("should return what was set to it", func() {
		for _, ns := range namespaces {
			nCtx := Set(ctx, ns)
			Expect(Get(nCtx)).To(Equal(ns))
		}
	})
})
