package game

import (
	"io/ioutil"
	"log"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	log.SetOutput(ioutil.Discard)
}

func TestStartGame(t *testing.T) {
	w := NewWinterIsComingEngine()

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		// Make sure the initial position is given when the game starts
		initialPosition := <-w.EnemyPosition()
		if assert.Equal(t, NewPosition(0, 0), initialPosition) {
			wg.Done()
		}
	}()

	w.StartGame("foo")
	assert.True(t, w.hasStarted())
	w.Stop()

	wg.Wait()
}

func TestStop(t *testing.T) {
	w := NewWinterIsComingEngine()
	wg := &sync.WaitGroup{}

	// Simulate a game start
	w.stop = make(chan struct{})
	go func() {
		wg.Add(1)
		<-w.stop
		wg.Done()
	}()

	w.Stop()

	// Wait until the stop signal is received
	wg.Wait()
}

func TestShootOnlyWorkingWhenGameHasStarted(t *testing.T) {
	w := NewWinterIsComingEngine()

	_, err := w.Shoot(0, 0)

	assert.Error(t, err)
}

func TestPlayerWinsIfHitsTheEnemy(t *testing.T) {
	w := NewWinterIsComingEngine()

	wg := &sync.WaitGroup{}
	wg.Add(2)

	// Hit when ready it!
	go func() {
		p := <-w.EnemyPosition()
		hit, err := w.Shoot(p.GetX(), p.GetY())
		if assert.Nil(t, err) && assert.True(t, hit) {
			wg.Done()
		}
	}()

	go func() {
		if assert.Equal(t, Win, <-w.GameOver(), "We should have win the game") {
			wg.Done()
		}
	}()

	w.StartGame("foo")
	wg.Wait()
}

func TestLosingAGame(t *testing.T) {
	// Let's make the game faster
	milliSecondsPerCell = 100

	// This board will ensure a win in one moves for the zombie
	b := NewBoard(2, 1)

	w := &WinterIsComing{
		board:         b,
		gameover:      make(chan Result),
		enemyPosition: make(chan *Position),
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		for {
			<-w.EnemyPosition()
		}
	}()

	go func() {
		result := <-w.GameOver()
		if assert.Equal(t, Lose, result, "We should have lost the game") {
			wg.Done()
		}
	}()

	w.StartGame("foo")

	wg.Wait()
}
