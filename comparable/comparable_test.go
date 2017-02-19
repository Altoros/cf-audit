package comparable_test

import (
	. "github.com/Altoros/cf-audit/comparable"
	. "github.com/Altoros/cf-audit/comparable/fakes"
	"github.com/davecgh/go-spew/spew"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Comparable", func() {
	Describe("FindCollisions", func() {
		var (
			tree1 ComparableItemWithName
			tree2 ComparableItemWithName
		)
		Describe("equal trees", func() {
			Describe("child free", func() {
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
			Describe("with family", func() {
				BeforeEach(func() {
					tree1 = NewNode("parent", "value", 42, NewChildren(
						NewNode("child1", "value1", 42, NewChildren(
							NewNode("grandson1", "value1", 42, NewChildren(
								NewNode("great-grandson1", "value1", 42, nil))))),
						NewNode("child2", "value2", 42, NewChildren(
							NewNode("grandson1", "value1", 42, nil))),
						NewNode("child3", "value3", 42, nil)))
					tree2 = NewNode("parent", "value", 42, NewChildren(
						NewNode("child1", "value1", 42, NewChildren(
							NewNode("grandson1", "value1", 42, NewChildren(
								NewNode("great-grandson1", "value1", 42, nil))))),
						NewNode("child2", "value2", 42, NewChildren(
							NewNode("grandson1", "value1", 42, nil))),
						NewNode("child3", "value3", 42, nil)))
				})
				It("returns no collisions", func() {
					collisions, err := FindCollisions(tree1, tree2)
					Expect(err).ShouldNot(HaveOccurred())
					Expect(len(collisions)).To(Equal(0))
				})
			})
		})

		Describe("trees with collisions", func() {
			Describe("different names", func() {
				BeforeEach(func() {
					tree1 = NewNode("name1", "value", 42, nil)
					tree2 = NewNode("name2", "value", 42, nil)
				})
				It("returns an error", func() {
					_, err := FindCollisions(tree1, tree2)
					Expect(err).Should(HaveOccurred())
					Expect(err.Error()).To(Equal("Different names"))
				})
			})
			Describe("different value", func() {
				BeforeEach(func() {
					tree1 = NewNode("name", "value1", 42, nil)
					tree2 = NewNode("name", "value2", 42, nil)
				})
				It("returns 1 collision (with 1 mismatch) because all nodes are the same, but value is wrong", func() {
					collisions, err := FindCollisions(tree1, tree2)
					Expect(err).ShouldNot(HaveOccurred())
					Expect(len(collisions)).To(Equal(1))
				})
			})
			Describe("with 2 different values", func() {
				BeforeEach(func() {
					tree1 = NewNode("name", "value1", 42, nil)
					tree2 = NewNode("name", "value2", 4242, nil)
				})
				It("returns 1 collision with 2 mismatches", func() {
					collisions, err := FindCollisions(tree1, tree2)
					Expect(err).ShouldNot(HaveOccurred())
					Expect(len(collisions)).To(Equal(1))
				})
			})

			Describe("with family", func() {
				BeforeEach(func() {
					tree1 = NewNode("parent", "value", 42, NewChildren(
						NewNode("child1", "value1", 42, NewChildren(
							NewNode("grandson1", "value1", 42, NewChildren(
								NewNode("great-grandson1-pretender", "value1", 42, nil))))),
						NewNode("child2", "value2", 42, NewChildren(
							NewNode("grandson1", "value1", 42, nil))),
						NewNode("child3", "value3", 42, nil)))
					tree2 = NewNode("parent", "value", 42, NewChildren(
						NewNode("child1", "value1", 42, NewChildren(
							NewNode("grandson1", "value1", 42, NewChildren(
								NewNode("great-grandson1", "value1", 42, nil))))),
						NewNode("child2", "value2", 42, NewChildren(
							NewNode("grandson1", "value1", 42, nil))),
						NewNode("child3", "value3", 42, nil)))
				})
				It("returns no collisions", func() {
					collisions, err := FindCollisions(tree1, tree2)
					Expect(err).ShouldNot(HaveOccurred())
					Expect(len(collisions)).To(Equal(2))
				})
			})
		})
	})

	Describe("FindCollisionsInItems", func() {
		var (
			list1, list2 ComparableNamedItems
		)
		Describe("equal lists", func() {
			Describe("child free", func() {
				BeforeEach(func() {
					list1 = NewNodeList(NewNode("name", "value", 42, nil)).ToComparableNamedItems()
					list2 = NewNodeList(NewNode("name", "value", 42, nil)).ToComparableNamedItems()
				})
				It("returns no collisions", func() {
					collisions, err := FindCollisionsInItems(list1, list2)
					Expect(err).ShouldNot(HaveOccurred())
					Expect(len(collisions)).To(Equal(0))
				})
			})
		})

		Describe("lists with collisions", func() {
			Describe("one different value", func() {
				BeforeEach(func() {
					list1 = NewNodeList(NewNode("node-name", "value", 42, nil)).ToComparableNamedItems()
					list2 = NewNodeList(NewNode("node-name", "value", 43, nil)).ToComparableNamedItems()
				})
				It("returns collisions", func() {
					collisions, err := FindCollisionsInItems(list1, list2)
					Expect(err).ShouldNot(HaveOccurred())
					Expect(len(collisions)).To(Equal(1))
					Expect(len(collisions[0].Mismatches)).To(Equal(1))
				})
			})
			Describe("several different values", func() {
				BeforeEach(func() {
					list1 = NewNodeList(NewNode("node-name", "value1", 42, nil)).ToComparableNamedItems()
					list2 = NewNodeList(NewNode("node-name", "value2", 43, nil)).ToComparableNamedItems()
				})
				It("returns collisions", func() {
					collisions, err := FindCollisionsInItems(list1, list2)
					Expect(err).ShouldNot(HaveOccurred())
					Expect(len(collisions)).To(Equal(1))
					Expect(len(collisions[0].Mismatches)).To(Equal(2))
				})
			})
			Describe("different names", func() {
				BeforeEach(func() {
					list1 = NewNodeList(NewNode("node-name-1", "value1", 42, nil)).ToComparableNamedItems()
					list2 = NewNodeList(NewNode("node-name-2", "value2", 43, nil)).ToComparableNamedItems()
				})
				It("returns collisions", func() {
					collisions, err := FindCollisionsInItems(list1, list2)
					Expect(err).ShouldNot(HaveOccurred())
					Expect(len(collisions)).To(Equal(2))
					spew.Dump(collisions)
				})
			})

		})

	})
})
