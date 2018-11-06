package heraldry

import (
	"math/rand"
)

// Tincture is a tincture
type Tincture struct {
	Type    string
	Name    string
	Hexcode string
}

// Charge is a charge
type Charge struct {
	Identifier string
	Name       string
	Noun       string
	NounPlural string
	Descriptor string
	Article    string
	SingleOnly bool
}

// ChargeGroup is a group of charges with a common tincture
type ChargeGroup struct {
	Charges []Charge
	Tincture
}

// Division is a division of the field
type Division struct {
	Name   string
	Blazon string
	Tincture
}

// Field is the field of a coat of arms
type Field struct {
	Division
	Tincture
}

// Device is the entire coat of arms
type Device struct {
	Field
	ChargeGroups []ChargeGroup
}

func randomCharge() Charge {
	return AvailableCharges[rand.Intn(len(AvailableCharges))]
}

func randomDivision() Division {
	return AvailableDivisions[rand.Intn(len(AvailableDivisions))]
}

func randomTincture() Tincture {
	return AvailableTinctures[rand.Intn(len(AvailableTinctures))]
}

func randomTinctureColor() Tincture {
	t := Colors[rand.Intn(len(Colors))]
	return t
}

func randomTinctureMetal() Tincture {
	t := Metals[rand.Intn(len(Metals))]
	return t
}

func randomComplementaryTincture(t Tincture) Tincture {
	var availableTinctures []Tincture
	if t.Type == "color" {
		for _, color := range Colors {
			if color.Name != t.Name {
				availableTinctures = append(availableTinctures, color)
			}
		}
	} else {
		for _, metal := range Metals {
			if metal.Name != t.Name {
				availableTinctures = append(availableTinctures, metal)
			}
		}
	}
	t2 := availableTinctures[rand.Intn(len(availableTinctures))]

	return t2
}

func randomContrastingTincture(t Tincture) Tincture {
	t2 := Tincture{}
	if t.Type == "metal" {
		t2 = randomTinctureColor()
	} else {
		t2 = randomTinctureMetal()
	}

	return t2
}

func shallWeIncludeCharges() bool {
	someRandomValue := rand.Intn(10)
	if someRandomValue > 2 {
		return true
	}
	return false
}

// Generate procedurally generates a random heraldic device and returns it.
func Generate() Device {
	fieldTincture1 := randomTincture()
	fieldTincture2 := randomComplementaryTincture(fieldTincture1)
	chargeTincture := randomContrastingTincture(fieldTincture1)

	fieldHasContrastingTinctures := false

	if rand.Intn(10) > 1 {
		fieldHasContrastingTinctures = true
	}

	if fieldHasContrastingTinctures {
		fieldTincture2 = randomContrastingTincture(fieldTincture1)
	}

	division := randomDivision()
	division.Tincture = fieldTincture2

	var charges []Charge
	var chargeGroups []ChargeGroup

	if shallWeIncludeCharges() {
		charge := randomCharge()

		chargeCountRanking := rand.Intn(10)
		countOfCharges := 1
		if chargeCountRanking >= 9 && !charge.SingleOnly {
			countOfCharges = 3
		} else if chargeCountRanking >= 7 && chargeCountRanking < 9 && !charge.SingleOnly {
			countOfCharges = 2
		}
		for i := 0; i < countOfCharges; i++ {
			charges = append(charges, charge)
		}
		chargeGroup := ChargeGroup{charges, chargeTincture}
		chargeGroups = append(chargeGroups, chargeGroup)
	}

	f := Field{Division: division, Tincture: fieldTincture1}
	d := Device{Field: f, ChargeGroups: chargeGroups}

	return d
}
