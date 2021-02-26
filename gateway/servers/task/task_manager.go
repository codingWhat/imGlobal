package task

import (
	"time"
)

var taskManager map[string]func()

func Start() {
	taskManager = make(map[string]func())
	Register("clearConns", 3*time.Second, 30*time.Second, cleanConns, "", nil, nil)
	Register("regService", 1*time.Second, 60*time.Second, regServer, "", remServer, nil)

	for _, fun := range taskManager {
		fun()
	}
}

func Register(taskName string, delay, tick time.Duration, fun TimerFunc, param interface{}, funcDefer TimerFunc, paramDefer interface{}) {
	taskManager[taskName] = func() {
		Timer(delay, tick, fun, param, funcDefer, paramDefer)
	}
}
