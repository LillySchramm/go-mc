package component

import (
	"io"

	"github.com/LillySchramm/go-mc/nbt/dynbt"
	pk "github.com/LillySchramm/go-mc/net/packet"
)

var _ DataComponent = (*BucketEntityData)(nil)

type BucketEntityData struct {
	dynbt.Value
}

// ID implements DataComponent.
func (BucketEntityData) ID() string {
	return "minecraft:bucket_entity_data"
}

// ReadFrom implements DataComponent.
func (b *BucketEntityData) ReadFrom(r io.Reader) (n int64, err error) {
	return pk.NBT(&b.Value).ReadFrom(r)
}

// WriteTo implements DataComponent.
func (b *BucketEntityData) WriteTo(w io.Writer) (n int64, err error) {
	return pk.NBT(&b.Value).WriteTo(w)
}
