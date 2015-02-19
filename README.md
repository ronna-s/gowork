# gowork

##Set up a bunch of workers to execute your code in paralel while maintaining full control.

##Why?
1. Simple API isnpired by some of the more common frameworks.
1. No need for external tools (the option to configure external tools will be added)

##How...

### ... To Set up a bunch of workers to run a block of code once
```go
	workerPool := gowork.NewPool(100)
	for w := range workerPool.Workers() {
	  //you may use closures here - if it makes life simpler for you
		w.DoOnce(func() {
		  time.Sleep(time.Duration(rand.Int31n(1000)) * time.Millisecond)
		})
```

### ... To Set up a bunch of workers to run a block of code forever

```go
	workerPool := gowork.NewPool(100)
	for w := range workerPool.Workers() {
	  //you may use closures here - if it makes life simpler for you
		w.Do(func() {
		  time.Sleep(time.Duration(rand.Int31n(1000)) * time.Millisecond)
		})
```


### ... To sync the workers and run another operation (both Do and DoOnce support this)

```go
	workerPool := gowork.NewPool(100)
	for w := range workerPool.Workers() {
	  //you may use closures here - if it makes life simpler for you
		w.Do(func() {
		  time.Sleep(time.Duration(rand.Int31n(1000)) * time.Millisecond)
		},func(){
		})
```

