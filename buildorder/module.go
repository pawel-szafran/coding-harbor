package buildorder

type (
	Module struct {
		Name ModuleName
		Deps []ModuleName
	}
	ModuleName string
)
