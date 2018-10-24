package main

import (
	"fmt"
	"strconv"

	"github.com/hmoragrega/winterfell/game"
)

const start = "START"
const shoot = "SHOOT"
const boom = "BOOM"
const walk = "WALK"
const gameOver = "GAMEOVER"
const gameError = "ERROR"

func executeGameCommand(g game.Engine, command string, parameters []string) string {
	switch command {
	case start:
		return executeStart(g, parameters)
	case shoot:
		return executeShoot(g, parameters)
	}

	return formatError("Unknown command %s", command)
}

func format(tag string, message string, extra ...interface{}) string {
	return fmt.Sprintf("%s %s\n", tag, fmt.Sprintf(message, extra...))
}

func formatError(message string, extra ...interface{}) string {
	return format(gameError, message, extra...)
}

func executeStart(g game.Engine, parameters []string) string {
	if len(parameters) != 1 {
		return formatError("Wrong number of parameters, pass the player name")
	}

	if err := g.StartGame(parameters[0]); err != nil {
		return formatError(err.Error())
	}

	return ""
}

func executeShoot(g game.Engine, parameters []string) string {
	if len(parameters) != 2 {
		return formatError("Wrong number of parameters, you must provide the shoot coordinates: SHOOT 1 3")
	}

	x, err := strconv.Atoi(parameters[0])
	if err != nil {
		return formatError("The first parameter X must be an integer: SHOOT 1 3")
	}

	y, err := strconv.Atoi(parameters[1])
	if err != nil {
		return formatError("The second parameter Y must be an integer: SHOOT 1 3")
	}

	hit, err := g.Shoot(x, y)
	if err != nil {
		return formatError(err.Error())
	}

	if hit {
		return format(boom, "%s 1 %s", g.GetPlayerName(), g.GetEnemyName())
	}

	return format(boom, "%s 0", g.GetPlayerName())
}

func getEnemyPositionMessage(g game.Engine, p *game.Position) string {
	return format(walk, "%s %d %d", g.GetEnemyName(), p.GetX(), p.GetY())
}

func getGameOverMessage(result game.Result) string {
	switch result {
	case game.Win:
		return format(gameOver, "You win!!")
	case game.Lose:
		return format(gameOver, "You lose!!")
	}

	return format(gameOver, "Unkown result")
}
