package scheduler

type Worker struct {
	JobChan chan *Job
	RetChan chan bool
	scheduler *Scheduler
}

func NewWorker(s *Scheduler) *Worker {
	return &Worker{
		JobChan: make(chan *Job, 1000),
		scheduler: s,
		RetChan:make(chan  bool, 100),
	}
}

func (w *Worker) Run() {
	for {
		w.scheduler.workerPool <- w
		select {
		case job := <-w.JobChan:
			err := getJobHandler(job.Handler)(job.Params)
			if err == nil {
				//通知consumer. mark offset
				w.RetChan <- true
			}  else {
				w.RetChan <- false
			}
		}
	}
}