package comparable_test

import (
	. "github.com/Altoros/cf-audit/comparable"
	. "github.com/Altoros/cf-audit/comparable/fakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Comparable", func() {
	var (
		tree1 ComparableItemWithName
		tree2 ComparableItemWithName
	)
	Describe("FindCollisions", func() {
		Describe("equal trees", func() {
			BeforeEach(func() {
				tree1 = NewNode("name", "value", 42, nil)
				tree2 = NewNode("name", "value", 42, nil)
			})
			It("returns no collisions", func() {
				collisions, err := FindCollisions(tree1, tree2)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(len(collisions)).To(Equal(0))
			})

		})
		// It("returns common pairs", func() {
		// 	commonPairs, _, _ := CompareByName(list1, list2)
		// 	Expect(len(commonPairs)).To(Equal(2))
		// 	Expect(commonPairs[0].First.GetName()).To(Equal("a2"))
		// 	Expect(commonPairs[0].First.GetName()).To(Equal(commonPairs[0].Second.GetName()))
		// 	Expect(commonPairs[1].First.GetName()).To(Equal("a3"))
		// 	Expect(commonPairs[1].First.GetName()).To(Equal(commonPairs[1].Second.GetName()))
		// })
		// It("returns alements that are not common", func() {
		// 	_, left1, left2 := CompareByName(list1, list2)
		// 	Expect(len(left1)).To(Equal(1))
		// 	Expect(len(left2)).To(Equal(1))
		// 	Expect(left1[0].GetName()).To(Equal("a1"))
		// 	Expect(left2[0].GetName()).To(Equal("a4"))
		// })
	})

})
