package world

import "github.com/LillySchramm/go-mc/level"

type EventsListener struct {
	LoadChunk   func(pos level.ChunkPos) error
	UnloadChunk func(pos level.ChunkPos) error
}
