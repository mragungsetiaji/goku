package goku

type Application interface {
	AddJob(j ...Job) error
	Enqueue(q string, p Payload) error
	GetAppName() string
	Start() error
}

type Options struct {
	PollFreq   int
	NumWorkers int
}

const (
	POLLFREQ   = 10
	NUMWORKERS = 3
)
