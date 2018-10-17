package model

type Memory struct {
	TotalKB  uint64 `json:"totalKB"`
	FreeKB   uint64 `json:"freeKB"`
	UsedKB   uint64 `json:"usedKB"`
	SharedKB uint64 `json:"sharedKB"`
	BufferKB uint64 `json:"bufferKB"`
	CacheKB  uint64 `json:"cacheKB"`
	SwapKB   uint64 `json:"swapKB"`
}
