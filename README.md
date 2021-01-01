# golang_prac


###### M:N Scheduler
- The go scheduler is aprt of the go runtime. It is known as M:N scheduler.
- Go scheduler runs in user space.
- GO scheduler uses OS threads to schedule go routines for execution
- Goroutines runs in the context of os threads.
- Go runtime creates number of worker OS threads, equal to `GOMAXPROCS`
- `GOMAXPROCS` - default value is number of processorts on machine
- Go scheduler distributes runnable goroutines over multiple worker OS threads. (4 cores ~ 4 OS threads)
- At anytime N goroutines can be scheduled on M OS threads that runs on at most `GOMAXPROCS` number of processors


###### Async Preemption
- GO scheduler implements async preemption
- This prevents long running goroutines from hoggin onto CPU, that could block other goroutines.
- The async preemption is tirggered based on a time condition. When a goroutine is running from more than 10 ms, go will try to preeempt it. 


###### Goroutine states

![Alt text](attachments/goroutine_states.png?raw=true "")


###### Go scheduler

![Alt text](attachments/go_scheduler.jpg?raw=true "")


###### Context switching

- Synchronous System call 
    - When goroutine makes synchronous call system call, Go scheduler brings a new OS thread from thread pool.
    - Moves the logical processor P to new thread. 
    - Goroutine which made the system call will still be attached to old thread.
    - Other goroutines in LRQ are scheduled for execution on new OS thread.
    - Once system call returns, goroutine is moved back to run queue on logical processor P and old thread is put to sleep.
-  Asynchronous System call
    - Go uses netpoller to handle async system call.
    - netpoller uses interface provided by OS to do polling on file descriptors and notifies the goroutine to try I/O operation when its ready.
    - Application complexity of managing async system call is moved to Go runtime, which manages it efficiently

###### Work Stealing
- If logical processor run out of go routines in its local run queue, it will steal go routines from other logical processors or global run queue.
