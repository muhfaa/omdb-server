package grpccontroller_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGrpcservice(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Grpccontroller Suite")
}
