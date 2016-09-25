package semaphore

type empty struct{}
type Semaphore chan empty

//acquire n resources
func (s Semaphore) P(n int) {
	e := empty{}
	for i := 0; i < n; i++ {
		s <- e
	}

}

// release n resources
func (s Semaphore) V(n int) {
	for i := 0; i < n; i++ {
		<-s
	}
}

/* mutexes */

func (s Semaphore) Lock() {
	s.P(1)
}

func (s Semaphore) UnLock() {
	s.V(1)
}

/* singal-wait */
func (s Semaphore) Signal() {
	s.P(1)
}

func (s Semaphore) Wait(n int) {
	s.V(n)
}
