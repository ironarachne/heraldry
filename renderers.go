package heraldry

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/ajstarks/svgo"
)

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
			renderBordureToSvg(canvas, charge.Tincture.Hexcode)
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
			renderEagleDisplayedToSvg(canvas, charge.Tincture.Hexcode, lineColor)
		case "dragon-passant":
			renderDragonPassantToSvg(canvas, charge.Tincture.Hexcode, lineColor)
		case "gryphon-passant":
			renderGryphonPassantToSvg(canvas, charge.Tincture.Hexcode, lineColor)
		case "fox-passant":
			renderFoxPassantToSvg(canvas, charge.Tincture.Hexcode, lineColor)
		case "antelope-passant":
			renderAntelopePassantToSvg(canvas, charge.Tincture.Hexcode, lineColor)
		case "antelope-rampant":
			renderAntelopeRampantToSvg(canvas, charge.Tincture.Hexcode, lineColor)
		case "bat-volant":
			renderBatVolantToSvg(canvas, charge.Tincture.Hexcode, lineColor)
		case "battleaxe":
			renderBattleaxeToSvg(canvas, charge.Tincture.Hexcode, lineColor)
		case "bear-head-couped":
			renderBearHeadCoupedToSvg(canvas, charge.Tincture.Hexcode, lineColor)
		case "bear-rampant":
			renderBearRampantToSvg(canvas, charge.Tincture.Hexcode, lineColor)
		case "bear-statant":
			renderBearStatantToSvg(canvas, charge.Tincture.Hexcode, lineColor)
		case "bee-volant":
			renderBeeVolantToSvg(canvas, charge.Tincture.Hexcode, lineColor)
		case "bell":
			renderBellToSvg(canvas, charge.Tincture.Hexcode, lineColor)
		case "boar-head-erased":
			renderBoarHeadErasedToSvg(canvas, charge.Tincture.Hexcode, lineColor)
		case "boar-passant":
			renderBoarPassantToSvg(canvas, charge.Tincture.Hexcode, lineColor)
		case "boar-rampant":
			renderBoarRampantToSvg(canvas, charge.Tincture.Hexcode, lineColor)
		case "bugle-horn":
			renderBugleHornToSvg(canvas, charge.Tincture.Hexcode, lineColor)
		case "bull-passant":
			renderBullPassantToSvg(canvas, charge.Tincture.Hexcode, lineColor)
		case "bull-rampant":
			renderBullRampantToSvg(canvas, charge.Tincture.Hexcode, lineColor)
		case "castle":
			renderCastleToSvg(canvas, charge.Tincture.Hexcode, lineColor)
		case "cock":
			renderCockToSvg(canvas, charge.Tincture.Hexcode, lineColor)
		case "cockatrice":
			renderCockatriceToSvg(canvas, charge.Tincture.Hexcode, lineColor)
		case "crown":
			renderCrownToSvg(canvas, charge.Tincture.Hexcode, lineColor)
		case "dolphin-hauriant":
			renderDolphinHauriantToSvg(canvas, charge.Tincture.Hexcode, lineColor)
		case "double-headed-eagle-displayed":
			renderDoubleHeadedEagleDisplayedToSvg(canvas, charge.Tincture.Hexcode, lineColor)
		case "dragon-rampant":
			renderDragonRampantToSvg(canvas, charge.Tincture.Hexcode, lineColor)
		case "eagles-head-erased":
			renderEaglesHeadErasedToSvg(canvas, charge.Tincture.Hexcode, lineColor)
		case "fox-sejant":
			renderFoxSejantToSvg(canvas, charge.Tincture.Hexcode, lineColor)
		case "gryphon-segreant":
			renderGryphonSegreantToSvg(canvas, charge.Tincture.Hexcode, lineColor)
		case "hare-salient":
			renderHareSalientToSvg(canvas, charge.Tincture.Hexcode, lineColor)
		case "hare":
			renderHareToSvg(canvas, charge.Tincture.Hexcode, lineColor)
		case "heron":
			renderHeronToSvg(canvas, charge.Tincture.Hexcode, lineColor)
		case "horse-passant":
			renderHorsePassantToSvg(canvas, charge.Tincture.Hexcode, lineColor)
		case "horse-rampant":
			renderHorseRampantToSvg(canvas, charge.Tincture.Hexcode, lineColor)
		case "leopard-passant":
			renderLeopardPassantToSvg(canvas, charge.Tincture.Hexcode, lineColor)
		case "lion-passant":
			renderLionPassantToSvg(canvas, charge.Tincture.Hexcode, lineColor)
		case "lion-rampant":
			renderLionRampantToSvg(canvas, charge.Tincture.Hexcode, lineColor)
		case "lions-head-erased":
			renderLionsHeadErasedToSvg(canvas, charge.Tincture.Hexcode, lineColor)
		case "owl":
			renderOwlToSvg(canvas, charge.Tincture.Hexcode, lineColor)
		case "pegasus-passant":
			renderPegasusPassantToSvg(canvas, charge.Tincture.Hexcode, lineColor)
		case "pegasus-rampant":
			renderPegasusRampantToSvg(canvas, charge.Tincture.Hexcode, lineColor)
		case "ram-rampant":
			renderRamRampantToSvg(canvas, charge.Tincture.Hexcode, lineColor)
		case "ram-statant":
			renderRamStatantToSvg(canvas, charge.Tincture.Hexcode, lineColor)
		case "rose":
			renderRoseToSvg(canvas, charge.Tincture.Hexcode, lineColor)
		case "sea-horse":
			renderSeaHorseToSvg(canvas, charge.Tincture.Hexcode, lineColor)
		case "serpent-nowed":
			renderSerpentNowedToSvg(canvas, charge.Tincture.Hexcode, lineColor)
		case "squirrel":
			renderSquirrelToSvg(canvas, charge.Tincture.Hexcode, lineColor)
		case "stag-lodged":
			renderStagLodgedToSvg(canvas, charge.Tincture.Hexcode, lineColor)
		case "stag-statant":
			renderStagStatantToSvg(canvas, charge.Tincture.Hexcode, lineColor)
		case "sun-in-splendor":
			renderSunInSplendorToSvg(canvas, charge.Tincture.Hexcode, lineColor)
		case "tiger-passant":
			renderTigerPassantToSvg(canvas, charge.Tincture.Hexcode, lineColor)
		case "tiger-rampant":
			renderTigerRampantToSvg(canvas, charge.Tincture.Hexcode, lineColor)
		case "tower":
			renderTowerToSvg(canvas, charge.Tincture.Hexcode, lineColor)
		case "two-axes-in-saltire":
			renderTwoAxesInSaltireToSvg(canvas, charge.Tincture.Hexcode, lineColor)
		case "two-bones-in-saltire":
			renderTwoBonesInSaltireToSvg(canvas, charge.Tincture.Hexcode, lineColor)
		case "unicorn-rampant":
			renderUnicornRampantToSvg(canvas, charge.Tincture.Hexcode, lineColor)
		case "unicorn-statant":
			renderUnicornStatantToSvg(canvas, charge.Tincture.Hexcode, lineColor)
		case "wolf-passant":
			renderWolfPassantToSvg(canvas, charge.Tincture.Hexcode, lineColor)
		case "wolf-rampant":
			renderWolfRampantToSvg(canvas, charge.Tincture.Hexcode, lineColor)
		case "wyvern":
			renderWyvernToSvg(canvas, charge.Tincture.Hexcode, lineColor)
		}

	}
	canvas.Path("m10.273 21.598v151.22c0 96.872 89.031 194.34 146.44 240.09 57.414-45.758 146.44-143.22 146.44-240.09v-151.22h-292.89z", "stroke:#000000;stroke-width:4;fill:none")
	canvas.Gend()
	canvas.End()

	defer writer.Close()
}
