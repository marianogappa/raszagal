// This file contains the types describing the replay header.

package rep

import (
	"bytes"
	"fmt"
	"time"

	"github.com/icza/screp/rep/repcore"
)

// Header models the replay header.
type Header struct {
	// Engine used to play the game and save the replay
	Engine *repcore.Engine

	// Frames is the number of frames. There are approximately ~23.81 frames in
	// a second. (1 frame = 0.042 second to be exact).
	Frames repcore.Frame

	// StartTime is the timestamp when the game started
	StartTime time.Time

	// Title is the game name / title
	Title string

	// Size of the map
	MapWidth, MapHeight uint16

	// AvailSlotsCount is the number of available slots
	AvailSlotsCount byte

	// Speed is the game speed
	Speed *repcore.Speed

	// Type is the game type
	Type *repcore.GameType

	// SubType indicates the size of the "Home" team.
	// For example, in case of 3v5 this is 3, in case of 7v1 this is 7.
	SubType uint16

	// Host is the game creator's name.
	Host string

	// Map name
	Map string

	// Slots contains all players of the game (including open/closed slots)
	Slots []*Player `json:"-"`

	// OrigPlayers contains the actual ("real") players of the game
	// in the order recorded in the replay.
	OrigPlayers []*Player `json:"-"`

	// Players contains the actual ("real") players of the game
	// in team order.
	Players []*Player

	// PIDPlayers maps from player ID to Player.
	PIDPlayers map[byte]*Player `json:"-"`
}

// Duration returns the game duration.
func (h *Header) Duration() time.Duration {
	return h.Frames.Duration()
}

// MapSize returns the map size in widthxheight format, e.g. "64x64".
func (h *Header) MapSize() string {
	return fmt.Sprint(h.MapWidth, "x", h.MapHeight)
}

// Matchup returns the matchup, the race letters of players in team order,
// inserting 'v' between different teams, e.g. "PvT" or "PTZvZTP".
func (h *Header) Matchup() string {
	m := make([]rune, 0, 9)
	var prevTeam byte
	for i, p := range h.Players {
		if i > 0 && p.Team != prevTeam {
			m = append(m, 'v')
		}
		m = append(m, p.Race.Letter)
		prevTeam = p.Team
	}
	return string(m)
}

// PlayerNames returns a comma separated list of player names in team order,
// inserting " VS " between different teams.
func (h *Header) PlayerNames() string {
	buf := &bytes.Buffer{}
	var prevTeam byte
	for i, p := range h.Players {
		if i > 0 {
			if p.Team != prevTeam {
				buf.WriteString(" VS ")
			} else {
				buf.WriteString(", ")
			}
		}
		buf.WriteString(p.Name)
		prevTeam = p.Team
	}
	return buf.String()
}

// Player represents a player of the game.
type Player struct {
	// SlotID is the slot ID
	SlotID uint16

	// ID of the player
	ID byte

	// Type is the player type
	Type *repcore.PlayerType

	// Race of the player
	Race *repcore.Race

	// Team of the player
	Team byte

	// Name of the player
	Name string

	// Color of the player
	Color *repcore.Color
}
