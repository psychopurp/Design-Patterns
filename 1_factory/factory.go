package factory

import "fmt"

type (
	mongoDB struct {
		database map[string]string
	}

	sqlite struct {
		database map[string]string
	}

	Database interface {
		GetData(string) string
		PutData(string, string)
	}
)

func (mdb *mongoDB) GetData(query string) string {
	if val, ok := mdb.database[query]; ok {
		fmt.Println("MongoDB")
		return val
	} else {
		return ""
	}
}

func (mdb *mongoDB) PutData(query string, data string) {
	mdb.database[query] = data
}

func (sql *sqlite) GetData(query string) string {
	if val, ok := sql.database[query]; ok {
		fmt.Println("sqlite")
		return val
	} else {
		return ""
	}
}

func (sql *sqlite) PutData(query string, data string) {
	sql.database[query] = data
}

func DatabaseFactory(env string) Database {
	switch env {
	case "production":
		return &mongoDB{database: make(map[string]string)}
	case "development":
		return &sqlite{database: make(map[string]string)}
	default:
		return nil
	}
}
