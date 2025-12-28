package save

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"errors"
	"io"

	"github.com/LillySchramm/go-mc/nbt"
)

// Chunk is 16* chunk
type Chunk struct {
	BlockEntities  []nbt.RawMessage `nbt:"block_entities"`
	BlockTicks     nbt.RawMessage   `nbt:"block_ticks"`
	CarvingMasks   map[string][]uint64
	DataVersion    int32
	Entities       []nbt.RawMessage    `nbt:"entities"`
	FluidTicks     nbt.RawMessage      `nbt:"fluid_ticks"`
	Heightmaps     map[string][]uint64 // keys: "WORLD_SURFACE_WG", "WORLD_SURFACE", "WORLD_SURFACE_IGNORE_SNOW", "OCEAN_FLOOR_WG", "OCEAN_FLOOR", "MOTION_BLOCKING", "MOTION_BLOCKING_NO_LEAVES"
	InhabitedTime  int64
	IsLightOn      byte `nbt:"isLightOn"`
	LastUpdate     int64
	Lights         []nbt.RawMessage
	PostProcessing nbt.RawMessage
	Sections       []Section `nbt:"sections"`
	Status         string
	Structures     nbt.RawMessage `nbt:"structures"`
	XPos           int32          `nbt:"xPos"`
	YPos           int32          `nbt:"yPos"`
	ZPos           int32          `nbt:"zPos"`
}

type Section struct {
	Y           int8
	BlockStates PaletteContainer[BlockState] `nbt:"block_states"`
	Biomes      PaletteContainer[BiomeState] `nbt:"biomes"`
	SkyLight    []byte
	BlockLight  []byte
}

type PaletteContainer[T any] struct {
	Palette []T      `nbt:"palette"`
	Data    []uint64 `nbt:"data"`
}

type Properties struct {
	Age                  string `nbt:"age,omitempty"`
	Attached             string `nbt:"attached,omitempty"`
	Attachment           string `nbt:"attachment,omitempty"`
	Axis                 string `nbt:"axis,omitempty"`
	Berries              string `nbt:"berries,omitempty"`
	Bites                string `nbt:"bites,omitempty"`
	Bloom                string `nbt:"bloom,omitempty"`
	Bottom               string `nbt:"bottom,omitempty"`
	Can_Summon           string `nbt:"can_summon,omitempty"`
	Candles              string `nbt:"candles,omitempty"`
	Charges              string `nbt:"charges,omitempty"`
	Conditional          string `nbt:"conditional,omitempty"`
	Copper_Golem_Pose    string `nbt:"copper_golem_pose,omitempty"`
	Cracked              string `nbt:"cracked,omitempty"`
	Crafting             string `nbt:"crafting,omitempty"`
	Creaking_Heart_State string `nbt:"creaking_heart_state,omitempty"`
	Delay                string `nbt:"delay,omitempty"`
	Disarmed             string `nbt:"disarmed,omitempty"`
	Distance             string `nbt:"distance,omitempty"`
	Down                 string `nbt:"down,omitempty"`
	Drag                 string `nbt:"drag,omitempty"`
	Dusted               string `nbt:"dusted,omitempty"`
	East                 string `nbt:"east,omitempty"`
	Eggs                 string `nbt:"eggs,omitempty"`
	Enabled              string `nbt:"enabled,omitempty"`
	Extended             string `nbt:"extended,omitempty"`
	Eye                  string `nbt:"eye,omitempty"`
	Face                 string `nbt:"face,omitempty"`
	Facing               string `nbt:"facing,omitempty"`
	Flower_Amount        string `nbt:"flower_amount,omitempty"`
	Half                 string `nbt:"half,omitempty"`
	Hanging              string `nbt:"hanging,omitempty"`
	Has_Book             string `nbt:"has_book,omitempty"`
	Has_Bottle_0         string `nbt:"has_bottle_0,omitempty"`
	Has_Bottle_1         string `nbt:"has_bottle_1,omitempty"`
	Has_Bottle_2         string `nbt:"has_bottle_2,omitempty"`
	Has_Record           string `nbt:"has_record,omitempty"`
	Hatch                string `nbt:"hatch,omitempty"`
	Hinge                string `nbt:"hinge,omitempty"`
	Honey_Level          string `nbt:"honey_level,omitempty"`
	Hydration            string `nbt:"hydration,omitempty"`
	In_Wall              string `nbt:"in_wall,omitempty"`
	Instrument           string `nbt:"instrument,omitempty"`
	Inverted             string `nbt:"inverted,omitempty"`
	Layers               string `nbt:"layers,omitempty"`
	Leaves               string `nbt:"leaves,omitempty"`
	Level                string `nbt:"level,omitempty"`
	Lit                  string `nbt:"lit,omitempty"`
	Locked               string `nbt:"locked,omitempty"`
	Mode                 string `nbt:"mode,omitempty"`
	Moisture             string `nbt:"moisture,omitempty"`
	Natural              string `nbt:"natural,omitempty"`
	North                string `nbt:"north,omitempty"`
	Note                 string `nbt:"note,omitempty"`
	Occupied             string `nbt:"occupied,omitempty"`
	Ominous              string `nbt:"ominous,omitempty"`
	Open                 string `nbt:"open,omitempty"`
	Orientation          string `nbt:"orientation,omitempty"`
	Part                 string `nbt:"part,omitempty"`
	Persistent           string `nbt:"persistent,omitempty"`
	Pickles              string `nbt:"pickles,omitempty"`
	Power                string `nbt:"power,omitempty"`
	Powered              string `nbt:"powered,omitempty"`
	Rotation             string `nbt:"rotation,omitempty"`
	Sculk_Sensor_Phase   string `nbt:"sculk_sensor_phase,omitempty"`
	Segment_Amount       string `nbt:"segment_amount,omitempty"`
	Shape                string `nbt:"shape,omitempty"`
	Short                string `nbt:"short,omitempty"`
	Shrieking            string `nbt:"shrieking,omitempty"`
	Side_Chain           string `nbt:"side_chain,omitempty"`
	Signal_Fire          string `nbt:"signal_fire,omitempty"`
	Slot_0_Occupied      string `nbt:"slot_0_occupied,omitempty"`
	Slot_1_Occupied      string `nbt:"slot_1_occupied,omitempty"`
	Slot_2_Occupied      string `nbt:"slot_2_occupied,omitempty"`
	Slot_3_Occupied      string `nbt:"slot_3_occupied,omitempty"`
	Slot_4_Occupied      string `nbt:"slot_4_occupied,omitempty"`
	Slot_5_Occupied      string `nbt:"slot_5_occupied,omitempty"`
	Snowy                string `nbt:"snowy,omitempty"`
	South                string `nbt:"south,omitempty"`
	Stage                string `nbt:"stage,omitempty"`
	Thickness            string `nbt:"thickness,omitempty"`
	Tilt                 string `nbt:"tilt,omitempty"`
	Tip                  string `nbt:"tip,omitempty"`
	Trial_Spawner_State  string `nbt:"trial_spawner_state,omitempty"`
	Triggered            string `nbt:"triggered,omitempty"`
	Type                 string `nbt:"type,omitempty"`
	Unstable             string `nbt:"unstable,omitempty"`
	Up                   string `nbt:"up,omitempty"`
	Vault_State          string `nbt:"vault_state,omitempty"`
	Vertical_Direction   string `nbt:"vertical_direction,omitempty"`
	Waterlogged          string `nbt:"waterlogged,omitempty"`
	West                 string `nbt:"west,omitempty"`
}

type BlockState struct {
	Name       string
	Properties Properties `nbt:"Properties"`
}

type BiomeState string

// Load read column data from []byte
func (c *Chunk) Load(data []byte) (err error) {
	var r io.Reader = bytes.NewReader(data[1:])

	switch data[0] {
	default:
		err = errors.New("unknown compression")
	case 1:
		r, err = gzip.NewReader(r)
	case 2:
		r, err = zlib.NewReader(r)
	case 3:
		// none compression
	}
	if err != nil {
		return err
	}

	d := nbt.NewDecoder(r)
	d.DisallowUnknownFields()
	_, err = d.Decode(c)
	return err
}

func (c *Chunk) Data(compressingType byte) ([]byte, error) {
	var buff bytes.Buffer

	buff.WriteByte(compressingType)
	var w io.Writer
	switch compressingType {
	default:
		return nil, errors.New("unknown compression")
	case 1:
		w = gzip.NewWriter(&buff)
	case 2:
		w = zlib.NewWriter(&buff)
	case 3:
		w = &buff
	}
	err := nbt.NewEncoder(w).Encode(c, "")
	return buff.Bytes(), err
}

type Entities struct {
	Pos, Motion  [3]float64
	Rotation     [3]float32
	FallDistance float32
	Fire, Air    int16

	OnGround       bool
	Invulnerable   bool
	PortalCooldown int32
	UUID           [4]int32

	CustomName        string
	CustomNameVisible bool
	Silent            bool
	NoGravity         bool
	Glowing           bool
	TicksFrozen       int32
	HasVisualFire     bool
	Tags              []string
}
