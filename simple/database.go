package simple

type Database struct {
	Name string
}

type DatabasePostgreSQL Database
type DatabaseMongoDB Database

func NewDatabasePostgre() *DatabasePostgreSQL {
	return (*DatabasePostgreSQL)(&Database{Name: "PostgreSQL"})
}
func NewDatabaseMongoDB() *DatabaseMongoDB {
	return (*DatabaseMongoDB)(&Database{Name: "MongoDB"})
}

type DatabaseRepo struct {
	DatabasePostgreSQL *DatabasePostgreSQL
	DatabaseMongoDB    *DatabaseMongoDB
}

func NewDatabaseRepo(postgresql *DatabasePostgreSQL,
	mongodb *DatabaseMongoDB) *DatabaseRepo {
	return &DatabaseRepo{DatabasePostgreSQL: postgresql, DatabaseMongoDB: mongodb}
}
