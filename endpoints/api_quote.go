package endpoints

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	md "github.com/UniversityRadioYork/API-go/models"
	ut "github.com/UniversityRadioYork/API-go/utils"
	memcache "github.com/bradfitz/gomemcache/memcache"
	"log"
	"net/http"
)

var GetAllQuotesSQL = "SELECT quote.quote_id, member.fname, member.sname, quote.text, quote.date, quote.suspended" +
	" FROM people.quote INNER JOIN public.member ON quote.source = member.memberid ORDER BY quote.date DESC;"

var GetAllQuotesBaseKey = "SQL:GetAllQuotes"

func GetAllQuotes(w http.ResponseWriter, r *http.Request) {
	var quotes []md.Quote
	var network bytes.Buffer
	//Attempt to retrieve data from memcache server
	data, err := ut.MemcacheClient.Get(GetAllQuotesBaseKey)
	if err != nil {
		log.Printf("Cache miss")
		//Run SQL query
		rows, err := ut.Database.Query(GetAllQuotesSQL)
		if err != nil {
			log.Print(err)
			w.WriteHeader(500)
			json.NewEncoder(w).Encode(md.APIError{
				OK:        false,
				ErrorCode: "database_fail",
				ErrorInfo: nil,
			})
			return
		}
		defer rows.Close()
		//Move results into struct
		for rows.Next() {
			var quote md.Quote
			err := rows.Scan(&quote.QuoteId, &quote.FName, &quote.SName,
				&quote.Text, &quote.Date, &quote.Suspended)
			if err != nil {
				log.Print(err)
				w.WriteHeader(500)
				json.NewEncoder(w).Encode(md.APIError{
					OK:        false,
					ErrorCode: "database_fail",
					ErrorInfo: nil,
				})
				return
			}
			quotes = append(quotes, quote)
		}
		//Encode struct into bytes
		enc := gob.NewEncoder(&network)
		enc.Encode(quotes)
		ut.MemcacheClient.Set(&memcache.Item{Key: GetAllQuotesBaseKey, Value: network.Bytes()})
	} else {
		log.Printf("Cache hit")
		//Decode bytes into struct
		network.Write(data.Value)
		dec := gob.NewDecoder(&network)
		dec.Decode(&quotes)
	}
	//Encode struct into json and respond to request
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(md.APIResponse{
		OK:   true,
		Data: quotes,
	})
}
