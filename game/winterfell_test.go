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

	// Ensure games stops after win
	go func() {
		wg.Add(1)
		<-w.stop
		wg.Done()
	}()

	// Hit when ready it!
	go func() {
		wg.Add(1)
		p := <-w.EnemyPosition()
		hit, err := w.Shoot(p.GetX(), p.GetY())
		if assert.Nil(t, err) && assert.True(t, hit) {
			wg.Done()
		}
	}()

	go func() {
		wg.Add(1)
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

	// This board will ensure a win in two moves for the zombie
	b := NewBoard(3, 1)
	expectedMoves := 2

	w := &WinterIsComing{
		board:         b,
		gameover:      make(chan Result),
		enemyPosition: make(chan *Position),
	}

	wg := &sync.WaitGroup{}

	go func() {
		wg.Add(1)
		for {
			<-w.EnemyPosition()
			expectedMoves--
			if expectedMoves == 0 {
				wg.Done()
			}
		}
	}()

	go func() {
		wg.Add(1)
		if assert.Equal(t, Lose, <-w.GameOver(), "We should have lost the game") {
			wg.Done()
		}
	}()

	w.StartGame("foo")

	wg.Wait()
	w.Stop()
}
