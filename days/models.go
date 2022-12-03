package days

type elf struct {
	meals         []int
	totalCalories int
}

func (e *elf) calcTotalCalories() {
	for _, meal := range e.meals {
		e.totalCalories += meal
	}
}

type rucksack struct {
	Compartment1 string
	Compartment2 string
	Items        string
}
