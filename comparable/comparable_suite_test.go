package comparable_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestComparable(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Comparable Suite")
}
