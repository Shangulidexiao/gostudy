package ch7

import (
	"flag"
	"fmt"
)

//Game game struct
type Game struct {
	name string
}

type gameFlag struct {
	Game
}

func (g *Game) String() string {
	return g.name
}

//Set is Game Flag function
func (g *Game) Set(s string) error {
	var name string
	n, err := fmt.Sscanf(s, "%s", &name)
	if n < 1 {
		return fmt.Errorf("name is empty")
	}
	if err != nil {
		return fmt.Errorf("read cmd arg err :%q", err)
	}
	g.name = name
	return nil
}

//GameFlag is defined CommandLine Flag
func GameFlag(name string, value Game, usage string) *Game {
	f := gameFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Game
}
