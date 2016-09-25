# Semaphores
Semaphores are a very general synchronization mechanism that can be used to implement mutexes, limit access to multiple resources, solve the readers-writers problem, etc.

There is no semaphore implementation in Go's sync package, but they can be emulated easily using buffered channels:

- the capacity of the buffered channel is the number of resources we wish to synchronize
- the length (number of elements currently stored) of the channel is the number of resources current being used
- the capacity minus the length of the channel is the number of free resources (the integer value of traditional semaphores)

