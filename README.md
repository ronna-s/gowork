# gowork

```bash
go get github.com/ronna-s/gowork
```

#####Set up a bunch of workers to execute your code in paralel while maintaining full control.

###Why?
1. Simple API isnpired by some of the more common frameworks.
1. No need for external tools (the option to configure external tools will be added)

###How to...

#### ... Run a job every x seconds:
```go
gowork.IterateEvery(1).Run(func() {
	fmt.Println("Hello world")
})
```

#### ... Setup a bunch of workers to run a block of code in parallel (RunInParallel or RunInParallelWithIndex)
```go
package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/ronna-s/gowork"
)

func main() {
	gowork.NewPool(10).RunInParallelWithIndex(func(i int) {
		rand.Seed(time.Now().UTC().UnixNano())
		time.Sleep(time.Duration(rand.Int31n(10)) * time.Millisecond)
		fmt.Println("Hello from", i)
	})
}

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
```
#### ... and if you combine the two, you schedule a job with x workers to run every x seconds:
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
#### ... Set up a bunch of workers to run a block of code forever

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
#### ... Sync the workers to run another operation

```go
func main() {
	pool := gowork.NewPool(2)
	pool.GetWorker().Do(func() {
		fmt.Println("Hello hello")
	})
	pool.GetWorker().Do(func() {
		fmt.Println("Hello world")
	})
	pool.Sync()
}		
```
Produces:
```
Hello world
Hello hello
```
