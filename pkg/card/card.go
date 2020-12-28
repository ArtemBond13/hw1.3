// card представляет собой абстракцию банковской карты
package card

import (
	"time"
)

type Card struct {
	Id           int64
	Issuer       string
	Balance      int64
	Currency     string
	Number       string
	Icon         string
	Transactions []Transaction
}

type Transaction struct {
	Id     int64
	Amount int64
	Time   time.Time
	MCC    string
	Status string
}

// AddTrancsation добавляет транзакцию в историю карты
func AddTrancsation(card *Card, transaction *Transaction) {
	card.Transactions = append(card.Transactions, *transaction)

}

// SumByMCC расчет транзакций по коду МСС
func SumByMCC(transactions []Transaction, mcc []string) int64 {
	sum := int64(0)
	for _, valueMCC := range mcc {
		for _, val := range transactions {
			if valueMCC == val.MCC {
				sum += val.Amount
			}
		}

	}
	//for i := 0; i < len(mcc); i++ {
	//	for j := 0; j < len(transactions); j++ {
	//		if mcc[i] == transactions[j].MCC {
	//			sum += transactions[j].Amount
	//		}
	//	}
	//
	//}
	return sum
}
