package space

type Planet string

const (
	eYear         = 365.25
	sec2DayDivFac = 86400
)

func Age(seconds float64, planet Planet) float64 {
	switch planet {
	case "Mercury":
		return CalculateRelativeAge(seconds, 0.2408467)
	case "Venus":
		return CalculateRelativeAge(seconds, 0.61519726)
	case "Earth":
		return CalculateRelativeAge(seconds, 1.0)
	case "Mars":
		return CalculateRelativeAge(seconds, 1.8808158)
	case "Jupiter":
		return CalculateRelativeAge(seconds, 11.862615)
	case "Saturn":
		return CalculateRelativeAge(seconds, 29.447498)
	case "Uranus":
		return CalculateRelativeAge(seconds, 84.016846)
	case "Neptune":
		return CalculateRelativeAge(seconds, 164.79132)
	default: // Earth
		return -1.0
	}
}

func CalculateRelativeAge(sec float64, orbPeriod float64) float64 {
	return ((sec / sec2DayDivFac) / (eYear * orbPeriod))
}
