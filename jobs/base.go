package jobs

type Job interface {
	Run() error
	GetID() string
}
