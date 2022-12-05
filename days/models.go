package days

type elf struct {
	meals               []int
	totalCalories       int
	cleaningSectorStart int
	cleaningSectorEnd   int
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

type CrateInstructions struct {
	NumberOfCrates   int
	OriginStack      int
	DestinationStack int
}
