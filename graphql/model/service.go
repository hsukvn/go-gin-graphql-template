package model

import (
	"github.com/hsukvn/go-gin-graphql-template/lib/unit"
)

type Service struct {
	Name          string `json:"name"`
	MainPID       uint64 `json: main_pid`
	ActiveState   int32  `json:"active_state"`
	UnitFileState int32  `json:"unit_file_state"`
}

func NewService(s *unit.Service) (*Service, error) {
	pid, err := s.GetMainPID()
	if err != nil {
		return nil, err
	}

	activeState, err := s.GetActiveState()
	if err != nil {
		return nil, err
	}

	unitFileState, err := s.GetUnitFileState()
	if err != nil {
		return nil, err
	}

	return &Service{
		Name:          s.Name,
		MainPID:       pid,
		ActiveState:   int32(activeState),
		UnitFileState: int32(unitFileState),
	}, nil
}
