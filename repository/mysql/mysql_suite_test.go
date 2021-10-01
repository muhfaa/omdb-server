package mysqldb_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMysqldb(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Mysqldb Suite")
}
