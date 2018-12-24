// This file contains general enum types.

package repcore

import "fmt"

// Enum is the base / common part of enum types.
type Enum struct {
	// Name of the entity
	Name string
}

// String returns the string representation of the enum (the name).
// Defined with value receiver so this gets called even if a non-pointer is used.
func (e Enum) String() string {
	return e.Name
}

// UnknownEnum constructs a new Enum for an unknown entity with a name:
//     "Unknown 0xID"
// ID must be an integer number.
func UnknownEnum(ID interface{}) Enum {
	return Enum{fmt.Sprintf("Unknown 0x%x", ID)}
}

// Engine is the StarCraft engine / extension.
type Engine struct {
	Enum

	// ID as it appears in replays
	ID byte

	// ShortName is a shorter name
	ShortName string
}

// Engines is an enumeration of the possible engines
var Engines = []*Engine{
	{Enum{"StarCraft"}, 0x00, "SC"},
	{Enum{"Brood War"}, 0x01, "BW"},
}

// Named engines
var (
	EngineStarCraft = Engines[0]
	EngineBroodWar  = Engines[1]
)

// EngineByID returns the Engine for a given ID.
// A new Engine with Unknown name is returned if one is not found
// for the given ID (preserving the unknown ID).
func EngineByID(ID byte) *Engine {
	if int(ID) < len(Engines) {
		return Engines[ID]
	}
	return &Engine{UnknownEnum(ID), ID, "Unk"}
}

// Speed is the game speed.
type Speed struct {
	Enum

	// ID as it appears in replays
	ID byte
}

// Speeds is an enumeration of the possible speeds
var Speeds = []*Speed{
	{Enum{"Slowest"}, 0x00},
	{Enum{"Slower"}, 0x01},
	{Enum{"Slow"}, 0x02},
	{Enum{"Normal"}, 0x03},
	{Enum{"Fast"}, 0x04},
	{Enum{"Faster"}, 0x05},
	{Enum{"Fastest"}, 0x06},
}

// Named speeds
var (
	SpeedSlowest = Speeds[0]
	SpeedSlower  = Speeds[1]
	SpeedSlow    = Speeds[2]
	SpeedNormal  = Speeds[3]
	SpeedFast    = Speeds[4]
	SpeedFaster  = Speeds[5]
	SpeedFastest = Speeds[6]
)

// SpeedByID returns the Speed for a given ID.
// A new Speed with Unknown name is returned if one is not found
// for the given ID (preserving the unknown ID).
func SpeedByID(ID byte) *Speed {
	if int(ID) < len(Speeds) {
		return Speeds[ID]
	}
	return &Speed{UnknownEnum(ID), ID}
}

// GameType is the game type.
type GameType struct {
	Enum

	// ID as it appears in replays
	ID uint16

	// ShortName is a shorter name
	ShortName string
}

// GameTypes is an enumeration of the possible game types
var GameTypes = []*GameType{
	{Enum{"None"}, 0x00, "None"},
	{Enum{"Custom"}, 0x01, "Custom"}, // Warcraft III
	{Enum{"Melee"}, 0x02, "Melee"},
	{Enum{"Free For All"}, 0x03, "FFA"},
	{Enum{"One on One"}, 0x04, "1v1"},
	{Enum{"Capture The Flag"}, 0x05, "CTF"},
	{Enum{"Greed"}, 0x06, "Greed"},
	{Enum{"Slaughter"}, 0x07, "Slaughter"},
	{Enum{"Sudden Death"}, 0x08, "Sudden Death"},
	{Enum{"Ladder"}, 0x09, "Ladder"},
	{Enum{"Use map settings"}, 0x0a, "UMS"},
	{Enum{"Team Melee"}, 0x0b, "Team Melee"},
	{Enum{"Team Free For All"}, 0x0c, "Team FFA"},
	{Enum{"Team Capture The Flag"}, 0x0d, "Team CTF"},
	{UnknownEnum(0x0e), 0x0e, "Unk"},
	{Enum{"Top vs Bottom"}, 0x0f, "TvB"},
	{Enum{"Iron Man Ladder"}, 0x10, "Iron Man Ladder"}, // Warcraft II
}

// Named valid game types
var (
	GameTypeNone          = GameTypes[0]
	GameTypeMelee         = GameTypes[2]
	GameTypeFFA           = GameTypes[3]
	GameType1v1           = GameTypes[4]
	GameTypeCTF           = GameTypes[5]
	GameTypeGreed         = GameTypes[6]
	GameTypeSlaughter     = GameTypes[7]
	GameTypeSuddenDeath   = GameTypes[8]
	GameTypeLadder        = GameTypes[9]
	GameTypeUMS           = GameTypes[10]
	GameTypeTeamMelee     = GameTypes[11]
	GameTypeTeamFFA       = GameTypes[12]
	GameTypeTeamCTF       = GameTypes[13]
	GameTypeTvB           = GameTypes[15]
	GameTypeIronManLadder = GameTypes[16]
)

// GameTypeByID returns the GameType for a given ID.
// A new GameType with Unknown name is returned if one is not found
// for the given ID (preserving the unknown ID).
func GameTypeByID(ID uint16) *GameType {
	if int(ID) < len(GameTypes) {
		return GameTypes[ID]
	}
	return &GameType{UnknownEnum(ID), ID, "Unk"}
}

// PlayerType describes a player (slot) type.
type PlayerType struct {
	Enum

	// ID as it appears in replays
	ID byte
}

// PlayerTypes is an enumeration of the possible player types
var PlayerTypes = []*PlayerType{
	{Enum{"Inactive"}, 0x00},
	{Enum{"Computer"}, 0x01},
	{Enum{"Human"}, 0x02},
	{Enum{"Rescue Passive"}, 0x03},
	{Enum{"(Unused)"}, 0x04},
	{Enum{"Computer Controlled"}, 0x05},
	{Enum{"Open"}, 0x06},
	{Enum{"Neutral"}, 0x07},
	{Enum{"Closed"}, 0x08},
}

// Named player types
var (
	PlayerTypeInactive           = PlayerTypes[0]
	PlayerTypeComputer           = PlayerTypes[1]
	PlayerTypeHuman              = PlayerTypes[2]
	PlayerTypeRescuePassive      = PlayerTypes[3]
	PlayerTypeUnused             = PlayerTypes[4]
	PlayerTypeComputerControlled = PlayerTypes[5]
	PlayerTypeOpen               = PlayerTypes[6]
	PlayerTypeNeutral            = PlayerTypes[7]
	PlayerTypeClosed             = PlayerTypes[8]
)

// PlayerTypeByID returns the PlayerType for a given ID.
// A new PlayerType with Unknown name is returned if one is not found
// for the given ID (preserving the unknown ID).
func PlayerTypeByID(ID byte) *PlayerType {
	if int(ID) < len(PlayerTypes) {
		return PlayerTypes[ID]
	}
	return &PlayerType{UnknownEnum(ID), ID}
}

// Race describes a race.
type Race struct {
	Enum

	// ID as it appears in replays
	ID byte

	// ShortName is a shorter name
	ShortName string

	// Letter is the letter of the race (first letter of its name)
	Letter rune
}

// Races is an enumeration of the possible races
var Races = []*Race{
	{Enum{"Zerg"}, 0x00, "zerg", 'Z'},
	{Enum{"Terran"}, 0x01, "ran", 'T'},
	{Enum{"Protoss"}, 0x02, "toss", 'P'},
}

// Named races
var (
	RaceZerg    = Races[0]
	RaceTerran  = Races[1]
	RaceProtoss = Races[2]
)

// RaceByID returns the Race for a given ID.
// A new Race with Unknown name is returned if one is not found
// for the given ID (preserving the unknown ID).
func RaceByID(ID byte) *Race {
	if int(ID) < len(Races) {
		return Races[ID]
	}
	return &Race{UnknownEnum(ID), ID, "Unk", 'U'}
}

// Color describes a color.
type Color struct {
	Enum

	// ID as it appears in replays
	ID uint32

	// RGB is the red, green, blue component of the color
	RGB uint32
}

// Colors is an enumeration of the possible colors
var Colors = []*Color{
	{Enum{"Red"}, 0x00, 0xf40404},
	{Enum{"Blue"}, 0x01, 0x0c48cc},
	{Enum{"Teal"}, 0x02, 0x2cb494},
	{Enum{"Purple"}, 0x03, 0x88409c},
	{Enum{"Orange"}, 0x04, 0xf88c14},
	{Enum{"Brown"}, 0x05, 0x703014},
	{Enum{"White"}, 0x06, 0xcce0d0},
	{Enum{"Yellow"}, 0x07, 0xfcfc38},
	{Enum{"Green"}, 0x08, 0x088008},
	{Enum{"Pale Yellow"}, 0x09, 0xfcfc7c},
	{Enum{"Tan"}, 0x0a, 0xecc4b0},
	{Enum{"Aqua"}, 0x0b, 0x4068d4},
	{Enum{"Pale Green"}, 0x0c, 0x74a47c},
	{Enum{"Blueish Grey"}, 0x0d, 0x9090b8},
	{Enum{"Pale Yellow2"}, 0x0e, 0xfcfc7c},
	{Enum{"Cyan"}, 0x0f, 0x00e4fc},
}

// Named colors
var (
	ColorRed         = Colors[0]
	ColorBlue        = Colors[1]
	ColorTeal        = Colors[2]
	ColorPurple      = Colors[3]
	ColorOrange      = Colors[4]
	ColorBrown       = Colors[5]
	ColorWhite       = Colors[6]
	ColorYellow      = Colors[7]
	ColorGreen       = Colors[8]
	ColorPaleYellow  = Colors[9]
	ColorTan         = Colors[10]
	ColorAqua        = Colors[11]
	ColorPaleGreen   = Colors[12]
	ColorBlueishGrey = Colors[13]
	ColorPaleYellow2 = Colors[14] // Same as the other with same name
	ColorCyan        = Colors[15]
)

// ColorByID returns the Color for a given ID.
// A new Color with Unknown name is returned if one is not found
// for the given ID (preserving the unknown ID).
func ColorByID(ID uint32) *Color {
	if int(ID) < len(Colors) {
		return Colors[ID]
	}
	return &Color{UnknownEnum(ID), ID, 0}
}

// TileSet describes a tile set.
type TileSet struct {
	Enum

	// ID as it appears in replays
	ID uint16
}

// TileSets is an enumeration of the possible tile sets
var TileSets = []*TileSet{
	{Enum{"Badlands"}, 0x00},
	{Enum{"Space Platform"}, 0x01},
	{Enum{"Installation"}, 0x02},
	{Enum{"Ashworld"}, 0x03},
	{Enum{"Jungle"}, 0x04},
	{Enum{"Desert"}, 0x05},
	{Enum{"Arctic"}, 0x06},
	{Enum{"Twilight"}, 0x07},
}

// Named tile sets
var (
	TileSetBadlands      = TileSets[0]
	TileSetSpacePlatform = TileSets[1]
	TileSetInstallation  = TileSets[2]
	TileSetAshworld      = TileSets[3]
	TileSetJungle        = TileSets[4]
	TileSetDesert        = TileSets[5]
	TileSetArctic        = TileSets[6]
	TileSetTwilight      = TileSets[7]
)

// TileSetByID returns the TileSet for a given ID.
// A new TileSet with Unknown name is returned if one is not found
// for the given ID (preserving the unknown ID).
func TileSetByID(ID uint16) *TileSet {
	if int(ID) < len(TileSets) {
		return TileSets[ID]
	}
	return &TileSet{UnknownEnum(ID), ID}
}
