package billing

func CalcBasePrice(item string) int {

	type Item struct {
		name  string
		price int
	}

	Butter := Item{"Butter", 4}
	Milk := Item{"Milk", 3}

	if item == Butter.name {
		return int(Butter.price)
	} else if item == Milk.name {
		return Milk.price
	} else {
		return -1
	}

}
