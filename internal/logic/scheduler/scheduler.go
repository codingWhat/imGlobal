package scheduler

type Scheduler struct {
	JobChan chan *Job
	workerPool chan *Worker
	workerSize int
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		JobChan: make(chan *Job, 1000),
		workerPool: make(chan *Worker, 0),
		workerSize: 10,
	}
}

var G_scheduler *Scheduler

func InitScheduler() {
	//注册handler
	RegLogicHandlers()

	G_scheduler = NewScheduler()

	for i := 0; i < G_scheduler.workerSize; i++ {
		worker := NewWorker(G_scheduler)
		go worker.Run()
	}
}

func (s *Scheduler) GetWorker() *Worker {
		return <-s.workerPool
}


