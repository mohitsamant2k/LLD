package taskScheduler

import (
	"container/heap"
	"fmt"
	"sync"
	"time"
)

type TaskScheduler struct {
	numOfThread int // max number of allowed worker thread
	taskQueue   *PriorityQueue
	mu          sync.Mutex
	signal      chan struct{}
	stop        chan struct{}
	timerChan   <-chan time.Time // receive only channel
}

func NewTaskScheduler(numThread int) *TaskScheduler {
	pq := &PriorityQueue{}
	heap.Init(pq)

	ret := &TaskScheduler{
		numOfThread: numThread,
		taskQueue:   pq,
		signal:      make(chan struct{}),
		stop:        make(chan struct{}),
	}

	for range numThread {
		go ret.Worker()
	}

	return ret
}

func (t *TaskScheduler) Worker() {
	// wait for timer or new task signal
	for {
		select {
		case <-t.stop:
			return
		case <-t.signal:
			t.mu.Lock()
			// fmt.Println("worker thread")
			if t.taskQueue.Len() == 0 {
				continue
			}
			task := heap.Pop(t.taskQueue).(*Task)
			heap.Push(t.taskQueue, task)
			now := time.Now()
			timer := time.NewTimer(task.runat.Sub(now))
			// fmt.Printf("%+v, %+v", now, task.runat)
			t.timerChan = timer.C
			t.mu.Unlock()
		case <-t.timerChan:
			//
			
			t.mu.Lock()
			if t.taskQueue.Len() == 0 {
				// fmt.Printf("continue")
				continue
			}
			task := heap.Pop(t.taskQueue).(*Task)
			t.signal <- struct{}{}
			// fmt.Println(" task is there")
			
			t.mu.Unlock()
			task.task()
		}
	}

}

type Task struct {
	task  func()
	runat time.Time
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Task

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].runat.Before(pq[j].runat)
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.(*Task))
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func (t *TaskScheduler) AddTask(task func(), delay int) {
	runat := time.Now().Add(time.Duration(delay) * time.Second)
	curr := &Task{
		task:  task,
		runat: runat, // convert the delay into duration
	}

	// When a new task comes, we will send a signal to the worker thread that a new task is available.
	// One thread can continue and check the highest priority task.
	// If that task's runat is before or equal to the current time, then run that task.
	// If not, then go to sleep for some time and wait for the timer signal.
	t.mu.Lock()
	heap.Push(t.taskQueue, curr)
	t.signal <- struct{}{}
	t.mu.Unlock()
}

// RunDemo demonstrates scheduling 3 tasks with different delays and waits for them to finish.
func RunDemo() {
	fmt.Println("Task Scheduler Demo start")
	sched := NewTaskScheduler(2)
	var wg sync.WaitGroup

	start := time.Now()
	add := func(name string, delay int) {
		wg.Add(1)
		sched.AddTask(func() {
			fmt.Printf("%s executed at %s (delay %ds)\n", name, time.Now().Format(time.StampMilli), delay)
			wg.Done()
		}, delay)
		fmt.Printf("Scheduled %s (runs in %ds)\n", name, delay)
	}
	fmt.Println("here is addong")
	add("love", 4)
	add("you", 5)
	add("--", 6)
	time.Sleep(2 * time.Second)
	add("I", 1)

	wg.Wait()
	fmt.Printf("All tasks done in %v\n", time.Since(start))
}
