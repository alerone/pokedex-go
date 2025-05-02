package pokeapi

import "math/rand/v2"


const minBaseExp = 1.0
const maxBaseExp = 635.0
func getCatchRate(baseExp, maxIntents int) float64 {
	return ((float64(baseExp) - minBaseExp) / (maxBaseExp - minBaseExp)) * float64(maxIntents)
}

func tryCatch(pokemon Pokemon, maxIntents int) bool {
	catchRes := rand.Float64() * float64(maxIntents)
	catchRate := getCatchRate(pokemon.BaseExperience, maxIntents)
	return catchRes >= catchRate
}



