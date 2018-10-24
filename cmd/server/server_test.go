package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"sync"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/hmoragrega/winterfell/mocks"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	log.SetOutput(ioutil.Discard)
}

func TestBroadcastMessages(t *testing.T) {
	server, client := net.Pipe()
	message := "foo"

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		var buffer = make([]byte, 1024)
		if _, err := client.Read(buffer); err != nil {
			return
		}

		if assert.Equal(t, message, cleanMessage(buffer)) {
			wg.Done()
		}
	}()

	time.Sleep(1)
	broadcastMessage(server, message)

	wg.Wait()
}

func TestExecuteStartCommand(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	engine := mocks.NewMockEngine(mockCtrl)
	server, client := net.Pipe()
	stop := make(chan struct{})

	username := "foo"

	engine.EXPECT().StartGame(username).Return(nil).Times(1)
	engine.EXPECT().Stop().Return().Times(1)

	go executePlayerCommands(server, engine, stop)

	client.Write([]byte(fmt.Sprintf("START %s", username)))
	client.Close()

	<-stop
}

func TestExecuteShootCommand(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	engine := mocks.NewMockEngine(mockCtrl)
	server, client := net.Pipe()
	stop := make(chan struct{})

	x := 0
	y := 1
	hit := true
	username := "foo"
	enemy := "bar"
	expected := fmt.Sprintf("BOOM %s 1 %s", username, enemy)

	engine.EXPECT().Shoot(x, y).Return(hit, nil).Times(1)
	engine.EXPECT().GetPlayerName().Return(username).Times(1)
	engine.EXPECT().GetEnemyName().Return(enemy).Times(1)
	engine.EXPECT().Stop().Return().Times(1)

	go executePlayerCommands(server, engine, stop)

	client.Write([]byte(fmt.Sprintf("SHOOT %d %d", x, y)))

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		var buffer = make([]byte, 1024)
		if _, err := client.Read(buffer); err != nil {
			return
		}

		if assert.Equal(t, expected, cleanMessage(buffer)) {
			wg.Done()
		}
	}()

	// We'll wait until we receive the correct response from the server
	wg.Done()
	client.Close()
	<-stop
}
