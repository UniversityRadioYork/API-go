package utils

type Config struct {
	ServerPort	string `toml:"serverport"`
	DBHost		string `toml:"dbhost"`
	DBPort		string `toml:"dbport"`
	DBUser		string `toml:"dbuser"`
	DBPassword	string `toml:"dbpassword"`
	DBName		string `toml:"dbname"`
	SSLMode		string `toml:"sslmode"`
	Memcache	string `toml:"memcache"`
}
