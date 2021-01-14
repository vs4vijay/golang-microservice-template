package app

type IDatabase interface {
}

type Database struct {
	Config Config
}

func NewDatabase(config Config) (*Database, error) {
	database := &Database{
		Config: config,
	}

	return database, nil
}
