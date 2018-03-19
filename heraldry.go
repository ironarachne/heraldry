package heraldry

import (
	"fmt"
	"github.com/ajstarks/svgo"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Tincture struct {
	Type    string
	Name    string
	Hexcode string
}

type Charge struct {
	Name string
	Tincture
}

type Division struct {
	Name   string
	Blazon string
	Tincture
}

type Field struct {
	Division
	Tincture
}

type Device struct {
	Field
	Charges []Charge
}

var (
	Or                 = Tincture{"metal", "Or", "#FFEC00"}
	Argent             = Tincture{"metal", "argent", "#FFFFFF"}
	Sable              = Tincture{"color", "sable", "#000000"}
	Gules              = Tincture{"color", "gules", "#EE0000"}
	Vert               = Tincture{"color", "vert", "#009900"}
	Azure              = Tincture{"color", "azure", "#0000FF"}
	Purpure            = Tincture{"color", "purpure", "#831CA6"}
	Metals             = [2]Tincture{Or, Argent}
	Colors             = [5]Tincture{Sable, Gules, Vert, Azure, Purpure}
	AvailableTinctures = [7]Tincture{Or, Argent, Sable, Gules, Vert, Azure, Purpure}

	Bend               = Division{"bend", "Per bend ", Tincture{}}
	BendSinister       = Division{"bendsinister", "Per bend sinister ", Tincture{}}
	Fess               = Division{"fess", "Per fess ", Tincture{}}
	Pale               = Division{"pale", "Per pale ", Tincture{}}
	Plain              = Division{"plain", "", Tincture{}}
	Quarterly          = Division{"quarterly", "Quarterly ", Tincture{}}
	Saltire            = Division{"saltire", "Per saltire ", Tincture{}}
	Chevron            = Division{"chevron", "Per chevron ", Tincture{}}
	AvailableDivisions = [8]Division{Bend, BendSinister, Fess, Pale, Plain, Quarterly, Saltire, Chevron}

	OrdinaryPale     = Charge{"pale", Tincture{}}
	OrdinaryFess     = Charge{"fess", Tincture{}}
	AvailableCharges = [2]Charge{OrdinaryPale, OrdinaryFess}
)

func randomCharge() Charge {
	rand.Seed(time.Now().UnixNano())
	return AvailableCharges[rand.Intn(len(AvailableCharges))]
}

func randomDivision() Division {
	rand.Seed(time.Now().UnixNano())
	return AvailableDivisions[rand.Intn(len(AvailableDivisions))]
}

func randomTincture() Tincture {
	rand.Seed(time.Now().UnixNano())
	return AvailableTinctures[rand.Intn(len(AvailableTinctures))]
}

func randomTinctureColor() Tincture {
	rand.Seed(time.Now().UnixNano())
	t := Colors[rand.Intn(len(Colors))]
	return t
}

func randomTinctureMetal() Tincture {
	rand.Seed(time.Now().UnixNano())
	t := Metals[rand.Intn(len(Metals))]
	return t
}

func randomComplementaryTincture(t Tincture) Tincture {
	rand.Seed(time.Now().UnixNano())

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
	rand.Seed(time.Now().UnixNano())
	t2 := Tincture{}
	if t.Type == "metal" {
		t2 = randomTinctureColor()
	} else {
		t2 = randomTinctureMetal()
	}

	return t2
}

func shallWeIncludeCharges() bool {
	rand.Seed(time.Now().UnixNano())
	someRandomValue := rand.Intn(10)
	if someRandomValue > 2 {
		return true
	}
	return false
}

func Generate() Device {
	fieldTincture1 := randomTincture()
	fieldTincture2 := randomComplementaryTincture(fieldTincture1)
	chargeTincture := randomContrastingTincture(fieldTincture1)

	division := randomDivision()
	division.Tincture = fieldTincture2

	var charges []Charge

	if shallWeIncludeCharges() {
		charge := randomCharge()
		charge.Tincture = chargeTincture
		charges = append(charges, charge)
	}

	f := Field{Division: division, Tincture: fieldTincture1}
	d := Device{Field: f, Charges: charges}

	return d
}

func (d Device) Blazon() string {
	blazon := ""
	if d.Field.Division.Name == "plain" {
		blazon = strings.Title(d.Field.Tincture.Name)
	} else {
		blazon = d.Field.Division.Blazon
		blazon += d.Field.Tincture.Name
	}

	if d.Field.Division.Name != "plain" {
		blazon += " and " + d.Field.Division.Tincture.Name
	}

	if len(d.Charges) > 0 {
		blazon += ", a " + d.Charges[0].Name + " " + d.Charges[0].Tincture.Name
	}

	return blazon
}

func (d Device) Svg(fileName string) {
	width := 320
	height := 420
	var writer io.WriteCloser
	writer, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	canvas := svg.New(writer)
	canvas.Start(width, height, "mask='url(#shieldmask)'")
	canvas.Def()
	canvas.Mask("shieldmask", 0, 0, width, height)
	canvas.Path("m10.273 21.598v151.22c0 96.872 89.031 194.34 146.44 240.09 57.414-45.758 146.44-143.22 146.44-240.09v-151.22h-292.89z", "fill:#FFFFFF")
	canvas.MaskEnd()
	canvas.DefEnd()
	canvas.Group()
	canvas.Rect(0, 0, width, height, "fill:"+d.Field.Tincture.Hexcode)
	switch d.Field.Division.Name {
	case "plain":
	case "pale":
		canvas.Rect(int(width/2), 0, int(width/2), height, "fill:"+d.Field.Division.Tincture.Hexcode)
	case "fess":
		canvas.Rect(0, 0, width, int(height/2), "fill:"+d.Field.Division.Tincture.Hexcode)
	case "bend":
		canvas.Polygon([]int{0, 0, width}, []int{0, height, height}, "fill:"+d.Field.Division.Tincture.Hexcode)
	case "bendsinister":
		canvas.Polygon([]int{0, width, 0}, []int{0, 0, height}, "fill:"+d.Field.Division.Tincture.Hexcode)
	case "chevron":
		canvas.Polygon([]int{0, int(width / 2), width}, []int{height, int(height / 2), height}, "fill:"+d.Field.Division.Tincture.Hexcode)
	case "quarterly":
		canvas.Rect(int(width/2), 0, int(width/2), int(height/2), "fill:"+d.Field.Division.Tincture.Hexcode)
		canvas.Rect(0, int(height/2), int(width/2), int(height/2), "fill:"+d.Field.Division.Tincture.Hexcode)
	case "saltire":
		canvas.Polygon([]int{0, int(width / 2), 0}, []int{0, int(height / 2), height}, "fill:"+d.Field.Division.Tincture.Hexcode)
		canvas.Polygon([]int{width, int(width / 2), width}, []int{0, int(height / 2), height}, "fill:"+d.Field.Division.Tincture.Hexcode)
	}
	for _, charge := range d.Charges {
		switch charge.Name {
		case "pale":
			canvas.Rect(int(width/3), 0, int(width/3), height, "fill:"+charge.Tincture.Hexcode)
		case "fess":
			canvas.Rect(0, int(height/3), width, int(height/3), "fill:"+charge.Tincture.Hexcode)
		}
	}
	canvas.Path("m10.273 21.598v151.22c0 96.872 89.031 194.34 146.44 240.09 57.414-45.758 146.44-143.22 146.44-240.09v-151.22h-292.89z", "stroke:#000000;stroke-width:4;fill:none")
	canvas.Gend()
	canvas.End()

	defer writer.Close()
}
