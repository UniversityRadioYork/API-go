package utils

import (
	"fmt"
	"log"
)

var getAPIKey = "SELECT revoked FROM myury.api_key WHERE key_string=%s"

// IsAPIKeyValid retrieves revoked state of an API key.
func IsAPIKeyValid(userAPIKey string) bool {
	dbQuery := fmt.Sprintf(getAPIKey, userAPIKey)
	rows, err := Database.Query(dbQuery)
	if err != nil {
		log.Print(err)
	}
	defer rows.Close()
	var isRevoked bool
	err = rows.Scan(&isRevoked)
	if err != nil {
		log.Print(err)
	}
	return isRevoked
}
