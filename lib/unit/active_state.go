package unit

type ActiveState int

const (
	ActiveStateError ActiveState = iota - 1
	ActiveStateInactive
	ActiveStateActive
	ActiveStateDeactivating
	ActiveStateActivating
	ActiveStateReloading
	ActiveStateFailed
)

var MapActiveState = map[string]ActiveState{
	"inactive":     ActiveStateInactive,
	"active":       ActiveStateActive,
	"deactivating": ActiveStateDeactivating,
	"activating":   ActiveStateActivating,
	"reloading":    ActiveStateReloading,
	"failed":       ActiveStateFailed,
}
