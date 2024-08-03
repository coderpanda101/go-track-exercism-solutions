package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgPreparationTimeInMin int) int {
	if avgPreparationTimeInMin == 0 {
		avgPreparationTimeInMin = 2
	}

	return len(layers) * avgPreparationTimeInMin
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {

	qtyNoodles := 0
	qtySauce := 0.0

	for i := 0; i < len(layers); i++ {
		if layers[i] == "noodles" {
			qtyNoodles += 50
		}
		if layers[i] == "sauce" {
			qtySauce += 0.2
		}
	}

	return qtyNoodles, qtySauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(frndList []string, myList []string) {
	(myList)[len(myList)-1] = (frndList)[len(frndList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portionsToCook int) []float64 {

	amtNeededForNPortions := make([]float64, len(quantities))

	for i := 0; i < len(quantities); i++ {
		amtNeededFor1Portion := quantities[i] / 2
		amtNeededForNPortions[i] = amtNeededFor1Portion * float64(portionsToCook)
	}

	return amtNeededForNPortions
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
