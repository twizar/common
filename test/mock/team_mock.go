// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/client/team.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	dto "github.com/twizar/common/pkg/dto"
)

// MockTeams is a mock of Teams interface.
type MockTeams struct {
	ctrl     *gomock.Controller
	recorder *MockTeamsMockRecorder
}

// MockTeamsMockRecorder is the mock recorder for MockTeams.
type MockTeamsMockRecorder struct {
	mock *MockTeams
}

// NewMockTeams creates a new mock instance.
func NewMockTeams(ctrl *gomock.Controller) *MockTeams {
	mock := &MockTeams{ctrl: ctrl}
	mock.recorder = &MockTeamsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTeams) EXPECT() *MockTeamsMockRecorder {
	return m.recorder
}

// AllTeams mocks base method.
func (m *MockTeams) AllTeams() ([]dto.Team, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllTeams")
	ret0, _ := ret[0].([]dto.Team)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllTeams indicates an expected call of AllTeams.
func (mr *MockTeamsMockRecorder) AllTeams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllTeams", reflect.TypeOf((*MockTeams)(nil).AllTeams))
}

// SearchTeams mocks base method.
func (m *MockTeams) SearchTeams(minRating float64, leagues []string, orderBy string, limit int) ([]dto.Team, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchTeams", minRating, leagues, orderBy, limit)
	ret0, _ := ret[0].([]dto.Team)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchTeams indicates an expected call of SearchTeams.
func (mr *MockTeamsMockRecorder) SearchTeams(minRating, leagues, orderBy, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchTeams", reflect.TypeOf((*MockTeams)(nil).SearchTeams), minRating, leagues, orderBy, limit)
}

// TeamsByID mocks base method.
func (m *MockTeams) TeamsByID(ids []string) ([]dto.Team, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TeamsByID", ids)
	ret0, _ := ret[0].([]dto.Team)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TeamsByID indicates an expected call of TeamsByID.
func (mr *MockTeamsMockRecorder) TeamsByID(ids interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TeamsByID", reflect.TypeOf((*MockTeams)(nil).TeamsByID), ids)
}
