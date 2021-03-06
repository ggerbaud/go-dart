package common

import "time"

// Sector the representation of a position hit by the dart on the board
type Sector struct {
	Val int
	Pos int
}

// IsValid tells if a Sector really exists (Triple 25 is not by example)
func (s Sector) IsValid() bool {
	if s.Val > 0 && s.Val <= 20 {
		return s.Pos > 0 && s.Pos <= 3
	} else if s.Val == 25 {
		return s.Pos == 1 || s.Pos == 2
	}
	return false
}

// PlayerState the player data (name, score, rank, ...)
type PlayerState struct {
	Name   string
	Score  int
	Rank   int
	Histo  map[string]int
	Board  string
	Visits []Sector
}

// ByRank implements sort.Interface
type ByRank []PlayerState

func (r ByRank) Len() int {
	return len(r)
}
func (r ByRank) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}
func (r ByRank) Less(i, j int) bool {
	if r[i].Rank == 0 {
		return r[j].Rank == 0
	} else if r[j].Rank == 0 {
		return true
	} else {
		return r[i].Rank < r[j].Rank
	}
}

// ByScore implements sort.Interface
type ByScore []PlayerState

func (r ByScore) Len() int {
	return len(r)
}
func (r ByScore) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}
func (r ByScore) Less(i, j int) bool {
	return r[i].Score > r[j].Score
}

// State the state of a Game : INITIALIZING, READY, PLAYING or OVER
type State int

const (
	// INITIALIZING Initial State of a Game
	INITIALIZING State = iota
	// READY Ready means the game could be start, the required number of player is achieved
	READY
	// PLAYING game has start
	PLAYING
	// ONHOLD game is waiting for next player signal
	ONHOLD
	// OVER game has ended
	OVER
)

// GameState the structured state of a Game, with all players and game data
type GameState struct {
	Players       []PlayerState
	Ongoing       State
	CurrentPlayer int
	CurrentDart   int
	LastMsg       string
	LastSector    Sector
	Round         int
}

// NewGameState the GameState constructor
func NewGameState() *GameState {

	g := new(GameState)
	g.Ongoing = INITIALIZING
	g.Players = make([]PlayerState, 0, 4)

	return g
}

// GameStyle the representation of a Game variant
type GameStyle struct {
	Name    string
	Code    string
	Rules   string
	Options []GameOption
}

// GameOption the representation of a Game option
type GameOption struct {
	Name    string
	Type    string
	Desc    string
	Default interface{}
}

const (
	// HealthCheckDelay is heart beat frequency
	HealthCheckDelay = 2 * time.Second
	// HealthCheckTimeout is the delay after which a board is considered disconnected
	HealthCheckTimeout = HealthCheckDelay + 1*time.Second
)
