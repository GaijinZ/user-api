package config

func GetPostgresConnectionString() string {
	database := "host=postgres user=userapi password=userapi dbname=userapi port=5432 sslmode=disable"
	return database
}
