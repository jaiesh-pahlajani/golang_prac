# Golang study notes


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
- Work stealing helps in better distribution of goroutines across all logical processors. 

###### Channels
- Communicate data between goroutines
- Synchronise goroutines
- typed
- thread safe
- The goroutine that creates, writes and closes the channel is ideally the owner of that channel. Goroutine that utilizes the channel only reads from the channel.

###### Ownership of channels avoid
- Deadlocking by writing to a nil channel
- closing a nil channel
- writing to a closed channel
- closing a channel more than once

###### Hchan structure
- hchan represents channel
- Its contains circular ring buffer and mutex lock
- There is no memory shared between goroutines
- Goroutines copy elements into and from hchan
- hchan is protected by mutex lock
- 'Do not communicate by sharing memory, instead share memory by communicating'

![Alt text](attachments/hchan.png?raw=true "")


###### Send and receive cases for buffered channels
- When a channel is full and goroutine tries to send value
    - Sender goroutine gets blocked, its parked on sendq
    - Data will be saved in elem field of sudog structure
    - Whenever receiver comes along it deques from value from buffer
    - Enqueues data from ele field to buffer
    - Pops the goroutine in sendq and puts it in runnable state. 
- When goroutine calls on empty buffer
    - goroutine is blocked, its parked into recq
    - elem field of the sudog structure holds reference to the stack variable of receiving goroutine
    - When sender comes along, sender finds goroutine in receiver queue. 
    - Sender copies the data directly onto the stack variable of the receiver goroutine
    - pops the goroutine in recq and puts in runnable state

![Alt text](attachments/senq_and_recq.png?raw=true "")

###### Send and receive cases for unbuffered channels
- Sender goroutine wants to send values on channel
    - If there is a corresponding reiver waiting in recvq, sender will write value directly into receiver goroutine stack variable.
    - The sender goroutine will put receiver goroutine back in runnable state.
    - If there is no receiver goroutine in recvq, sender gets parked into sendq
    - Data is saved in elem field in sudog struct
    - Receiver comes and copies the data, puts the sender into runnable state again. 
- Receiver goroutine wants to receive value
    - If its finds sender goroutine in sendq, receiver copies the value in elem field to its stack variable
    - Puts the sender goroutine in runnable state.
    - If there was no sender goroutine in sendq, then receiver gets parked into recvq.
    - Reference to variable is saved in elem field in sudog struct.
    - Sender comes along it copies data directly into the receiver stack variable. Puts the variable back into runnable state.
    
###### Select 
- Select is like a switch statement with each case statement specifying channel operation
- Select will block until any of the case statement is ready
- With select we can implement timeout and non blocking communication
- Select on nil channel will block forever

###### Mutex
- Used for protecting shared resources
- Caches and states
- Critical section represents the bottleneck between the goroutine.

###### Atomic 
- Low level operations on memory
- Lockless operations
- Useful for counters

###### Conditional variable
- Conditional variable is used to synchronise execution of goroutines
- Wait suspends the execution of goroutine
- Signal wakes one goroutine waiting on c
- Broadcast wakes all goroutines waiting on c