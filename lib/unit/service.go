package unit

import (
	"fmt"
	"os/exec"
	"strconv"
)

type Service struct {
	Name     string
	Property Property
}

func NewService(name string) (*Service, error) {
	p, err := NewProperty(name)
	if err != nil {
		return nil, err
	}

	n, err := p.GetName()
	if err != nil {
		return nil, err
	}

	return &Service{
		Name:     n,
		Property: p,
	}, nil
}

func (s *Service) GetMainPID() (uint64, error) {
	if p, ok := s.Property["MainPID"]; ok {
		pid, err := strconv.ParseUint(p, 10, 64)

		if err != nil {
			return 0, err
		}

		return pid, nil
	}

	return 0, fmt.Errorf("property MainPID does not exist")
}

func (s *Service) GetActiveState() (ActiveState, error) {
	if st, ok := s.Property["ActiveState"]; ok {
		if state, ok := MapActiveState[st]; ok {
			return state, nil
		}

		return ActiveStateError, fmt.Errorf("property ActiveState unrecognized")
	}

	return ActiveStateError, fmt.Errorf("property ActiveState does not exist")
}

func (s *Service) GetUnitFileState() (UnitFileState, error) {
	if st, ok := s.Property["UnitFileState"]; ok {
		if state, ok := MapUnitFileState[st]; ok {
			return state, nil
		}

		return UnitFileStateError, fmt.Errorf("property UnitFileState unrecognized")
	}

	return UnitFileStateError, fmt.Errorf("property UnitFileState does not exist")
}

func (s *Service) Start() error {
	state, err := s.GetActiveState()
	if err != nil {
		return err
	}

	if state == ActiveStateActive {
		return nil
	}

	err = exec.Command("/bin/systemctl", "start", s.Name).Run()
	if err != nil {
		return err
	}

	s.Property, err = NewProperty(s.Name)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) Stop() error {
	state, err := s.GetActiveState()
	if err != nil {
		return err
	}

	if state == ActiveStateInactive {
		return nil
	}

	err = exec.Command("/bin/systemctl", "stop", s.Name).Run()
	if err != nil {
		return err
	}

	s.Property, err = NewProperty(s.Name)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) Enable() error {
	state, err := s.GetUnitFileState()
	if err != nil {
		return err
	}

	if state == UnitFileStateEnabled {
		return nil
	}

	err = exec.Command("/bin/systemctl", "enable", s.Name).Run()
	if err != nil {
		return err
	}

	s.Property, err = NewProperty(s.Name)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) Disable() error {
	state, err := s.GetUnitFileState()
	if err != nil {
		return err
	}

	if state == UnitFileStateDisabled {
		return nil
	}

	err = exec.Command("/bin/systemctl", "disable", s.Name).Run()
	if err != nil {
		return err
	}

	s.Property, err = NewProperty(s.Name)
	if err != nil {
		return err
	}

	return nil
}
