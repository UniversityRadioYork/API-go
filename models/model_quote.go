package models

type Quote struct {
	QuoteId		int32 	`json:"quote_id"`
	FName		string	`json:"fname"`
	SName		string	`json:"sname"`
	Text		string	`json:"text"`
	Date		string	`json:"date"`
	Suspended	bool	`json:"suspended"`
}