package comparator_test

import (
	. "github.com/Altoros/cf-audit/comparator"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type FakeComparable []int

var _ = Describe("Comparable", func() {
	var (
		firstComparable  Comparable
		secondComparable Comparable
	)

	BeforeEach(func() {
		// secondComparable =
	})

	Describe("Getting collisions", func() {
		It("returns 404", func() {

			Expect(statusCode).To(Equal(404))
		})
	})

})
