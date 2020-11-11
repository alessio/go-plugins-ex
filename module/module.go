package module

type Interface interface {
	Name() string
	Permissions() string
	Configure() error
	Start() error
	Stop() error
}

type Getter func() Interface
