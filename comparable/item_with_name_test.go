package comparable_test

import (
	. "github.com/Altoros/cf-audit/comparable"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type FakeItem struct {
	Name string
}

func (i FakeItem) GetName() string {
	return i.Name
}

var _ = Describe("ItemWithName", func() {
	var (
		list1 NamedItems
		list2 NamedItems
	)
	Describe("CompareByName", func() {
		BeforeEach(func() {
			list1 = NamedItems{FakeItem{Name: "a1"}, FakeItem{Name: "a2"}, FakeItem{Name: "a3"}}
			list2 = NamedItems{FakeItem{Name: "a2"}, FakeItem{Name: "a3"}, FakeItem{Name: "a4"}}
		})
		It("returns common pairs", func() {
			commonPairs, _, _ := CompareByName(list1, list2)
			Expect(len(commonPairs)).To(Equal(2))
			Expect(commonPairs[0].First.GetName()).To(Equal("a2"))
			Expect(commonPairs[0].First.GetName()).To(Equal(commonPairs[0].Second.GetName()))
			Expect(commonPairs[1].First.GetName()).To(Equal("a3"))
			Expect(commonPairs[1].First.GetName()).To(Equal(commonPairs[1].Second.GetName()))
		})
		It("returns alements that are not common", func() {
			_, left1, left2 := CompareByName(list1, list2)
			Expect(len(left1)).To(Equal(1))
			Expect(len(left2)).To(Equal(1))
			Expect(left1[0].GetName()).To(Equal("a1"))
			Expect(left2[0].GetName()).To(Equal("a4"))
		})
	})
})
