package api

type Resolver struct {
	DB *DB
	ES *ES
}

type Person struct {
	ID   string
	Name string
}

type PersonResolver struct {
	Person Person
}
