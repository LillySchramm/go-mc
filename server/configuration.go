package server

import (
	"github.com/LillySchramm/go-mc/chat"
	"github.com/LillySchramm/go-mc/data/packetid"
	"github.com/LillySchramm/go-mc/net"
	pk "github.com/LillySchramm/go-mc/net/packet"
	"github.com/LillySchramm/go-mc/registry"
)

type ConfigHandler interface {
	AcceptConfig(conn *net.Conn) error
}

type Configurations struct {
	Registries registry.Registries
}

func (c *Configurations) AcceptConfig(conn *net.Conn) error {
	err := conn.WritePacket(pk.Marshal(
		packetid.ClientboundConfigRegistryData,
		pk.NBT(c.Registries),
	))
	if err != nil {
		return err
	}
	err = conn.WritePacket(pk.Marshal(
		packetid.ClientboundConfigFinishConfiguration,
	))
	return err
}

type ConfigFailErr struct {
	reason chat.Message
}

func (c ConfigFailErr) Error() string {
	return "config error: " + c.reason.ClearString()
}
