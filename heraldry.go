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

	OrdinaryPale               = Charge{"pale", "pale", Tincture{}}
	OrdinaryFess               = Charge{"fess", "fess", Tincture{}}
	OrdinaryCross              = Charge{"cross", "cross", Tincture{}}
	OrdinaryBend               = Charge{"bend", "bend", Tincture{}}
	OrdinarySaltire            = Charge{"saltire", "saltire", Tincture{}}
	OrdinaryChevron            = Charge{"chevron", "chevron", Tincture{}}
	OrdinaryChief              = Charge{"chief", "chief", Tincture{}}
	OrdinaryPile               = Charge{"pile", "pile", Tincture{}}
	OrdinaryPall               = Charge{"pall", "pall", Tincture{}}
	OrdinaryBordure            = Charge{"bordure", "bordure", Tincture{}}
	OrdinaryLozenge            = Charge{"lozenge", "lozenge", Tincture{}}
	OrdinaryRoundel            = Charge{"roundel", "roundel", Tincture{}}
	EagleDisplayed             = Charge{"eagle-displayed", "eagle displayed", Tincture{}}
	DragonPassant              = Charge{"dragon-passant", "dragon passant", Tincture{}}
	GryphonPassant             = Charge{"gryphon-passant", "gryphon passant", Tincture{}}
	FoxPassant                 = Charge{"fox-passant", "fox passant", Tincture{}}
	AntelopePassant            = Charge{"antelope-passant", "antelope passant", Tincture{}}
	AntelopeRampant            = Charge{"antelope-rampant", "antelope rampant", Tincture{}}
	BatVolant                  = Charge{"bat-volant", "bat volant", Tincture{}}
	Battleaxe                  = Charge{"battleaxe", "battleaxe", Tincture{}}
	BearHeadCouped             = Charge{"bear-head-couped", "bear head couped", Tincture{}}
	BearRampant                = Charge{"bear-rampant", "bear rampant", Tincture{}}
	BearStatant                = Charge{"bear-statant", "bear statant", Tincture{}}
	BeeVolant                  = Charge{"bee-volant", "bee volant", Tincture{}}
	Bell                       = Charge{"bell", "bell", Tincture{}}
	BoarHeadErased             = Charge{"boar-head-erased", "boar head erased", Tincture{}}
	BoarPassant                = Charge{"boar-passant", "boar passant", Tincture{}}
	BoarRampant                = Charge{"boar-rampant", "boar rampant", Tincture{}}
	BugleHorn                  = Charge{"bugle-horn", "bugle horn", Tincture{}}
	BullPassant                = Charge{"bull-passant", "bull passant", Tincture{}}
	BullRampant                = Charge{"bull-rampant", "bull rampant", Tincture{}}
	Castle                     = Charge{"castle", "castle", Tincture{}}
	Cock                       = Charge{"cock", "cock", Tincture{}}
	Cockatrice                 = Charge{"cockatrice", "cockatrice", Tincture{}}
	Crown                      = Charge{"crown", "crown", Tincture{}}
	DolphinHauriant            = Charge{"dolphin-hauriant", "dolphin hauriant", Tincture{}}
	DoubleHeadedEagleDisplayed = Charge{"double-headed-eagle-displayed", "double-headed eagle displayed", Tincture{}}
	DragonRampant              = Charge{"dragon-rampant", "dragon rampant", Tincture{}}
	EaglesHeadErased           = Charge{"eagles-head-erased", "eagle's head erased", Tincture{}}
	FoxSejant                  = Charge{"fox-sejant", "fox sejant", Tincture{}}
	GryphonSegreant            = Charge{"gryphon-segreant", "gryphon segreant", Tincture{}}
	HareSalient                = Charge{"hare-salient", "hare salient", Tincture{}}
	Hare                       = Charge{"hare", "hare", Tincture{}}
	Heron                      = Charge{"heron", "heron", Tincture{}}
	HorsePassant               = Charge{"horse-passant", "horse passant", Tincture{}}
	HorseRampant               = Charge{"horse-rampant", "horse rampant", Tincture{}}
	LeopardPassant             = Charge{"leopard-passant", "leopard passant", Tincture{}}
	LionPassant                = Charge{"lion-passant", "lion passant", Tincture{}}
	LionRampant                = Charge{"lion-rampant", "lion rampant", Tincture{}}
	LionsHeadErased            = Charge{"lions-head-erased", "lion's head erased", Tincture{}}
	Owl                        = Charge{"owl", "owl", Tincture{}}
	PegasusPassant             = Charge{"pegasus-passant", "pegasus passant", Tincture{}}
	PegasusRampant             = Charge{"pegasus-rampant", "pegasus rampant", Tincture{}}
	RamRampant                 = Charge{"ram-rampant", "ram rampant", Tincture{}}
	RamStatant                 = Charge{"ram-statant", "ram statant", Tincture{}}
	Rose                       = Charge{"rose", "rose", Tincture{}}
	SeaHorse                   = Charge{"sea-horse", "sea horse", Tincture{}}
	SerpentNowed               = Charge{"serpent-nowed", "serpent nowed", Tincture{}}
	Squirrel                   = Charge{"squirrel", "squirrel", Tincture{}}
	StagLodged                 = Charge{"stag-lodged", "stag lodged", Tincture{}}
	StagStatant                = Charge{"stag-statant", "stag statant", Tincture{}}
	SunInSplendor              = Charge{"sun-in-splendor", "sun in splendor", Tincture{}}
	TigerPassant               = Charge{"tiger-passant", "tiger passant", Tincture{}}
	TigerRampant               = Charge{"tiger-rampant", "tiger rampant", Tincture{}}
	Tower                      = Charge{"tower", "tower", Tincture{}}
	TwoAxesInSaltire           = Charge{"two-axes-in-saltire", "two axes in saltire", Tincture{}}
	TwoBonesInSaltire          = Charge{"two-bones-in-saltire", "two bones in saltire", Tincture{}}
	UnicornRampant             = Charge{"unicorn-rampant", "unicorn rampant", Tincture{}}
	UnicornStatant             = Charge{"unicorn-statant", "unicorn statant", Tincture{}}
	WolfPassant                = Charge{"wolf-passant", "wolf passant", Tincture{}}
	WolfRampant                = Charge{"wolf-rampant", "wolf rampant", Tincture{}}
	Wyvern                     = Charge{"wyvern", "wyvern", Tincture{}}
	AvailableCharges           = [72]Charge{OrdinaryPale, OrdinaryFess, OrdinaryCross, OrdinaryBend, OrdinarySaltire, OrdinaryChevron, OrdinaryChief, OrdinaryPile, OrdinaryPall, OrdinaryBordure, OrdinaryLozenge, OrdinaryRoundel, EagleDisplayed, DragonPassant, GryphonPassant, FoxPassant, AntelopePassant, AntelopeRampant, BatVolant, Battleaxe, BearHeadCouped, BearRampant, BearStatant, BeeVolant, Bell, BoarHeadErased, BoarPassant, BoarRampant, BugleHorn, BullPassant, BullRampant, Castle, Cock, Cockatrice, Crown, DolphinHauriant, DoubleHeadedEagleDisplayed, DragonRampant, EaglesHeadErased, FoxSejant, GryphonSegreant, HareSalient, Hare, Heron, HorsePassant, HorseRampant, LeopardPassant, LionPassant, LionRampant, LionsHeadErased, Owl, PegasusPassant, PegasusRampant, RamRampant, RamStatant, Rose, SeaHorse, SerpentNowed, Squirrel, StagLodged, StagStatant, SunInSplendor, TigerPassant, TigerRampant, Tower, TwoAxesInSaltire, TwoBonesInSaltire, UnicornRampant, UnicornStatant, WolfPassant, WolfRampant, Wyvern}
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
