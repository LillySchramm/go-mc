package bot

import (
	"errors"
	"sync"

	"github.com/google/uuid"

	"github.com/LillySchramm/go-mc/data/packetid"
	"github.com/LillySchramm/go-mc/net"
	pk "github.com/LillySchramm/go-mc/net/packet"
	"github.com/LillySchramm/go-mc/net/queue"
	"github.com/LillySchramm/go-mc/registry"
)

// Client is used to access Minecraft server
type Client struct {
	Conn *Conn
	Auth Auth

	// These are filled when login process
	Name       string
	UUID       uuid.UUID
	Registries registry.Registries
	Cookies    map[string][]byte

	// Ingame packet handlers
	Events Events

	// Login plugins
	LoginPlugin map[string]CustomPayloadHandler

	// Configuration handler
	ConfigHandler

	CustomReportDetails map[string]string
}

// CustomPayloadHandler is a function handling custom payload
type CustomPayloadHandler func(data []byte) ([]byte, error)

func (c *Client) Close() error {
	return c.Conn.Close()
}

// NewClient init and return a new Client.
//
// A new Client has default name "Steve" and zero UUID.
// It is usable for an offline-mode game.
//
// For online-mode, you need login your Mojang account
// and load your Name, UUID and AccessToken to client.
func NewClient() *Client {
	return &Client{
		Auth:                Auth{Name: "Steve"},
		Registries:          registry.NewNetworkCodec(),
		Events:              Events{handlers: make([][]PacketHandler, packetid.ClientboundPacketIDGuard)},
		LoginPlugin:         make(map[string]CustomPayloadHandler),
		ConfigHandler:       NewDefaultConfigHandler(),
		CustomReportDetails: make(map[string]string),
	}
}

// Conn is a concurrently-safe warpper of net.Conn with packet queue.
// Note that not all methods are concurrently-safe.
type Conn struct {
	*net.Conn
	send, recv queue.Queue[pk.Packet]
	pool       sync.Pool // pool of recv packet data
	rerr       error
}

func warpConn(c *net.Conn, qr, qw queue.Queue[pk.Packet]) *Conn {
	wc := Conn{
		Conn: c,
		send: qw,
		recv: qr,
		pool: sync.Pool{New: func() any { return []byte{} }},
		rerr: nil,
	}
	go func() {
		for {
			// take a buffer from pool, after the packet is handled we put it back
			p := pk.Packet{Data: wc.pool.Get().([]byte)}
			if err := c.ReadPacket(&p); err != nil {
				wc.rerr = err
				break
			}
			if ok := wc.recv.Push(p); !ok {
				wc.rerr = errors.New("receive queue is full")
				break
			}
		}
		wc.recv.Close()
	}()
	go func() {
		for {
			p, ok := wc.send.Pull()
			if !ok {
				break
			}
			if err := c.WritePacket(p); err != nil {
				break
			}
		}
	}()

	return &wc
}

func (c *Conn) ReadPacket(p *pk.Packet) error {
	packet, ok := c.recv.Pull()
	if !ok {
		return c.rerr
	}
	*p = packet
	return nil
}

func (c *Conn) WritePacket(p pk.Packet) error {
	ok := c.send.Push(p)
	if !ok {
		return errors.New("queue is full")
	}
	return nil
}

func (c *Conn) Close() error {
	c.send.Close()
	return c.Conn.Close()
}

// Position is a 3D vector.
type Position struct {
	X, Y, Z int
}
