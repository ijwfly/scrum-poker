package app

import "scrum-poker/internal/poker"

type App struct {
	Poker poker.PokerSessionController
}

func NewApp(poker poker.PokerSessionController) *App {
	return &App{
		Poker: poker,
	}
}
