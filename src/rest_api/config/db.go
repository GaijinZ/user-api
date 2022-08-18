package config

func GetPostgresConnectionString() string {
	database := "host=192.168.33.1 user=userapi password=userapi dbname=userapidb port=8221 sslmode=disable"
	return database
}
