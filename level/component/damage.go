package component

import pk "github.com/LillySchramm/go-mc/net/packet"

var _ DataComponent = (*Damage)(nil)

type Damage struct {
	pk.VarInt
}

// ID implements DataComponent.
func (Damage) ID() string {
	return "minecraft:damage"
}
