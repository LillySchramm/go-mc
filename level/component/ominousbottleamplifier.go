package component

import (
	pk "github.com/LillySchramm/go-mc/net/packet"
)

var _ DataComponent = (*OminousBottleAmplifier)(nil)

type OminousBottleAmplifier struct {
	pk.VarInt
}

// ID implements DataComponent.
func (OminousBottleAmplifier) ID() string {
	return "minecraft:ominous_bottle_amplifier"
}
