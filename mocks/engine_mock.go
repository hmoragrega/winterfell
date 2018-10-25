package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	"github.com/hmoragrega/winterfell/game"
)

// MockEngine is a mock of Engine interface
type MockEngine struct {
	ctrl     *gomock.Controller
	recorder *MockEngineMockRecorder
}

// MockEngineMockRecorder is the mock recorder for MockEngine
type MockEngineMockRecorder struct {
	mock *MockEngine
}

// NewMockEngine creates a new mock instance
func NewMockEngine(ctrl *gomock.Controller) *MockEngine {
	mock := &MockEngine{ctrl: ctrl}
	mock.recorder = &MockEngineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEngine) EXPECT() *MockEngineMockRecorder {
	return m.recorder
}

// StartGame mocks base method
func (m *MockEngine) StartGame(playerName string) error {
	ret := m.ctrl.Call(m, "StartGame", playerName)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartGame indicates an expected call of StartGame
func (mr *MockEngineMockRecorder) StartGame(playerName interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartGame", reflect.TypeOf((*MockEngine)(nil).StartGame), playerName)
}

// Stop mocks base method
func (m *MockEngine) Stop() {
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop
func (mr *MockEngineMockRecorder) Stop() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockEngine)(nil).Stop))
}

// Shoot mocks base method
func (m *MockEngine) Shoot(x, y int) (bool, error) {
	ret := m.ctrl.Call(m, "Shoot", x, y)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Shoot indicates an expected call of Shoot
func (mr *MockEngineMockRecorder) Shoot(x, y interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shoot", reflect.TypeOf((*MockEngine)(nil).Shoot), x, y)
}

// EnemyPosition mocks base method
func (m *MockEngine) EnemyPosition() chan *game.Position {
	ret := m.ctrl.Call(m, "EnemyPosition")
	ret0, _ := ret[0].(chan *game.Position)
	return ret0
}

// EnemyPosition indicates an expected call of EnemyPosition
func (mr *MockEngineMockRecorder) EnemyPosition() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnemyPosition", reflect.TypeOf((*MockEngine)(nil).EnemyPosition))
}

// GameOver mocks base method
func (m *MockEngine) GameOver() chan game.Result {
	ret := m.ctrl.Call(m, "GameOver")
	ret0, _ := ret[0].(chan game.Result)
	return ret0
}

// GameOver indicates an expected call of GameOver
func (mr *MockEngineMockRecorder) GameOver() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GameOver", reflect.TypeOf((*MockEngine)(nil).GameOver))
}

// GetPlayerName mocks base method
func (m *MockEngine) GetPlayerName() string {
	ret := m.ctrl.Call(m, "GetPlayerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetPlayerName indicates an expected call of GetPlayerName
func (mr *MockEngineMockRecorder) GetPlayerName() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPlayerName", reflect.TypeOf((*MockEngine)(nil).GetPlayerName))
}

// GetEnemyName mocks base method
func (m *MockEngine) GetEnemyName() string {
	ret := m.ctrl.Call(m, "GetEnemyName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetEnemyName indicates an expected call of GetEnemyName
func (mr *MockEngineMockRecorder) GetEnemyName() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEnemyName", reflect.TypeOf((*MockEngine)(nil).GetEnemyName))
}
