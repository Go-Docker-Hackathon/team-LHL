package resource

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Resource", func() {
	Describe("test get image", func(){
		It("should get 0c689384ea3b when resource is ruby", func(){
			Expect(GetImage([]string{"ruby"})).To(Equal("0c689384ea3b"))
		})
	})
})
