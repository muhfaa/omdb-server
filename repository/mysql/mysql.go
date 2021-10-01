package mysqldb

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/muhfaa/omdb-server/repository"
)

type MysqlDBDSL struct {
	Username string
	Password string
	Host     string
	Port     string
	DBName   string
}

func (m MysqlDBDSL) GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", m.Username, m.Password, m.Host, m.Port, m.DBName)
}

type MysqlDBImpl struct {
	dbInstance repository.DB
}

func NewMysqlDB(dbInstance repository.DB) MysqlDBImpl {
	return MysqlDBImpl{
		dbInstance: dbInstance,
	}
}

func (m MysqlDBImpl) SaveSearchActivity(searchWord string) error {
	_, err := m.dbInstance.Query("INSERT INTO search_activities (search_word, created_at) VALUES ( ?, ? )", searchWord, time.Now())
	if err != nil {
		return err
	}
	return nil
}
