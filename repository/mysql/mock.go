package mysqldb

type MockMysqlDB struct {
	MockSaveSearchActivity func(searchWord string) error
}

func (m MockMysqlDB) SaveSearchActivity(searchWord string) error {
	return m.MockSaveSearchActivity(searchWord)
}

func NewMock() MockMysqlDB {
	return MockMysqlDB{
		MockSaveSearchActivity: func(searchWord string) error {
			return nil
		},
	}
}
