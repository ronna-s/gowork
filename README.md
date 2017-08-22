# gowork

```bash
go get github.com/ronna-s/gowork
```

#####Set up a bunch of workers to execute your code in paralel while maintaining full control.

###Why?
1. Simple API isnpired by some of the more common frameworks.
1. No need for external tools (the option to configure external tools will be added)

###How...

###### ... To run a job every x seconds use IterateEvery
```go
gowork.IterateEvery(1).Run(func() {
	fmt.Println("Hello world")
})
```

###### ... To setup a bunch of workers to run a block of code in parallel use RunInParallel or RunInParallelWithIndex
```go
gowork.NewPool(10).RunInParallelWithIndex(func(i int) {
	rand.Seed(time.Now().UTC().UnixNano())
	time.Sleep(time.Duration(rand.Int31n(10)) * time.Millisecond)
	fmt.Println("Hello from", i)
})
fmt.Println("Hello I'm done")

```
Can produce:
```
Hello from 9
Hello from 7
Hello from 0
Hello from 5
Hello from 4
Hello from 8
Hello from 3
Hello from 6
Hello from 1
Hello from 2
Hello I'm done
```
RunInParlalle and RunInParallelWithIndex will wait (Sync) until all workers have finished.
###### ... and if you combine the two, you schedule a job with x workers to run every x seconds:
If the workers are not done in time, the next iteration will start immediately after the last worker is done.
```go
p := gowork.NewPool(10)
gowork.IterateEvery(1).Run(func() {
	p.RunInParallelWithIndex(func(i int) {
		rand.Seed(time.Now().UTC().UnixNano())
		time.Sleep(time.Duration(rand.Int31n(10)) * time.Millisecond)
		fmt.Println("Hello from", i)
	})
	fmt.Println("==========")
})
```
###### ... to set up a bunch of workers to run a block of code forever range over Workers.

```go
workerPool := gowork.NewPool(100)
for w := range workerPool.Workers() {
	//you may use closures here - if it makes life simpler for you
	w.Do(func() {
		rand.Seed(time.Now().UTC().UnixNano())
		time.Sleep(time.Duration(rand.Int31n(1000)) * time.Millisecond)
	})
}
```
###### ... get rid of WaitGroups - simply use Sync()

```go
func main() {
	p := gowork.NewPool(2)

	p.GetWorker().Do(func() {
		fmt.Println("Hello from 1")
	})

	p.GetWorker().Do(func() {
		fmt.Println("Hello from 2")
	})
	//wait - what? but we only have 2 worker
	p.GetWorker().Do(func() {
		fmt.Println("Hello from 3")
	})

	p.GetWorker().Do(func() {
		fmt.Println("Hello from 4")
	})

	fmt.Println("hello from foo")

	p.Sync()
	fmt.Println("all workers are done")
}
```
Produces:
```
Hello world
Hello hello
```
###TODO:
* support priorities
* support schedules
