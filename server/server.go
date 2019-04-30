package server

import (
	"net"

	"github.com/makyo/mugz/game"
)

type server struct {
	game *game.Game
	ln   net.Listener
}

func New(g *game.Game, addr string) (*server, error) {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, error
	}

	s := &server{
		game: g,
		ln:   ln,
	}
	return s, nil
}
