package heraldry

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/ajstarks/svgo"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func fetchChargePathData(charge string) string {
	pathData, err := ioutil.ReadFile("charges/" + charge)
	check(err)
	return string(pathData)
}

// RenderToBlazon renders a device as its blazon and returns it.
func RenderToBlazon(device Device) string {
	blazon := ""
	if device.Field.Division.Name == "plain" {
		blazon = strings.Title(device.Field.Tincture.Name)
	} else {
		blazon = device.Field.Division.Blazon
		blazon += device.Field.Tincture.Name
	}

	if device.Field.Division.Name != "plain" {
		blazon += " and " + device.Field.Division.Tincture.Name
	}

	if len(device.Charges) > 0 {
		blazon += ", a " + device.Charges[0].Name + " " + device.Charges[0].Tincture.Name
	}

	return blazon
}

// RenderToSvg renders a device as SVG and writes to fileName.
func RenderToSvg(device Device, fileName string, width int, height int) {
	var writer io.WriteCloser
	writer, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	centerX := int(width / 2)
	centerY := int(height / 2)

	lineColor := "#000000"

	canvas := svg.New(writer)
	canvas.Start(width, height, "mask='url(#shieldmask)'")
	canvas.Def()
	canvas.Mask("shieldmask", 0, 0, width, height)
	canvas.Path("m10.273 21.598v151.22c0 96.872 89.031 194.34 146.44 240.09 57.414-45.758 146.44-143.22 146.44-240.09v-151.22h-292.89z", "fill:#FFFFFF")
	canvas.MaskEnd()
	canvas.DefEnd()
	canvas.Group()
	canvas.Rect(0, 0, width, height, "fill:"+device.Field.Tincture.Hexcode)
	switch device.Field.Division.Name {
	case "plain":
	case "pale":
		canvas.Rect(int(width/2), 0, int(width/2), height, "fill:"+device.Field.Division.Tincture.Hexcode)
	case "fess":
		canvas.Rect(0, 0, width, int(height/2), "fill:"+device.Field.Division.Tincture.Hexcode)
	case "bend":
		canvas.Polygon([]int{0, 0, width}, []int{0, height, height}, "fill:"+device.Field.Division.Tincture.Hexcode)
	case "bendsinister":
		canvas.Polygon([]int{0, width, 0}, []int{0, 0, height}, "fill:"+device.Field.Division.Tincture.Hexcode)
	case "chevron":
		canvas.Polygon([]int{0, int(width / 2), width}, []int{height, int(height / 2), height}, "fill:"+device.Field.Division.Tincture.Hexcode)
	case "quarterly":
		canvas.Rect(int(width/2), 0, int(width/2), int(height/2), "fill:"+device.Field.Division.Tincture.Hexcode)
		canvas.Rect(0, int(height/2), int(width/2), int(height/2), "fill:"+device.Field.Division.Tincture.Hexcode)
	case "saltire":
		canvas.Polygon([]int{0, int(width / 2), 0}, []int{0, int(height / 2), height}, "fill:"+device.Field.Division.Tincture.Hexcode)
		canvas.Polygon([]int{width, int(width / 2), width}, []int{0, int(height / 2), height}, "fill:"+device.Field.Division.Tincture.Hexcode)
	}
	for _, charge := range device.Charges {
		if charge.Tincture.Name == "sable" {
			lineColor = "#FFFFFF"
		}
		switch charge.Identifier {
		case "pale":
			canvas.Rect(int(width/3), 0, int(width/3), height, "fill:"+charge.Tincture.Hexcode)
		case "fess":
			canvas.Rect(0, int(height/3), width, int(height/3), "fill:"+charge.Tincture.Hexcode)
		case "cross":
			crossHalfWidth := 40
			canvas.Polygon(
				[]int{centerX - crossHalfWidth, centerX + crossHalfWidth, centerX + crossHalfWidth, width, width, centerX + crossHalfWidth, centerX + crossHalfWidth, centerX - crossHalfWidth, centerX - crossHalfWidth, 0, 0, centerX - crossHalfWidth},
				[]int{0, 0, centerY - crossHalfWidth, centerY - crossHalfWidth, centerY + crossHalfWidth, centerY + crossHalfWidth, height, height, centerY + crossHalfWidth, centerY + crossHalfWidth, centerY - crossHalfWidth, centerY - crossHalfWidth},
				"fill:"+charge.Tincture.Hexcode)
		case "bend":
			canvas.TranslateRotate(-100, 135, -45)
			canvas.Rect(int(width/3), 0, int(width/3), height, "fill:"+charge.Tincture.Hexcode)
			canvas.Gend()
		case "saltire":
			saltireHalfWidth := 30
			canvas.Polygon(
				[]int{0, saltireHalfWidth, centerX, width - saltireHalfWidth, width, width, centerX + saltireHalfWidth, width, width, width - saltireHalfWidth, centerX, saltireHalfWidth, 0, 0, centerX - saltireHalfWidth, 0},
				[]int{0, 0, centerY - saltireHalfWidth, 0, 0, saltireHalfWidth, centerY, height - saltireHalfWidth, height, height, centerY + saltireHalfWidth, height, height, height - saltireHalfWidth, centerY, saltireHalfWidth},
				"fill:"+charge.Tincture.Hexcode)
		case "chevron":
			chevronHalfWidth := 40
			canvas.Polygon(
				[]int{0, centerX, width, width, centerX, 0},
				[]int{height - int(height/4), height - int(height/3) - (3 * chevronHalfWidth), height - int(height/4), height - int(height/4) + (2 * chevronHalfWidth), height - int(height/3) - chevronHalfWidth, height - int(height/4) + (2 * chevronHalfWidth)},
				"fill:"+charge.Tincture.Hexcode)
		case "chief":
			canvas.Rect(0, 0, width, int(height/3), "fill:"+charge.Tincture.Hexcode)
		case "pile":
			canvas.Polygon(
				[]int{0, width, int(width / 2)},
				[]int{0, 0, int(height / 2)},
				"fill:"+charge.Tincture.Hexcode)
		case "pall":
			pallHalfWidth := 40
			canvas.Polygon(
				[]int{0, pallHalfWidth, centerX, width - pallHalfWidth, width, width, centerX + pallHalfWidth - int(pallHalfWidth/3), centerX + pallHalfWidth - int(pallHalfWidth/3), centerX - pallHalfWidth + int(pallHalfWidth/3), centerX - pallHalfWidth + int(pallHalfWidth/3), 0},
				[]int{0, 0, centerY - pallHalfWidth, 0, 0, pallHalfWidth, centerY + pallHalfWidth - int(pallHalfWidth/3), height, height, centerY + pallHalfWidth - int(pallHalfWidth/3), pallHalfWidth},
				"fill:"+charge.Tincture.Hexcode)
		case "bordure":
			pathData := fetchChargePathData("bordure")
			canvas.Path(pathData, "stroke:"+charge.Tincture.Hexcode+";stroke-width:100;fill:none")
		case "lozenge":
			lozengeHalfWidth := 80
			canvas.Polygon(
				[]int{centerX, centerX + lozengeHalfWidth, centerX, centerX - lozengeHalfWidth},
				[]int{centerY - lozengeHalfWidth - int(lozengeHalfWidth/2), centerY, centerY + lozengeHalfWidth + int(lozengeHalfWidth/2), centerY},
				"fill:"+charge.Tincture.Hexcode)
		case "roundel":
			roundelRadius := 100
			canvas.Circle(
				int(width/2),
				int(height/2),
				roundelRadius,
				"fill:"+charge.Tincture.Hexcode)
		case "eagle-displayed":
			pathData0 := fetchChargePathData("eagle-displayed-0")
			pathData1 := fetchChargePathData("eagle-displayed-1")
			canvas.Translate(10, 50)
			canvas.Path(pathData0, "fill:"+charge.Tincture.Hexcode+";fill-opacity:1")
			canvas.Path(pathData1, "fill:"+lineColor)
			canvas.Gend()
		case "dragon-passant":
			pathData0 := fetchChargePathData("dragon-passant-0")
			pathData1 := fetchChargePathData("dragon-passant-1")
			canvas.Translate(10, 50)
			canvas.Path(pathData0, "fill:"+charge.Tincture.Hexcode+";fill-opacity:1")
			canvas.Path(pathData1, "fill:"+lineColor)
			canvas.Gend()
		case "gryphon-passant":
			pathData0 := fetchChargePathData("gryphon-passant-0")
			pathData1 := fetchChargePathData("gryphon-passant-1")
			canvas.Translate(10, 50)
			canvas.Path(pathData0, "fill:"+charge.Tincture.Hexcode+";fill-opacity:1")
			canvas.Path(pathData1, "fill:"+lineColor)
			canvas.Gend()
		case "fox-passant":
			pathData0 := fetchChargePathData("fox-passant-0")
			pathData1 := fetchChargePathData("fox-passant-1")
			canvas.Translate(10, 50)
			canvas.Path(pathData0, "fill:"+charge.Tincture.Hexcode+";fill-opacity:1")
			canvas.Path(pathData1, "fill:"+lineColor)
			canvas.Gend()
		}

	}
	canvas.Path("m10.273 21.598v151.22c0 96.872 89.031 194.34 146.44 240.09 57.414-45.758 146.44-143.22 146.44-240.09v-151.22h-292.89z", "stroke:#000000;stroke-width:4;fill:none")
	canvas.Gend()
	canvas.End()

	defer writer.Close()
}
