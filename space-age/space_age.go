package space

type Planet string

func Age(seconds float64, planet Planet) float64{
	planetsToRatios := map[Planet]float64{
		"Earth" : 1.0,
		"Mercury" : 0.2408467, 
		"Venus" : 0.61519726,
		"Mars" : 1.8808158,
		"Jupiter" : 11.862615,
		"Saturn" : 29.447498,
		"Uranus" : 84.016846,
		"Neptune" : 164.79132,
		}
	ratio, ok := planetsToRatios[planet]
	if !ok{
		return -1.0
	}
	return (seconds /( 3600* 24* 365.25)) / ratio
}