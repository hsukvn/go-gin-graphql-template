package unit

type UnitFileState int

const (
	UnitFileStateError UnitFileState = iota - 1
	UnitFileStateDisabled
	UnitFileStateEnabled
	UnitFileStateStatic
	UnitFileStateMasked
	UnitFileStateLinked
)

var MapUnitFileState = map[string]UnitFileState{
	"disabled":        UnitFileStateDisabled,
	"enabled":         UnitFileStateEnabled,
	"enabled-runtime": UnitFileStateEnabled,
	"static":          UnitFileStateStatic,
	"masked":          UnitFileStateMasked,
	"masked-runtime":  UnitFileStateMasked,
	"linked":          UnitFileStateLinked,
	"linked-runtime":  UnitFileStateLinked,
}
