package app

type App interface {
	Serve()
	Stop()
}

type Options struct {
}
