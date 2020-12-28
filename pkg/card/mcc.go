package card

func TranslateMCC(code string) string {
	// представим, что mcc читается из файла (научимся позже)
	mcc := map[string]string{
		// TODO: ваш код (необходимые вам MCC-коды)
		"4013": "Men's and Boys' Clothing",
		"4123": "Women's Ready to Wear Stories",
		"4156": "Women's Accessory and Special Stores",
		"4978": "Children's and Infants' Wear Stores",
		"5411": "Family Clothing Stores",
	}
	var searchResults = "Категория не указана"
	for key, value := range mcc {
		if code == key {
			searchResults = value
		}
	}

	// TODO: ваш код
	return searchResults
}
