package main

import (
	"encoding/json"
	"io"
)

// Unique ID for a game type, so hotel can multiplex across games.
type GameIDType string

// Unique ID for a game server instance.
type ServerIDType string

// An connectable game server instance.
type GameServer struct {
	ID         ServerIDType `json:"id"`
	GameID     GameIDType   `json:"gameId"`
	Name       string       `json:"name"`
	Host       string       `json:"host"`
	Port       string       `json:"port"`
	NumPlayers int          `json:"numPlayers"`
	MaxPlayers int          `json:"maxPlayers"`
}

// Decode a JSON version of GameServer, validate it, and return it.
// If any steps fail, returns an error.
func DecodeAndValidateServer(reader io.Reader) (GameServer, error) {
	decoder := json.NewDecoder(reader)
	var server GameServer
	err := decoder.Decode(&server)
	if err != nil {
		return server, err
	}
	err = server.Validate()
	if err != nil {
		return server, err
	}
	return server, nil
}

func (s *GameServer) Validate() error {
	return nil
}

func (s *GameServer) UpdateFrom(other *GameServer) error {
	return nil
}
