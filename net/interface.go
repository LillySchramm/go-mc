package net

import pk "github.com/LillySchramm/go-mc/net/packet"

type Writer interface {
	WritePacket(p pk.Packet) error
}

type Reader interface {
	ReadPacket() (pk.Packet, error)
}

type ReadWriter interface {
	Reader
	Writer
}
