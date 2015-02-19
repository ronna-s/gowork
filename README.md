# gowork

#####Set up a bunch of workers to execute your code in paralel while maintaining full control.

###Why?
1. Simple API isnpired by some of the more common frameworks.
1. No need for external tools (the option to configure external tools will be added)

###How to...

##### ... Set up a bunch of workers to run a block of code once
```go
	workerPool := gowork.NewPool(100)
	for w := range workerPool.Workers() {
	  //you may use closures here - if it makes life simpler for you
		w.DoOnce(func() {
		  time.Sleep(time.Duration(rand.Int31n(1000)) * time.Millisecond)
		})
	}
```

##### ... Set up a bunch of workers to run a block of code forever

```go
	workerPool := gowork.NewPool(100)
	for w := range workerPool.Workers() {
	  //you may use closures here - if it makes life simpler for you
		w.Do(func() {
		  time.Sleep(time.Duration(rand.Int31n(1000)) * time.Millisecond)
		})
	}
```


##### ... Sync the workers and run another operation (both `Do` and `DoOnce` support this)

```go
	workerPool := gowork.NewPool(100)
	for w := range workerPool.Workers() {
	  //you may use closures here - if it makes life simpler for you
		w.Do(func() {
		  time.Sleep(time.Duration(rand.Int31n(1000)) * time.Millisecond)
		},
		func(){
		  fmt.Println("All 100 workers are not synced - this code is executed once")
		})
	}		
```

