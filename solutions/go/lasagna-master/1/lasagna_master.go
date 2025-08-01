package lasagna

func PreparationTime(layers []string, timePerLayer int) int {
	var defaultTimePerLayer int = 2
	if timePerLayer == 0 {
		timePerLayer = defaultTimePerLayer
	}
	return timePerLayer * len(layers)
}

func Quantities(layers []string) (noodlesQuantity int, sauceQuantity float64) {
	for _, layer := range layers {
		if layer == "noodles" {
			noodlesQuantity += 50
		}
		if layer == "sauce" {
			sauceQuantity += 0.2
		}
	}
	return noodlesQuantity, sauceQuantity
}

func AddSecretIngredient(list []string, list2 []string) {
	list2[len(list2)-1] = list[len(list)-1]
}

func ScaleRecipe(quantities []float64, scale int) []float64 {
	scaledQuantities := make([]float64, len(quantities))
	const originalRecipePortions = 2
	for i, val := range quantities {
		scaledQuantities[i] = val * (float64(scale) / originalRecipePortions)
	}
	return scaledQuantities
}