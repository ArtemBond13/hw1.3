package main

import (
	"fmt"
	"hw1.3/pkg/card"
	"time"
)

func main() {
	master := card.Card{
		Id:       0001,
		Issuer:   "MasterCard",
		Balance:  500000,
		Currency: "RUB",
		Number:   "3453 5345 3445 3552",
		Icon:     "http://...",
		Transactions: []card.Transaction{{
			Id:     1,
			Amount: 73555,
			Time:   time.Now(),
			MCC:    "5411",
			Status: " ",
		},
			{
				Id:     2,
				Amount: 120391,
				Time:   time.Now(),
				MCC:    "5411",
				Status: " ",
			}},
	}
	fmt.Println(&master)
	card.AddTrancsation(&master, &card.Transaction{Id: 3, Time: time.Now(), Amount: 1000, MCC: "5411"})
	fmt.Println(&master)

	mcc := []string{
		"4511",
		"4512",
		"3423",
		"5411",
	}

	//fmt.Println(mcc)
	totalResult := card.SumByMCC(master.Transactions, mcc)
	fmt.Printf("Итого трат за месяц %d", totalResult)
}
