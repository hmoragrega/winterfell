package game

// Result represents the game result
type Result bool

// Win result obtained if the player wins
const Win = Result(true)

// Lose result obtained if the player loses
const Lose = Result(false)

// Engine the engine that runs all the player games
type Engine interface {
	StartGame(playerName string) error
	Stop()
	Shoot(x int, y int) (bool, error)
	EnemyPosition() chan *Position
	GameOver() chan Result
	GetPlayerName() string
	GetEnemyName() string
}
