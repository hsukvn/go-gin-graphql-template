package resolver

import (
	"context"
	"fmt"

	linuxproc "github.com/c9s/goprocinfo/linux"
	"github.com/hsukvn/go-gin-graphql-template/graphql/model"
	"github.com/hsukvn/go-gin-graphql-template/graphql/scalar"
)

const procMeminfoPath = "/proc/meminfo"

func (r *Resolver) Memory(ctx context.Context) (*memResolver, error) {
	info, err := linuxproc.ReadMemInfo(procMeminfoPath)
	if err != nil {
		return nil, fmt.Errorf("memory: Fail to open (%v), err: (%v)", procMeminfoPath, err)
	}

	/*
		https://github.com/hishamhm/htop/blob/8af4d9f453ffa2209e486418811f7652822951c6/linux/LinuxProcessList.c#L802-L833
		https://github.com/hishamhm/htop/blob/1f3d85b6174f690a7e354bbadac19404d5e75e78/linux/Platform.c#L198-L208
	*/
	totalUsed := info.MemTotal - info.MemFree
	buffer := info.Buffers
	cache := info.Cached + info.SReclaimable - info.Shmem
	used := totalUsed - (buffer + cache)
	swap := info.SwapTotal - info.SwapFree

	mem := model.Memory{
		TotalKB:  info.MemTotal,
		FreeKB:   info.MemFree,
		UsedKB:   used,
		SharedKB: info.Shmem,
		BufferKB: buffer,
		CacheKB:  cache,
		SwapKB:   swap,
	}

	return &memResolver{memory: &mem}, nil
}

type memResolver struct {
	memory *model.Memory
}

func (r *memResolver) Total() *scalar.Uint64 {
	total := scalar.Uint64(r.memory.TotalKB)
	return &total
}

func (r *memResolver) Free() *scalar.Uint64 {
	free := scalar.Uint64(r.memory.FreeKB)
	return &free
}

func (r *memResolver) Used() *scalar.Uint64 {
	used := scalar.Uint64(r.memory.UsedKB)
	return &used
}

func (r *memResolver) Shared() *scalar.Uint64 {
	shared := scalar.Uint64(r.memory.SharedKB)
	return &shared
}

func (r *memResolver) Buffer() *scalar.Uint64 {
	buffer := scalar.Uint64(r.memory.BufferKB)
	return &buffer
}

func (r *memResolver) Cache() *scalar.Uint64 {
	cache := scalar.Uint64(r.memory.CacheKB)
	return &cache
}

func (r *memResolver) Swap() *scalar.Uint64 {
	swap := scalar.Uint64(r.memory.SwapKB)
	return &swap
}
