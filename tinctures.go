package heraldry

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
)
