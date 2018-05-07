package heraldry

import (
	"math/rand"
	"time"
)

type Tincture struct {
	Type    string
	Name    string
	Hexcode string
}

type Charge struct {
	Identifier string
	Name       string
	Noun       string
	NounPlural string
	Descriptor string
	Article    string
	SingleOnly bool
}

type ChargeGroup struct {
	Charges []Charge
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
	ChargeGroups []ChargeGroup
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

	OrdinaryPale               = Charge{"pale", "pale", "pale", "pales", "", "a", false}
	OrdinaryFess               = Charge{"fess", "fess", "fess", "fesses", "", "a", true}
	OrdinaryCross              = Charge{"cross", "cross", "cross", "crosses", "", "a", true}
	OrdinaryBend               = Charge{"bend", "bend", "bend", "bends", "", "a", true}
	OrdinarySaltire            = Charge{"saltire", "saltire", "saltire", "saltires", "", "a", true}
	OrdinaryChevron            = Charge{"chevron", "chevron", "chevron", "chevrons", "", "a", true}
	OrdinaryChief              = Charge{"chief", "chief", "chief", "chiefs", "", "a", true}
	OrdinaryPile               = Charge{"pile", "pile", "pile", "piles", "", "a", true}
	OrdinaryPall               = Charge{"pall", "pall", "pall", "palls", "", "a", true}
	OrdinaryBordure            = Charge{"bordure", "bordure", "bordure", "bordures", "", "a", true}
	OrdinaryLozenge            = Charge{"lozenge", "lozenge", "lozenge", "lozenges", "", "a", false}
	OrdinaryRoundel            = Charge{"roundel", "roundel", "roundel", "roundels", "", "a", false}
	EagleDisplayed             = Charge{"eagle-displayed", "eagle displayed", "eagle", "eagles", "displayed", "an", false}
	DragonPassant              = Charge{"dragon-passant", "dragon passant", "dragon", "dragons", "passant", "a", false}
	GryphonPassant             = Charge{"gryphon-passant", "gryphon passant", "gryphon", "gryphons", "passant", "a", false}
	FoxPassant                 = Charge{"fox-passant", "fox passant", "fox", "foxes", "passant", "a", false}
	AntelopePassant            = Charge{"antelope-passant", "antelope passant", "antelope", "antelopes", "passant", "an", false}
	AntelopeRampant            = Charge{"antelope-rampant", "antelope rampant", "antelope", "antelopes", "rampant", "an", false}
	BatVolant                  = Charge{"bat-volant", "bat volant", "bat", "bats", "volant", "a", false}
	Battleaxe                  = Charge{"battleaxe", "battleaxe", "battleaxe", "battleaxes", "", "a", false}
	BearHeadCouped             = Charge{"bear-head-couped", "bear's head couped", "bear's head", "bear's heads", "couped", "a", false}
	BearRampant                = Charge{"bear-rampant", "bear rampant", "bear", "bears", "rampant", "a", false}
	BearStatant                = Charge{"bear-statant", "bear statant", "bear", "bears", "statant", "a", false}
	BeeVolant                  = Charge{"bee-volant", "bee volant", "bee", "bees", "volant", "a", false}
	Bell                       = Charge{"bell", "bell", "bell", "bells", "", "a", false}
	BoarHeadErased             = Charge{"boar-head-erased", "boar's head erased", "boar's head", "boar's heads", "erased", "a", false}
	BoarPassant                = Charge{"boar-passant", "boar passant", "boar", "boars", "passant", "a", false}
	BoarRampant                = Charge{"boar-rampant", "boar rampant", "boar", "boars", "rampant", "a", false}
	BugleHorn                  = Charge{"bugle-horn", "bugle horn", "bugle horn", "bugle horns", "", "a", false}
	BullPassant                = Charge{"bull-passant", "bull passant", "bull", "bulls", "passant", "a", false}
	BullRampant                = Charge{"bull-rampant", "bull rampant", "bull", "bulls", "rampant", "a", false}
	Castle                     = Charge{"castle", "castle", "castle", "castles", "", "a", false}
	Cock                       = Charge{"cock", "cock", "cock", "cocks", "", "a", false}
	Cockatrice                 = Charge{"cockatrice", "cockatrice", "cockatrice", "cockatrices", "", "a", false}
	Crown                      = Charge{"crown", "crown", "crown", "crowns", "", "a", false}
	DolphinHauriant            = Charge{"dolphin-hauriant", "dolphin hauriant", "dolphin", "dolphins", "hauriant", "a", false}
	DoubleHeadedEagleDisplayed = Charge{"double-headed-eagle-displayed", "double-headed eagle displayed", "double-headed eagle", "double-headed eagles", "displayed", "a", false}
	DragonRampant              = Charge{"dragon-rampant", "dragon rampant", "dragon", "dragons", "rampant", "a", false}
	EaglesHeadErased           = Charge{"eagles-head-erased", "eagle's head erased", "eagle's head", "eagle's heads", "erased", "an", false}
	FoxSejant                  = Charge{"fox-sejant", "fox sejant", "fox", "foxes", "sejant", "a", false}
	GryphonSegreant            = Charge{"gryphon-segreant", "gryphon segreant", "gryphon", "gryphons", "segreant", "a", false}
	HareSalient                = Charge{"hare-salient", "hare salient", "hare", "hares", "salient", "a", false}
	Hare                       = Charge{"hare", "hare", "hare", "hares", "", "a", false}
	Heron                      = Charge{"heron", "heron", "heron", "herons", "", "a", false}
	HorsePassant               = Charge{"horse-passant", "horse passant", "horse", "horses", "passant", "a", false}
	HorseRampant               = Charge{"horse-rampant", "horse rampant", "horse", "horses", "rampant", "a", false}
	LeopardPassant             = Charge{"leopard-passant", "leopard passant", "leopard", "leopards", "passant", "a", false}
	LionPassant                = Charge{"lion-passant", "lion passant", "lion", "lions", "passant", "a", false}
	LionRampant                = Charge{"lion-rampant", "lion rampant", "lion", "lions", "rampant", "a", false}
	LionsHeadErased            = Charge{"lions-head-erased", "lion's head erased", "lion's head", "lion's heads", "erased", "a", false}
	Owl                        = Charge{"owl", "owl", "owl", "owls", "", "an", false}
	PegasusPassant             = Charge{"pegasus-passant", "pegasus passant", "pegasus", "pegasi", "passant", "a", false}
	PegasusRampant             = Charge{"pegasus-rampant", "pegasus rampant", "pegasus", "pegasi", "rampant", "a", false}
	RamRampant                 = Charge{"ram-rampant", "ram rampant", "ram", "rams", "rampant", "a", false}
	RamStatant                 = Charge{"ram-statant", "ram statant", "ram", "rams", "statant", "a", false}
	Rose                       = Charge{"rose", "rose", "rose", "roses", "", "a", false}
	SeaHorse                   = Charge{"sea-horse", "sea horse", "sea horse", "sea horses", "", "a", false}
	SerpentNowed               = Charge{"serpent-nowed", "serpent nowed", "serpent", "serpents", "nowed", "a", false}
	Squirrel                   = Charge{"squirrel", "squirrel", "squirrel", "squirrels", "", "a", false}
	StagLodged                 = Charge{"stag-lodged", "stag lodged", "stag", "stags", "", "a", false}
	StagStatant                = Charge{"stag-statant", "stag statant", "stag", "stags", "", "a", false}
	SunInSplendor              = Charge{"sun-in-splendor", "sun in splendor", "sun", "suns", "in splendor", "a", false}
	TigerPassant               = Charge{"tiger-passant", "tiger passant", "tiger", "tigers", "passant", "a", false}
	TigerRampant               = Charge{"tiger-rampant", "tiger rampant", "tiger", "tigers", "rampant", "a", false}
	Tower                      = Charge{"tower", "tower", "tower", "towers", "", "a", false}
	TwoAxesInSaltire           = Charge{"two-axes-in-saltire", "two axes in saltire", "two axes", "two axes", "in saltire", "", false}
	TwoBonesInSaltire          = Charge{"two-bones-in-saltire", "two bones in saltire", "two bones", "two bones", "in saltire", "", false}
	UnicornRampant             = Charge{"unicorn-rampant", "unicorn rampant", "unicorn", "unicorns", "rampant", "a", false}
	UnicornStatant             = Charge{"unicorn-statant", "unicorn statant", "unicorn", "unicorns", "statant", "a", false}
	WolfPassant                = Charge{"wolf-passant", "wolf passant", "wolf", "wolves", "passant", "a", false}
	WolfRampant                = Charge{"wolf-rampant", "wolf rampant", "wolf", "wolves", "rampant", "a", false}
	Wyvern                     = Charge{"wyvern", "wyvern", "wyvern", "wyverns", "", "a", false}
	AvailableCharges           = [68]Charge{OrdinaryPale, OrdinaryFess, OrdinaryCross, OrdinaryBend, OrdinarySaltire, OrdinaryChevron, OrdinaryChief, OrdinaryPile, OrdinaryPall, OrdinaryBordure, OrdinaryLozenge, OrdinaryRoundel, EagleDisplayed, DragonPassant, GryphonPassant, FoxPassant, AntelopeRampant, BatVolant, Battleaxe, BearHeadCouped, BearRampant, BearStatant, BeeVolant, Bell, BoarHeadErased, BoarPassant, BoarRampant, BugleHorn, BullPassant, BullRampant, Castle, Cock, Cockatrice, Crown, DolphinHauriant, DoubleHeadedEagleDisplayed, DragonRampant, EaglesHeadErased, FoxSejant, HareSalient, Hare, Heron, HorsePassant, HorseRampant, LeopardPassant, LionPassant, LionRampant, LionsHeadErased, Owl, PegasusPassant, RamRampant, RamStatant, Rose, SeaHorse, Squirrel, StagLodged, StagStatant, SunInSplendor, TigerPassant, TigerRampant, Tower, TwoAxesInSaltire, TwoBonesInSaltire, UnicornStatant, WolfPassant, WolfRampant, Wyvern}
	//BrokenCharges = [4]Charge{AntelopePassant, GryphonSegreant, UnicornRampant, SerpentNowed}
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

// Generate procedurally generates a random heraldic device and returns it.
func Generate() Device {
	fieldTincture1 := randomTincture()
	fieldTincture2 := randomComplementaryTincture(fieldTincture1)
	chargeTincture := randomContrastingTincture(fieldTincture1)

	division := randomDivision()
	division.Tincture = fieldTincture2

	var charges []Charge
	var chargeGroups []ChargeGroup

	if shallWeIncludeCharges() {
		charge := randomCharge()
		rand.Seed(time.Now().UnixNano())
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
