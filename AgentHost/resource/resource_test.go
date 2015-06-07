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
	Describe("test get image tags", func(){
		It("should get ruby when resource is 0c689384ea3b", func(){
			Expect(GetImageTags("0c689384ea3b")).To(Equal([]string{"ruby"}))
		})
	})
})
