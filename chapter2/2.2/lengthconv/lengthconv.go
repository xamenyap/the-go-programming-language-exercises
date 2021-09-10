package lengthconv

type Feet float64
type Meter float64

func FeetToMeter(f Feet) Meter {
	return Meter(f / 3.28)
}

func MeterToFeet(m Meter) Feet {
	return Feet(m * 3.28)
}
