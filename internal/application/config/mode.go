package config

type mode struct{}

func (mode) Local() string {
	return "local"
}

func (mode) Dev() string {
	return "dev"
}

func (mode) Prod() string {
	return "prod"
}

var Mode = mode{}
