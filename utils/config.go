package utils

type Config struct {
	Host		string `toml:"host"`
	Port		string `toml:"port"`
	User		string `toml:"user"`
	Password	string `toml:"password"`
	DBName		string `toml:"dbname"`
	SSLMode		string `toml:"sslmode"`
	Memcache	string `toml:"memcache"`
}
