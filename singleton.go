package main

import "sync"

type Single struct{}

var (
	singleInstance *Single
	singleMutex    sync.Mutex
)

func GetSingleInstance() *Single {
	if singleInstance == nil {
		singleMutex.Lock()
		defer singleMutex.Unlock()

		// double check
		if singleInstance == nil {
			singleInstance = new(Single)
		}
	}

	return singleInstance
}

var singleOnce sync.Once

func GetSingleInstance2() *Single {
	singleOnce.Do(func() {
		singleInstance = new(Single)
	})

	return singleInstance
}
