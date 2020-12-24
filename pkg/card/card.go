package card

import "time"

type Transaction struct {
	Id          int64
	Amount      int64
	Time        time.Time
	MCC         string
	Status      string
	transaction []*Transaction
}
