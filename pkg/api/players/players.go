package players

import (
	"github.com/resourcemod/api/pkg/api/entities"
	"strconv"
)

type Player struct {
	Slot      int
	Name      string
	SteamID   string
	SteamID32 int32
	Entity    entities.Entity
}

func NewPlayer(slot int, name string, steamID32 int32, entity entities.Entity) Player {
	return Player{
		slot,
		name,
		steamInt32ToString(steamID32),
		steamID32,
		entity,
	}
}

// https://github.com/MrWaggel/gosteamconv/blob/master/gosteamconv.go thanks!
func steamInt32ToString(steamID32 int32) string {
	if steamID32 == 0 {
		return "BOT"
	}
	remainder := steamID32 % 2
	steamID32 = steamID32 / 2
	return "STEAM_0:" + strconv.FormatInt(int64(remainder), 10) + ":" + strconv.FormatInt(int64(steamID32), 10)
}
