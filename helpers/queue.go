package helpers

// JobQueue 작업 요청을 보낼 수있는 버퍼링된 채널.
var JobQueue chan Queuable

// Queuable Queuable Job의 인터페이스
type Queuable interface {
	Handle() error
}

//Dispatcher ... worker dispatcher
type Dispatcher struct {
	maxWorkers int
	WorkerPool chan chan Queuable
	Workers    []Worker
}

type Worker struct {
	Name       string
	WorkerPool chan chan Queuable
	JobChannel chan Queuable
	quit       chan bool
}
