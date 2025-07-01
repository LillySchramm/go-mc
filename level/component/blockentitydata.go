package component

import (
	"io"

	"github.com/LillySchramm/go-mc/nbt/dynbt"
	pk "github.com/LillySchramm/go-mc/net/packet"
)

var _ DataComponent = (*BlockEntityData)(nil)

type BlockEntityData struct {
	dynbt.Value
}

// ID implements DataComponent.
func (BlockEntityData) ID() string {
	return "minecraft:block_entity_data"
}

// ReadFrom implements DataComponent.
func (b *BlockEntityData) ReadFrom(r io.Reader) (n int64, err error) {
	return pk.NBT(&b.Value).ReadFrom(r)
}

// WriteTo implements DataComponent.
func (b *BlockEntityData) WriteTo(w io.Writer) (n int64, err error) {
	return pk.NBT(&b.Value).WriteTo(w)
}
