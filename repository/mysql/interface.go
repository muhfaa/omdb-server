package mysqldb

type Repository interface {
	SaveSearchActivity(searchWord string) error
}
