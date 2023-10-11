<img src="https://avatars2.githubusercontent.com/u/24763891?s=400&u=c1150e7da5667f47159d433d8e49dad99a364f5f&v=4"  width="256px" height="256px" align="right" alt="Multiverse OS Logo">

## Multiverse: `signal` Handling Library
**URL** [multiverse-os.org](https://multiverse-os.org)

A signal handling library built ontop of the Go standard library `os/signal`. 

Designed for broad usage beyond simple shutdown signal handling, while still
including extra functionality to streamline and simplify usage for shutdown
signal handling. Beyond signal handling and running one function or more in
response to a OS signal, the library is capable of ignoring signals using the
newer functionality added to the underlying library.


#### Usage 
There is a variety of ways to interact with the library, and a different ways to
initialize the functions to run in response to OS signals:

```go
package main

import (
  "fmt"

  signal "github.com/multiverse-os/signal"
)

func main() {
  // Shutdown is a group of signals that include: SIGINT, SIGTERM, SIGQUIT,
  // SIGHUP, and SIGKILL
  optionOne := signals.Handler()
  optionOne.OnShutdown(func(s os.Signal){
    fmt.Println("received a signal:", s)
  })

  // Additional hooks can be added with chainable hook initializing functions
  // Calling a function to add a function in response to a signal will append
  // it to the existing functions 
  optionOne.OnTerminate(func(s os.Signal){
    fmt.Println("terminated")
  }).OnHangup(func(s os.Signal){
    fmt.Println("terminal hung up")
  })

  // And we can reset what we have done with, stopping signals comming down the
  // channels and removing the hooks from our map (see the code for more fine
  // grain clear/removal and cancelation)
  optionOne.Reset()

  // And we can ignore signals using
  optionOne.Ignore(signal.Teminate, signal.SIGINT)

  // Or very simple one line interaction is possible
  optionTwo := signals.ShutdownHandler(func(s os.Signal){
    fmt.Println("received a singal:", s)
  })
}
```

These examples above do not yet document all the possibilities and functionality
contained in the ~150 line of code (LOC) file. And until they are documented the
source code is written to be human-readable as possible without being noisy. 

#### Contributing
Volunteers are wanted to help improve the quality and competeness Any feature
requests, push requets, documentation fixes are welcome, please just create a
pull request and a code review will be initiated. 
