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

	OrdinaryPale               = Charge{"pale", "pale", "pale", "pales", "", "a", Tincture{}}
	OrdinaryFess               = Charge{"fess", "fess", "fess", "fesses", "", "a", Tincture{}}
	OrdinaryCross              = Charge{"cross", "cross", "cross", "crosses", "", "a", Tincture{}}
	OrdinaryBend               = Charge{"bend", "bend", "bend", "bends", "", "a", Tincture{}}
	OrdinarySaltire            = Charge{"saltire", "saltire", "saltire", "saltires", "", "a", Tincture{}}
	OrdinaryChevron            = Charge{"chevron", "chevron", "chevron", "chevrons", "", "a", Tincture{}}
	OrdinaryChief              = Charge{"chief", "chief", "chief", "chiefs", "", "a", Tincture{}}
	OrdinaryPile               = Charge{"pile", "pile", "pile", "piles", "", "a", Tincture{}}
	OrdinaryPall               = Charge{"pall", "pall", "pall", "palls", "", "a", Tincture{}}
	OrdinaryBordure            = Charge{"bordure", "bordure", "bordure", "bordures", "", "a", Tincture{}}
	OrdinaryLozenge            = Charge{"lozenge", "lozenge", "lozenge", "lozenges", "", "a", Tincture{}}
	OrdinaryRoundel            = Charge{"roundel", "roundel", "roundel", "roundels", "", "a", Tincture{}}
	EagleDisplayed             = Charge{"eagle-displayed", "eagle displayed", "eagle", "eagles", "displayed", "an", Tincture{}}
	DragonPassant              = Charge{"dragon-passant", "dragon passant", "dragon", "dragons", "passant", "a", Tincture{}}
	GryphonPassant             = Charge{"gryphon-passant", "gryphon passant", "gryphon", "gryphons", "passant", "a", Tincture{}}
	FoxPassant                 = Charge{"fox-passant", "fox passant", "fox", "foxes", "passant", "a", Tincture{}}
	AntelopePassant            = Charge{"antelope-passant", "antelope passant", "antelope", "antelopes", "passant", "an", Tincture{}}
	AntelopeRampant            = Charge{"antelope-rampant", "antelope rampant", "antelope", "antelopes", "rampant", "an", Tincture{}}
	BatVolant                  = Charge{"bat-volant", "bat volant", "bat", "bats", "volant", "a", Tincture{}}
	Battleaxe                  = Charge{"battleaxe", "battleaxe", "battleaxe", "battleaxes", "", "a", Tincture{}}
	BearHeadCouped             = Charge{"bear-head-couped", "bear's head couped", "bear's head", "bear's heads", "couped", "a", Tincture{}}
	BearRampant                = Charge{"bear-rampant", "bear rampant", "bear", "bears", "rampant", "a", Tincture{}}
	BearStatant                = Charge{"bear-statant", "bear statant", "bear", "bears", "statant", "a", Tincture{}}
	BeeVolant                  = Charge{"bee-volant", "bee volant", "bee", "bees", "volant", "a", Tincture{}}
	Bell                       = Charge{"bell", "bell", "bell", "bells", "", "a", Tincture{}}
	BoarHeadErased             = Charge{"boar-head-erased", "boar's head erased", "boar's head", "boar's heads", "erased", "a", Tincture{}}
	BoarPassant                = Charge{"boar-passant", "boar passant", "boar", "boars", "passant", "a", Tincture{}}
	BoarRampant                = Charge{"boar-rampant", "boar rampant", "boar", "boars", "rampant", "a", Tincture{}}
	BugleHorn                  = Charge{"bugle-horn", "bugle horn", "bugle horn", "bugle horns", "", "a", Tincture{}}
	BullPassant                = Charge{"bull-passant", "bull passant", "bull", "bulls", "passant", "a", Tincture{}}
	BullRampant                = Charge{"bull-rampant", "bull rampant", "bull", "bulls", "rampant", "a", Tincture{}}
	Castle                     = Charge{"castle", "castle", "castle", "castles", "", "a", Tincture{}}
	Cock                       = Charge{"cock", "cock", "cock", "cocks", "", "a", Tincture{}}
	Cockatrice                 = Charge{"cockatrice", "cockatrice", "cockatrice", "cockatrices", "", "a", Tincture{}}
	Crown                      = Charge{"crown", "crown", "crown", "crowns", "", "a", Tincture{}}
	DolphinHauriant            = Charge{"dolphin-hauriant", "dolphin hauriant", "dolphin", "dolphins", "hauriant", "a", Tincture{}}
	DoubleHeadedEagleDisplayed = Charge{"double-headed-eagle-displayed", "double-headed eagle displayed", "double-headed eagle", "double-headed eagles", "displayed", "an", Tincture{}}
	DragonRampant              = Charge{"dragon-rampant", "dragon rampant", "dragon", "dragons", "rampant", "a", Tincture{}}
	EaglesHeadErased           = Charge{"eagles-head-erased", "eagle's head erased", "eagle's head", "eagle's heads", "erased", "an", Tincture{}}
	FoxSejant                  = Charge{"fox-sejant", "fox sejant", "fox", "foxes", "sejant", "a", Tincture{}}
	GryphonSegreant            = Charge{"gryphon-segreant", "gryphon segreant", "gryphon", "gryphons", "segreant", "a", Tincture{}}
	HareSalient                = Charge{"hare-salient", "hare salient", "hare", "hares", "salient", "a", Tincture{}}
	Hare                       = Charge{"hare", "hare", "hare", "hares", "", "a", Tincture{}}
	Heron                      = Charge{"heron", "heron", "heron", "herons", "", "a", Tincture{}}
	HorsePassant               = Charge{"horse-passant", "horse passant", "horse", "horses", "passant", "a", Tincture{}}
	HorseRampant               = Charge{"horse-rampant", "horse rampant", "horse", "horses", "rampant", "a", Tincture{}}
	LeopardPassant             = Charge{"leopard-passant", "leopard passant", "leopard", "leopards", "passant", "a", Tincture{}}
	LionPassant                = Charge{"lion-passant", "lion passant", "lion", "lions", "passant", "a", Tincture{}}
	LionRampant                = Charge{"lion-rampant", "lion rampant", "lion", "lions", "rampant", "a", Tincture{}}
	LionsHeadErased            = Charge{"lions-head-erased", "lion's head erased", "lion's head", "lion's heads", "erased", "a", Tincture{}}
	Owl                        = Charge{"owl", "owl", "owl", "owls", "", "an", Tincture{}}
	PegasusPassant             = Charge{"pegasus-passant", "pegasus passant", "pegasus", "pegasi", "passant", "a", Tincture{}}
	PegasusRampant             = Charge{"pegasus-rampant", "pegasus rampant", "pegasus", "pegasi", "rampant", "a", Tincture{}}
	RamRampant                 = Charge{"ram-rampant", "ram rampant", "ram", "rams", "rampant", "a", Tincture{}}
	RamStatant                 = Charge{"ram-statant", "ram statant", "ram", "rams", "statant", "a", Tincture{}}
	Rose                       = Charge{"rose", "rose", "rose", "roses", "", "a", Tincture{}}
	SeaHorse                   = Charge{"sea-horse", "sea horse", "sea horse", "sea horses", "", "a", Tincture{}}
	SerpentNowed               = Charge{"serpent-nowed", "serpent nowed", "serpent", "serpents", "nowed", "a", Tincture{}}
	Squirrel                   = Charge{"squirrel", "squirrel", "squirrel", "squirrels", "", "a", Tincture{}}
	StagLodged                 = Charge{"stag-lodged", "stag lodged", "stag", "stags", "", "a", Tincture{}}
	StagStatant                = Charge{"stag-statant", "stag statant", "stag", "stags", "", "a", Tincture{}}
	SunInSplendor              = Charge{"sun-in-splendor", "sun in splendor", "sun", "suns", "in splendor", "a", Tincture{}}
	TigerPassant               = Charge{"tiger-passant", "tiger passant", "tiger", "tigers", "passant", "a", Tincture{}}
	TigerRampant               = Charge{"tiger-rampant", "tiger rampant", "tiger", "tigers", "rampant", "a", Tincture{}}
	Tower                      = Charge{"tower", "tower", "tower", "towers", "", "a", Tincture{}}
	TwoAxesInSaltire           = Charge{"two-axes-in-saltire", "two axes in saltire", "two axes", "two axes", "in saltire", "", Tincture{}}
	TwoBonesInSaltire          = Charge{"two-bones-in-saltire", "two bones in saltire", "two bones", "two bones", "in saltire", "", Tincture{}}
	UnicornRampant             = Charge{"unicorn-rampant", "unicorn rampant", "unicorn", "unicorns", "rampant", "a", Tincture{}}
	UnicornStatant             = Charge{"unicorn-statant", "unicorn statant", "unicorn", "unicorns", "statant", "a", Tincture{}}
	WolfPassant                = Charge{"wolf-passant", "wolf passant", "wolf", "wolves", "passant", "a", Tincture{}}
	WolfRampant                = Charge{"wolf-rampant", "wolf rampant", "wolf", "wolves", "rampant", "a", Tincture{}}
	Wyvern                     = Charge{"wyvern", "wyvern", "wyvern", "wyverns", "", "a", Tincture{}}
	AvailableCharges           = [68]Charge{OrdinaryPale, OrdinaryFess, OrdinaryCross, OrdinaryBend, OrdinarySaltire, OrdinaryChevron, OrdinaryChief, OrdinaryPile, OrdinaryPall, OrdinaryBordure, OrdinaryLozenge, OrdinaryRoundel, EagleDisplayed, DragonPassant, GryphonPassant, FoxPassant, AntelopeRampant, BatVolant, Battleaxe, BearHeadCouped, BearRampant, BearStatant, BeeVolant, Bell, BoarHeadErased, BoarPassant, BoarRampant, BugleHorn, BullPassant, BullRampant, Castle, Cock, Cockatrice, Crown, DolphinHauriant, DoubleHeadedEagleDisplayed, DragonRampant, EaglesHeadErased, FoxSejant, GryphonSegreant, HareSalient, Hare, Heron, HorsePassant, HorseRampant, LeopardPassant, LionPassant, LionRampant, LionsHeadErased, Owl, PegasusPassant, RamRampant, RamStatant, Rose, SeaHorse, Squirrel, StagLodged, StagStatant, SunInSplendor, TigerPassant, TigerRampant, Tower, TwoAxesInSaltire, TwoBonesInSaltire, UnicornStatant, WolfPassant, WolfRampant, Wyvern}
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

	if shallWeIncludeCharges() {
		charge := randomCharge()
		charge.Tincture = chargeTincture
		charges = append(charges, charge)
	}

	f := Field{Division: division, Tincture: fieldTincture1}
	d := Device{Field: f, Charges: charges}

	return d
}
