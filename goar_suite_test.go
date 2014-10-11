package goar

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestGoar(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Goar Suite")
}
