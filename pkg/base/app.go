package base

type (
	Initialize func(app App) *error

	App interface {
		Name() string
	}
)