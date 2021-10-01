package mysqldb_test

import (
	"database/sql"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/muhfaa/omdb-server/repository"
	mysqlrepo "github.com/muhfaa/omdb-server/repository/mysql"
)

var _ = Describe("Impl", func() {

	Describe("MysqlDBDSL", func() {
		Context("GetDSN", func() {
			It("correct", func() {
				dsn := mysqlrepo.MysqlDBDSL{
					Username: "jajak",
					Password: "dah",
					Host:     "localhost",
					Port:     "3306",
					DBName:   "jajak",
				}

				str := dsn.GetDSN()
				Expect(str).To(Equal("jajak:dah@tcp(localhost:3306)/jajak"))
			})
		})
	})

	Describe("MysqlDBImpl", func() {

		db := repository.NewMockDB()

		Context("no error", func() {
			It("ok", func() {
				db.MockQuery = func(query string, args ...interface{}) (*sql.Rows, error) {
					return &sql.Rows{}, nil
				}
				op := mysqlrepo.NewMysqlDB(db)

				err := op.SaveSearchActivity("Pontianak")
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("any error", func() {
			It("error", func() {
				db.MockQuery = func(query string, args ...interface{}) (*sql.Rows, error) {
					return nil, sql.ErrNoRows
				}
				op := mysqlrepo.NewMysqlDB(db)

				err := op.SaveSearchActivity("Pontianak")
				Expect(err).To(HaveOccurred())
			})
		})
	})

})
