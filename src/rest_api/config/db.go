package config

func GetPostgresConnectionString() string {
	database := "host=192.168.33.2 user=userapi password=userapi dbname=userapi port=5432 sslmode=disable"
	return database
}
