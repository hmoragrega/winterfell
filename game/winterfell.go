package game

import (
	"fmt"
	"log"
	"sync"
)

// @TODO This variabls should parameters coming from env variables, flags, etc.
var boardWidth = 10
var boardHeight = 30
var milliSecondsPerCell = 2000 // How many seconds before an enemy moves again
var cellMovement = 1           // How many cell moves an enemy on each turn
var enemyName = "knight-king"  // Default enemy name

// WinterIsComing variation of the game with zombies as enemies
type WinterIsComing struct {
	board         *Board
	player        *Player
	enemy         Enemy
	gameover      chan Result
	stop          chan struct{}
	enemyPosition chan *Position
	rm            sync.RWMutex
}

// NewWinterIsComingEngine returns the winter is coming version of the game
// We could have diferent variations with diferent enemies
func NewWinterIsComingEngine() *WinterIsComing {
	return &WinterIsComing{
		board:         NewBoard(boardWidth, boardHeight),
		gameover:      make(chan Result),
		enemyPosition: make(chan *Position),
	}
}

// StartGame a new game for the given player
func (g *WinterIsComing) StartGame(playerName string) error {
	if g.hasStarted() {
		return fmt.Errorf("The game has already started")
	}

	player, err := NewPlayer(playerName)
	if err != nil {
		return err
	}

	g.rm.Lock()
	g.stop = make(chan struct{})
	g.rm.Unlock()

	g.player = player
	g.enemy = NewZombie(enemyName, &Position{0, 0}, milliSecondsPerCell)

	go g.loop()

	return nil
}

// Stop stops the current game
func (g *WinterIsComing) Stop() {
	log.Println("Stopping game")
	if g.hasStarted() {
		g.stop <- struct{}{}
		log.Println("Game stopped")
	}
}

// GameOver this channels gets notified once the games has finsihed with the result
func (g *WinterIsComing) GameOver() chan Result {
	return g.gameover
}

// Shoot calculates if a shot has hit the enemy
func (g *WinterIsComing) Shoot(x int, y int) (bool, error) {
	if !g.hasStarted() {
		return false, fmt.Errorf("The games has not started yet, use the START command")
	}

	hit := g.enemy.IsCurrentPosition(NewPosition(x, y))

	if hit {
		g.Stop()
		g.gameOver(Win)
	}

	return hit, nil
}

// GetPlayerName Retrieves the current player name
// Note: for the sake of simplicity it doesn't return an error if the games is stopped
func (g *WinterIsComing) GetPlayerName() string {
	if g.player == nil {
		return ""
	}

	return g.player.GetName()
}

// GetEnemyName Retrieves the current enemy name
// Note: for the sake of simplicity it doesn't return an error if the games is stopped
func (g *WinterIsComing) GetEnemyName() string {
	if g.enemy == nil {
		return ""
	}

	return g.enemy.GetName()
}

// EnemyPosition notifies changes in the enemy position
func (g *WinterIsComing) EnemyPosition() chan *Position {
	return g.enemyPosition
}

func (g *WinterIsComing) loop() {
	g.enemyPosition <- g.enemy.GetCurrentPosition()
	enemyTurn := g.enemy.StartMoving()

	defer func() {
		g.enemy.StopMoving()

		g.rm.Lock()
		close(g.stop)
		g.stop = nil
		g.rm.Unlock()
	}()

	for {
		select {
		case <-enemyTurn:
			g.moveEnemy()
			if g.isEnemyWinner() {
				g.gameOver(Lose)
				return
			}
		case <-g.stop:
			log.Println("Ending game loop")
			return
		}
	}
}

func (g *WinterIsComing) moveEnemy() {
	if err := g.enemy.Move(g.board); err != nil {
		log.Println(err.Error())
	}

	g.enemyPosition <- g.enemy.GetCurrentPosition()
}

func (g *WinterIsComing) isEnemyWinner() bool {
	return g.enemy.GetCurrentPosition().GetX() == g.board.width
}

func (g *WinterIsComing) hasStarted() bool {
	g.rm.RLock()
	defer g.rm.RUnlock()

	return g.stop != nil
}

func (g *WinterIsComing) gameOver(result Result) {
	log.Printf("Game over reached %v", result)
	g.gameover <- result
}
