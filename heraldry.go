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
	Or        = Tincture{"metal", "Or", "#FFFF00"}
	Argent    = Tincture{"metal", "argent", "#FFFFFF"}
	Sable     = Tincture{"color", "sable", "#000000"}
	Gules     = Tincture{"color", "gules", "#EE0000"}
	Vert      = Tincture{"color", "vert", "#009900"}
	Azure     = Tincture{"color", "azure", "#0000FF"}
	Purpure   = Tincture{"color", "purpure", "#CC00CC"}
	Metals    = [2]Tincture{Or, Argent}
	Colors    = [5]Tincture{Sable, Gules, Vert, Azure, Purpure}
	Tinctures = [7]Tincture{Or, Argent, Sable, Gules, Vert, Azure, Purpure}

	Bend         = Division{"bend", "Per bend ", Tincture{}}
	BendSinister = Division{"bendsinister", "Per bend sinister ", Tincture{}}
	Fess         = Division{"fess", "Per fess ", Tincture{}}
	Pale         = Division{"pale", "Per pale ", Tincture{}}
	Plain        = Division{"plain", "", Tincture{}}
	Quarterly    = Division{"quarterly", "Quarterly ", Tincture{}}
	Saltire      = Division{"saltire", "Per saltire ", Tincture{}}
	Chevron      = Division{"chevron", "Per chevron ", Tincture{}}
	Divisions    = [8]Division{Bend, BendSinister, Fess, Pale, Plain, Quarterly, Saltire, Chevron}
)

func randomDivision() Division {
	rand.Seed(time.Now().UnixNano())
	return Divisions[rand.Intn(len(Divisions))]
}

func randomTincture() Tincture {
	rand.Seed(time.Now().UnixNano())
	return Tinctures[rand.Intn(len(Tinctures))]
}

func randomContrastingTincture(t Tincture) Tincture {
	rand.Seed(time.Now().UnixNano())
	t2 := Tincture{}
	if t.Type == "metal" {
		t2 = Colors[rand.Intn(len(Colors))]
	} else {
		t2 = Metals[rand.Intn(len(Metals))]
	}

	return t2
}

func Generate() Device {
	fieldTincture1 := randomTincture()
	fieldTincture2 := randomContrastingTincture(fieldTincture1)

	division := randomDivision()
	division.Tincture = fieldTincture2

	f := Field{Division: division, Tincture: fieldTincture1}
	d := Device{Field: f}

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
	canvas.Path("m10.273 21.598v151.22c0 96.872 89.031 194.34 146.44 240.09 57.414-45.758 146.44-143.22 146.44-240.09v-151.22h-292.89z", "stroke:#000000;stroke-width:4;fill:none")
	canvas.Gend()
	canvas.End()

	defer writer.Close()
}
