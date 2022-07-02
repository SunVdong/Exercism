package lasagna

// 'PreparationTime()' function
func PreparationTime(layers []string, minutes int) int {
	if minutes == 0 {
		minutes = 2
	}
	return len(layers) * minutes
}

// 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
	var n int
	var s float64
	for _, layer := range layers {
		switch layer {
		case "noodles":
			n += 50
		case "sauce":
			s += 0.2
		}
	}

	return n, s

}

// 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, amounts int) []float64 {
	var newQuantities []float64
	for _, quantity := range quantities {
		newQuantities = append(newQuantities, quantity/2*float64(amounts))
	}
	return newQuantities
}
