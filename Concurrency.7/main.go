package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"

	"golang.org/x/net/html"
)

/********
### Goroutines and Channels ###
*********/

/*
Goroutines are lightweight processes that are disposed of by the
by the Go language runtime.
When you start a Go program, the Go language runtime creates
several threads to execute it and starts one goroutine
*/
// nice to watch https://www.youtube.com/watch?v=B9uR2gLM80E and easy to understand
func main() {

	fetchdata() // if we not using concurrency we will wait along time until all func finish
	// fetchdb()
	// fetchcache()
	// fetchapi()

	fmt.Println("done")
	fmt.Println("------")

	// concurrency
	fmt.Println("concurrency")

	// go func() { // anonymous function
	// 	// go func can never be executed
	// 	start := time.Now()
	// 	fetchdata()
	// 	fetchcache()
	// 	fetchapi()
	// 	fetchdb()

	// 	end := time.Since(start)
	// 	fmt.Println("fetchdata took: ", end)
	// }()

	// another way

	// go fetchdata()
	// go fetchdb()

	// go fetchcache()
	// go fetchapi()
	// If the main stops running, the goroutines will be executed
	// so
	// "To ensure goroutines execute concurrently, we must use synchronization mechanisms to coordinate their execution."

	// With channels

	// channel is used to communicate between goroutines

	// you need always read from a channel and write to channel

	/*
		Like maps, channels are a reference type.
		When you pass a channel to a function, it is actually
		passed a pointer to the channel.
		And just as with maps and slices, the nil value for channels is the nil value
	*/

	fmt.Println("Channel")
	fmt.Println("------")

	// channels like the tunnel for goroutines imagine this and this tunnel has 2 ends

	// to initialize a channel we use make
	// chan is a keyword
	// buffered channel
	ch := make(chan int, 2) // make(chan Type|value, length|2)
	// its the buffered channel

	// we can append values to the channel
	// writing example
	ch <- 1 // <- writing to channel (1)
	ch <- 2 // <- writing to channel (2)

	// we can read values from the channel
	// reading example
	fmt.Println(<-ch) // <- reading from channel
	// print 1

	// or...

	result := <-ch
	fmt.Println(result)
	// print 2

	/*
		Each value written to a channel can be read only once. If
		If several goroutines read from one and the same channel,
		only one of these goroutines reads the value written to it.

		If the buffer fills up before any read operations are
		performed from the channel, the next write operation to this channel
		will suspend the writing goroutine.
		The next write operation to that channel will suspend
		the writing goroutine
		until a read from the channel is performed. To the same blocking,
		as in the case of writing to a channel with a full buffer, an attempt
		to read from a channel with an empty buffer


	*/

	// channel in go always block if its full

	// unbeffered channel

	// in unbeffered channels

	ch2 := make(chan string)

	// ch2 <- "fool" // -> is now full and it will block

	// how we can avoid this?

	go func() { // read from channel

		result2 := <-ch2
		fmt.Println(result2)

	}()

	ch2 <- "fool" // write in channel

	/*
		After each operation.
		writing to an open unbuffered channel,
		the writing goroutine pauses until another goroutine reads from this channel.
		pauses until another goroutine reads from this channel.
		And similarly, after each read operation from an open unbuffered channel,
		the read-producing goroutine pauses until the
		another goroutine writes to that channel. This means that in order to read
		or writing to an unbuffered channel requires at least *two goroutines* working in parallel.
	*/

	fmt.Println("------")
	fmt.Println("another example")

	msgch := make(chan string, 8)

	msgch <- "hello"
	msgch <- "world"
	msgch <- "bye"

	close(msgch) // use this to close channel after all values are written

	// msg := <-msgch
	// fmt.Println(msg) // every time when we read from channel we proceed sequentially

	// msg = <-msgch
	// fmt.Println(msg)

	// msg = <-msgch
	// fmt.Println(msg)

	// but we can do this more elegantly

	// This approach is incorrect because len(msgch) always returns the capacity of the channel, not the number of elements in it.

	for m := 0; m < len(msgch); m++ {

		msg := <-msgch
		fmt.Println(msg)
	}

	// or we can use for range

	// loop continues until the channel
	// is closed or the break or return operator is encountered

	// so we need to use this close(channel)

	// This is the preferred method for reading from channels, as it automatically handles channel closure and iteration.
	for msg := range msgch {
		fmt.Println(msg)
	}

	// or we can use this loop

	// This is another correct way to read from a channel after it has been closed.

	for {
		msg, ok := <-msgch
		if !ok { // check if channel is closed
			break
		}
		fmt.Println(msg)
	}

	MakeExampleConcurrency()

	// MakeExampleConcurrency2()

	MakeExample()

	MakeChannel()

	MakeExampleWaitGroup()

	MakeExampleWaitGroup2()

	MakeExampleMutex()
	MakeExampleWithoutMutex()

	MakeExampleAtomic()

	MakeExampleLockFree()

	MakeExampleContext()

	MakeParserContext()

}

func fetchdata() {
	time.Sleep(5 * time.Second) // import time to use Sleep function to wait for 5 seconds

	fmt.Println("fetch data")

}

// what if we have other functions that need to wait for 5 seconds

func fetchdb() {
	fmt.Println("fetch db")
	time.Sleep(5 * time.Second)

}

func fetchcache() {
	fmt.Println("fetch cache")
	time.Sleep(5 * time.Second)

}

func fetchapi() {
	fmt.Println("fetch api")
	time.Sleep(5 * time.Second)

}

// best practice with goroutine
//goroutines are usually triggered by a closure wrapping the business logic

func ProccessValue(value int) (int, error) {

	return 0, nil

}

func RunConcurrency(in <-chan int, out chan<- int) {
	go func() {
		for v := range in {
			result, err := ProccessValue(v)
			if err != nil {
				fmt.Println(err)
				close(out)
				return
			}
			out <- result
		}
		close(out)
	}()
}

//  examples using async

type Server struct {
	ch chan struct{} // for use 0 bytes
	// ch chan bool use 1 byte
	msgch chan string
}

func NewServer() *Server { // allocating memory to heap if we return pointer to struct
	return &Server{
		ch:    make(chan struct{}),
		msgch: make(chan string, 8),
	}
}

func (s *Server) ServeHTTP(w string, r string) {
	fmt.Println(w, "Hello", r)
	fmt.Println("SERVER STARTING")
	go s.ServeLoop()

	/*
		The ServeHTTP method uses go s.ServeLoop() to run ServeLoop in a separate goroutine.
		This allows the main program to continue execution after calling ServeHTTP.

	*/

}
func (s *Server) ServeLoop() {
	/*
		Since the select operator provides data exchange
		with several channels, it is often embedded into a for loop:
	*/
	for {
		select { // select for channels

		/*
						The select statement allows the goroutine to read or write
						to one of several channels and is much like
						the empty switch statement.
			            Of several channels and in many respects resembles
						the empty *switch* operator

		*/
		case <-s.ch: // receive from ch channel ?
			fmt.Println("SERVER STOPPING")
			return // stop server
		//if no receive from  ch channel
		case msg := <-s.msgch:
			s.HandleMessage(msg)

		default:

		}
	}
}

func (s *Server) SendMessage(msg string) {

	s.msgch <- msg
}

func (s *Server) HandleMessage(msg string) {
	fmt.Println(msg)

}

func (s *Server) Stop() {

	s.ch <- struct{}{}
	//or
	//close(s.ch)
}

func MakeExampleConcurrency() {

	fmt.Println("------")

	fmt.Println("MakeExampleConcurrency")

	server := NewServer()

	server.ServeHTTP("GET", "/")

	server.SendMessage("message from server")
	time.Sleep(time.Second) // to give the time for goroutine to process the message before program termination.
	// This helps to make sure that the message is processed before the main program terminates.

}

// another example using goroutine channels and select

type RealServ struct {
	taskCh chan string // taskCh - channel for task transfer

	resultCh chan string // resultCh - the channel for transmitting results.

	closeCh chan struct{} // closeCh - channel for closing the server

	/*

		an empty struct{} structure does not take up memory, this is an efficient way of signaling when
		you don't need to transfer data but just need the fact of completion.

	*/
}

// create constuctor for server

func NewRealServer() *RealServ {
	return &RealServ{
		/*
			Channels are created using the make function,
			which is used to initialize slices, maps, and channels in Go.

		*/
		taskCh:   make(chan string), // we can't use only taskch chan string because is nil
		resultCh: make(chan string),
		closeCh:  make(chan struct{}),
	}
}

func (r *RealServ) Start() { // start our server and handler request
	// start handler in another goroutine
	go r.TaskHandler()

	http.HandleFunc("/", r.HandleRequest)
	http.ListenAndServe(":8080", nil)
}

func (r *RealServ) TaskHandler() { // handler for task
	for { // loop always waiting for tasks or signal for close
		select {
		case task := <-r.taskCh:
			// simulation another process
			time.Sleep(time.Second * 2)
			result := "Processed: " + task

			r.resultCh <- result

		case <-r.closeCh:
			// end task handler
			close(r.resultCh)
			close(r.closeCh)
			close(r.taskCh)
			return

		}
	}
}

func (r *RealServ) HandleRequest(w http.ResponseWriter, s *http.Request) {
	// read task
	task := s.URL.Query().Get("task") // Extract the task parameter from the URL request.
	if task == "" {
		task = "no task"
		http.Error(w, "no task", http.StatusBadRequest)
		return

	}

	r.taskCh <- task // add task to channel

	// Send a task to the taskCh channel so that another goroutine can process it.

	// read result
	select {
	//We use the select statement to expect either a result from the resultCh channel or a timeout after 5 seconds.
	case result := <-r.resultCh:
		fmt.Fprintf(w, "Result : %s\n", result)
	case <-time.After(time.Second * 5):
		http.Error(w, "timeout", http.StatusRequestTimeout)
		// If 5 seconds have passed and no result is received, we return an error with the code 504 Gateway Timeout.
	}

}

func (r *RealServ) Stop() {
	r.closeCh <- struct{}{}
}

func MakeExampleConcurrency2() {
	fmt.Println("------")

	fmt.Println("MakeExampleConcurrency2")
	server := NewRealServer()

	go server.Start()

	fmt.Println("server started at port 8080")

	time.Sleep(time.Second * 5)
	server.Stop()
	fmt.Println("server stopped")
	// to test use http://localhost:8080/task?task=your_task_here
}

/*Imagine that you have several workers (goroutines) who perform tasks
(functions) in parallel. These workers use boxes (channels) to send messages
to each other. Sometimes, workers have multiple tasks and they choose
which task to perform first based on which box is ready (select).

Here is a simple example that combines all three concepts:

*/

func MakeExample() {

	fmt.Println("------")

	fmt.Println("MakeExampleBlocking")
	ch := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}

}

func MakeChannel() { // unbeffered
	ch := make(chan string) // 1

	go func() { // goroutine wait for block another goroutine *block until a value is received*
		message := <-ch      // get ping // 4
		fmt.Println(message) // print ping // 5
		ch <- "pong"         // block when at the time of admission and read and we again go to another goroutine in fmt.Println(<-ch) *block until the value is read* // 6

	}()

	ch <- "ping"      // send ping // 2 // block
	fmt.Println(<-ch) // read pong // 3

	/*
		Why is it done in this sequence?
		Locking and unlocking channels: Channels in Go are synchronous by default, which means that a write operation to a channel is locked until someone starts reading from that channel, and vice versa.

		Sending a "ping" unblocks a goroutine: When the main goroutine sends a "ping" to a channel, this unblocks the goroutine waiting to read from the channel.

		The goroutine sends "pong" and is blocked: After receiving a "ping", the goroutine sends "pong" and is blocked again until the main goroutine reads this value from the channel.

		Main goroutine unblocks and reads "pong": The main goroutine reads "pong" from the channel, ending program execution.

	*/
}

/********
### Waitgroups and Mutex   ###
*********/

/*
Sometimes it is necessary for one goroutine to wait for several other goroutines to complete their work. When you are waiting for a single goroutine to complete, you can
use the pattern described earlier based on the done channel. But when you're waiting
for several goroutines to complete, you need to use the WaitGroup type from the
sync package of the standard library
*/

func MakeExampleWaitGroup() {
	fmt.Println("------")

	fmt.Println("MakeExampleWaitGroup")

	var wg sync.WaitGroup // we need only declared this variable

	// don't use var wg *sync.WaitGroup

	// sync.waithroup has 3 methods

	// add
	// add adds delta, which may be negative, to the WaitGroup counter.
	// If the counter becomes zero, all goroutines blocked on Wait are released.
	// If the counter goes negative, Add panics.

	// done
	// done decrements the WaitGroup counter by one.

	// wait
	// wait blocks until the WaitGroup counter is zero.

	wg.Add(3) // we need add it before goroutine
	// don't add wg.add into goroutine

	go func() {
		defer wg.Done() // its good practice
		//defer wg.Done() ensures that the Done call will be executed even
		// if an error or exception (panic) occurs inside the goroutine.
		fmt.Println("1")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("2")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("3")
	}()

	wg.Wait() // wait for 3 goroutines to finish

}

// but how we can guarantee the fulfillment in a certain order ?

// let's put all of our knowledge

func MakeExampleWaitGroup2() {

	fmt.Println("------") // 1

	fmt.Println("MakeExampleWaitGroup2") // 2

	var wq sync.WaitGroup // 3

	wq.Add(3) // 4
	// create channels for control
	ch1 := make(chan struct{}) // we can use it only for give signal // 5
	ch2 := make(chan struct{}) // its uses 0 memory // 6

	go func() {
		defer wq.Done()
		fmt.Println("1") // 7
		close(ch1)       // signal that 1 goroutine  is done //  8

	}()

	go func() {
		defer wq.Done()
		<-ch1            // wait for 1 goroutine to start block channel // 9
		fmt.Println("2") // 10
		close(ch2)       // signal that 2 goroutine  is done //  11

	}()

	go func() {
		defer wq.Done()
		<-ch2            // wait for 2 goroutine to start block channel  // 12
		fmt.Println("3") // 13

	}()

	wq.Wait() // wait for 3 goroutines to finish

}

// scenario for example

/*
The second goroutine is terminated by the first one.

Goroutine 2 starts executing, but it is blocked at <-ch1 (line 9) because ch1 is not yet closed.
Goroutine 3 also remains blocked at <-ch2 (line 12).
Goroutine 1 is executed and prints "1" (line 7).
It then closes the ch1 channel (line 8).
Goroutine 2 unlocks, receives the ch1 close signal, and prints "2" (line 10).
 Goroutine 2 then closes the ch2 channel (line 11).
 Goroutine 3 unlocks, receives the ch2 close signal and prints "3" (line 13).


*/

// another way control goroutine using mutex

type Count struct {
	counter int
	mu      sync.Mutex
}

func (c *Count) Inc() {
	c.mu.Lock() //   Capture the mutex before accessing the counter
	// Capture mutex to block access to counter other goroutines.
	defer c.mu.Unlock() // // Release the mutex when the function is terminated
	c.counter++
	fmt.Println(c.counter)
}

// MakeExampleMutex creates an instance of Count and starts 10 goroutines
// each of which calls the Inc method to increment the counter by 100000 units
func MakeExampleMutex() {

	count := &Count{}

	fmt.Println("------")

	fmt.Println("MakeExampleMutex")
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(c *Count) { // we need to avoid problem with
			// capturing the variables from the closures
			defer wg.Done()
			for j := 0; j < 10; j++ {
				c.Inc()
			}
		}(count)
	}

	wg.Wait()
}

/*
------

Data race - When multiple goroutines simultaneously try to change the same variable without synchronization, it results in a data race condition
Incorrect counter value - Due to data races, the final counter value is likely to be incorrect because the increment operations will not be performed atomically.
Potential crashes and errors

*/

// lets check

type Count2 struct {
	counter int
}

func (c *Count2) Inc2() {
	c.counter++
	fmt.Println(c.counter)
}

func MakeExampleWithoutMutex() {

	count := &Count2{}

	fmt.Println("------")

	fmt.Println("MakeExampleWithoutMutex")

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(c *Count2) {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				c.Inc2()
			}
		}(count)
	}

	wg.Wait() // you can check what will be if we not using wg.wait()

	// in result of this function you will see the data race
}

// to check data race use *go run -race main.go*

// go has alternative for mutex

/********
### Atomic value  ###
*********/

// you can never see it in go but its important to know

// see why we need it in tips.md

// hard example but in this case atomic is optional

/*
	Atomic operations: Although the atomic.AddInt64 and atomic.LoadInt64 operations
	are atomic, they do not prevent competition for access to the count variable. Each goroutine can modify and read the count value,
	which can lead to undefined ordering of operations and output.

	If you need to squeeze out all possible performance and you are an
   expert in writing competitive code, you will be pleased with Go's support for atomic operations


*/

func MakeExampleAtomic() {
	fmt.Println("------")

	fmt.Println("MakeExampleAtomic")

	var wg sync.WaitGroup

	var count int64

	wg.Add(1)
	go func(c int64) {
		defer wg.Done()
		for j := 0; j < 100; j++ {
			// AddInt64 atomically adds delta to *addr and returns the new value.
			atomic.AddInt64(&c, 1) // useful for increment or decrement
			// func AddInt64(addr *int64, delta int64) (new int64)

			// LoadInt64 atomically loads *addr.
			fmt.Println(atomic.LoadInt64(&c)) // Loads the value pointed to by *addr
			// func LoadInt64(addr *int64) (val int64)

		}
	}(count)

	wg.Wait()

}

// another example

type Node struct {
	value int   // value
	next  *Node // pointer to next node
}

type Stack struct {
	top unsafe.Pointer // top point to top element of stack. Use unsafe.Pointer for atomic operations on pointers
}

// add element to stack
func (s *Stack) Push(value int) {
	newNode := &Node{value: value}
	for { // try add new node to top stack
		// LoadPointer atomically loads *addr.
		// func LoadPointer(addr *unsafe.Pointer) (val unsafe.Pointer)

		// load current top
		oldTop := (*Node)(atomic.LoadPointer(&s.top)) // use LoadPointer to read current top

		newNode.next = oldTop

		// replace top with new node
		if atomic.CompareAndSwapPointer(&s.top, unsafe.Pointer(oldTop), unsafe.Pointer(newNode)) {
			break
		}
	}
}

// delete and return element from stack
func (s *Stack) Pop() (int, bool) {
	for {
		// load current top
		oldTop := (*Node)(atomic.LoadPointer(&s.top))

		if oldTop == nil {
			return 0, false // stack is empty
		}

		// replace top with next
		if atomic.CompareAndSwapPointer(&s.top, unsafe.Pointer(oldTop), unsafe.Pointer(oldTop.next)) {
			return oldTop.value, true
		}
	}
}

func MakeExampleLockFree() {
	fmt.Println("------")

	fmt.Println("MakeExampleLockFree")

	stack := &Stack{}

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				stack.Push(i*10 + j)
				fmt.Println(stack.Pop())
			}
		}(i)
	}

	wg.Wait()

	// get elements from stack

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				value, ok := stack.Pop()

				if !ok {
					break
				}
				fmt.Println(value)
			}
		}()
	}

	wg.Wait()

}

/*

atomic.LoadPointer and atomic.CompareAndSwapPointer
are used to safely read and modify pointers without locks.


if you interesting to know more about it

https://pkg.go.dev/sync/atomic
https://pkg.go.dev/unsafe
https://github.com/valyala




*/

/********
### Context ###
*********/

// its used a lot in go
// it's important very important to know

// context always declared in functions as first parameter

// why do wee need context
// imitation of api call
func httpCallToApi(ctx context.Context, s string) (string, error) {

	ch := make(chan string) // add a channel to communicate with the goroutine

	go func(s string) { // create a goroutine to call the long api

		time.Sleep(50 * time.Second) // imitation long api call to get data

		ch <- s
	}(s)

	go func(s string) {

		time.Sleep(50 * time.Millisecond) // imitation long api call to get data
		ch <- s + "goroutine"
	}(s)

	// We use select to wait for either the context to complete or for the result to be retrieved from the channel.

	select {
	case <-ctx.Done(): // # Check if the context has finished
		// context has been canceled or timed out
		return "", ctx.Err() // #  Return an error if the context has ended
	case res := <-ch: // # Wait for the result from the goroutine
		return res, nil // # return the result
	}
}

func MakeExampleContext() {

	fmt.Println("------")

	fmt.Println("MakeExampleContext")
	start := time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)

	defer cancel() // Provide a cancel call to release resources
	result, err := httpCallToApi(ctx, "test-ctx")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("the response took %v:   %+v\n", time.Since(start), result)
}

/*

Contexts in Go are used to manage
runtimes, cancel tasks, and pass values between goroutines


The simulated API call will now fail
because the time set in the context (50 milliseconds) has expired,
since the API call itself takes 5 seconds. This demonstrates
how context is used to control the timing of tasks and cancel them
if necessary.


*/

/********

The main functions and methods of the context package:
context.Background() and context.TODO():

context.Background(): Creates an empty context. Typically used in the main goroutine and to start a call chain.
context.TODO(): Used as a temporary context when the context is not yet defined.
context.WithCancel(parent Context) (ctx Context, cancel CancelFunc):

Creates a new context that can be canceled manually by calling cancel().
context.WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc):

Creates a new context with a set timeout time. The context will be automatically canceled after the specified timeout time.
context.WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc):

Creates a new context with a specified deadline time.
context.WithValue(parent Context, key, val interface{}) Context:

Creates a new context with an added key-value pair. This is useful for passing additional data through the context.



******/

/// another example

func exctractText(ctx context.Context, url string) (string, error) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to parse HTML: %w", err)
	}

	var extractText func(*html.Node) string
	extractText = func(n *html.Node) string {
		if n.Type == html.TextNode {
			return n.Data
		}

		text := ""

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			text += extractText(c) + " "
		}
		return text

	}
	return extractText(doc), nil
}

func MakeParserContext() {

	fmt.Println("------")

	fmt.Println("MakeParserContext")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	defer cancel()

	url := "https://easyoffer.ru/rating/golang_developer"

	result, err := exctractText(ctx, url)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("the response took ", result)

}
