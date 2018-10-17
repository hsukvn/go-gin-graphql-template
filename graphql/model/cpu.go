package model

import (
	linuxproc "github.com/c9s/goprocinfo/linux"
)

type CPU struct {
	ID         string `json:"id"`
	TotalTick  uint64 `json:"totalTick"`
	UserTick   uint64 `json:"userTick"`
	SystemTick uint64 `json:"systemTick"`
	IdleTick   uint64 `json:"idleTick"`
	IOWaitTick uint64 `json:"iowaitTick"`
}

func NewCPU(cs *linuxproc.CPUStat) *CPU {
	total := cs.User + cs.Nice + cs.System + cs.Idle + cs.IOWait +
		cs.IRQ + cs.SoftIRQ + cs.Steal
	user := cs.User + cs.Nice
	sys := cs.System + cs.IRQ + cs.SoftIRQ + cs.Steal
	idle := cs.Idle
	iowait := cs.IOWait

	return &CPU{
		ID:         cs.Id,
		TotalTick:  total,
		UserTick:   user,
		SystemTick: sys,
		IdleTick:   idle,
		IOWaitTick: iowait,
	}
}
