package heraldry

var (
	Argent          = Tincture{"metal", "argent", "#FFFFFF"}
	Azure           = Tincture{"color", "azure", "#0000FF"}
	Gules           = Tincture{"color", "gules", "#EE0000"}
	Murrey          = Tincture{"stain", "murrey", "#591334"}
	Or              = Tincture{"metal", "Or", "#FFEC00"}
	Purpure         = Tincture{"color", "purpure", "#831CA6"}
	Sable           = Tincture{"color", "sable", "#000000"}
	Sanguine        = Tincture{"stain", "sanguine", "#720101"}
	Tenne           = Tincture{"stain", "tenné", "#7C4215"}
	Vert            = Tincture{"color", "vert", "#009900"}
	Metals          = [2]Tincture{Or, Argent}
	Colors          = [5]Tincture{Sable, Gules, Vert, Azure, Purpure}
	Stains          = [3]Tincture{Murrey, Sanguine, Tenne}
	tinctureChances = map[string]int{
		"color": 8,
		"metal": 8,
		"stain": 1,
	}
	colorOrStainChances = map[string]int{
		"color": 9,
		"stain": 1,
	}
)
