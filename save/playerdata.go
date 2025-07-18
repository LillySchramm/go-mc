package save

import (
	"io"

	"github.com/LillySchramm/go-mc/nbt"
)

type PlayerData struct {
	DataVersion int32

	Dimension    string
	Pos          [3]float64
	Motion       [3]float64
	Rotation     [2]float32
	FallDistance float32
	FallFlying   byte
	OnGround     byte

	UUID [4]int32

	PlayerGameType  int32 `nbt:"playerGameType"`
	Air             int16
	DeathTime       int16
	Fire            int16
	HurtTime        int16
	Health          float32
	HurtByTimestamp int32
	PortalCooldown  int32

	Invulnerable     byte
	SeenCredits      byte `nbt:"seenCredits"`
	SelectedItemSlot int32
	Score            int32
	AbsorptionAmount float32

	Inventory, EnderItems []Item

	XpLevel int32
	XpP     float32
	XpTotal int32
	XpSeed  int32

	FoodExhaustionLevel float32 `nbt:"foodExhaustionLevel"`
	FoodLevel           int32   `nbt:"foodLevel"`
	FoodSaturationLevel float32 `nbt:"foodSaturationLevel"`
	FoodTickTimer       int32   `nbt:"foodTickTimer"`

	Attributes []struct {
		Base float64
		Name string
	}

	Abilities struct {
		FlySpeed     float32 `nbt:"flySpeed"`
		WalkSpeed    float32 `nbt:"walkSpeed"`
		Flying       byte    `nbt:"flying"`
		InstantBuild byte    `nbt:"instabuild"`
		Invulnerable byte    `nbt:"invulnerable"`
		MayBuild     byte    `nbt:"mayBuild"`
		MayFly       byte    `nbt:"mayfly"`
	} `nbt:"abilities"`

	RecipeBook struct {
		IsFilteringCraftable        byte `nbt:"isFilteringCraftable"`
		IsFurnaceFilteringCraftable byte `nbt:"isFurnaceFilteringCraftable"`
		IsFurnaceGUIOpen            byte `nbt:"isFurnaceGuiOpen"`
		IsGUIOpen                   byte `nbt:"isGuiOpen"`
	} `nbt:"recipeBook"`
}

type Item struct {
	Count byte
	Slot  byte
	ID    string         `nbt:"id"`
	Tag   map[string]any `nbt:"tag"`
}

func ReadPlayerData(r io.Reader) (data PlayerData, err error) {
	_, err = nbt.NewDecoder(r).Decode(&data)
	return
}
