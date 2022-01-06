package scheduler

type ReadResult struct {
	Data  interface{}
	Error error
}

type WriteResult struct {
	Status int
	Error  error
}
