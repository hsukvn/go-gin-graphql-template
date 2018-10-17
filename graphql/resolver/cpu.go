package resolver

import (
	"context"
	"fmt"

	linuxproc "github.com/c9s/goprocinfo/linux"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/hsukvn/go-gin-graphql-template/graphql/model"
	"github.com/hsukvn/go-gin-graphql-template/graphql/scalar"
)

const procStatPath = "/proc/stat"

type cpuArgs struct {
	ID string
}

func (r *Resolver) CPU(ctx context.Context, args cpuArgs) (*cpuResolver, error) {
	stat, err := linuxproc.ReadStat(procStatPath)
	if err != nil {
		return nil, fmt.Errorf("cpu: Fail to open (%v), err: (%v)", procStatPath, err)
	}

	for i, s := range stat.CPUStats {
		if s.Id == args.ID {
			return &cpuResolver{
				cpu: model.NewCPU(&stat.CPUStats[i]),
			}, nil
		}
	}

	return nil, fmt.Errorf("cpu: ID (%v) is not exist", args.ID)
}

func (r *Resolver) CPUs(ctx context.Context) (*[]*cpuResolver, error) {
	stat, err := linuxproc.ReadStat(procStatPath)
	if err != nil {
		return nil, fmt.Errorf("cpu: Fail to open (%v), err: (%v)", procStatPath, err)
	}

	cpus := make([]*cpuResolver, 0)
	for i, _ := range stat.CPUStats {
		cpus = append(cpus, &cpuResolver{
			cpu: model.NewCPU(&stat.CPUStats[i]),
		})
	}

	return &cpus, nil
}

type cpuResolver struct {
	cpu *model.CPU
}

func (r *cpuResolver) ID() *graphql.ID {
	id := graphql.ID(r.cpu.ID)
	return &id
}

func (r *cpuResolver) Total() *scalar.Uint64 {
	total := scalar.Uint64(r.cpu.TotalTick)
	return &total
}

func (r *cpuResolver) User() *scalar.Uint64 {
	user := scalar.Uint64(r.cpu.UserTick)
	return &user
}

func (r *cpuResolver) System() *scalar.Uint64 {
	system := scalar.Uint64(r.cpu.SystemTick)
	return &system
}

func (r *cpuResolver) Idle() *scalar.Uint64 {
	idle := scalar.Uint64(r.cpu.IdleTick)
	return &idle
}

func (r *cpuResolver) IOWait() *scalar.Uint64 {
	iowait := scalar.Uint64(r.cpu.IOWaitTick)
	return &iowait
}
