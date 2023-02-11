package custom

type CoffeePot string

func (c CoffeePot) String() string {
	return string(c) + "coffee pot"
}
