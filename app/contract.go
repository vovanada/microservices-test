package app

type Service interface {
	Start(addr string) error
}
