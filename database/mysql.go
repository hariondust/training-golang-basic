package database

var connection string

// gunakan init untuk menjalankan fungsi ketika package dipanggil
func init() {
	connection = "MySQL"
}

func GetDatabase() string {
	return connection
}
