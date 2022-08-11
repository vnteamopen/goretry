
# Go Retry

[![Go Reference](https://pkg.go.dev/badge/github.com/vnteamopen/goretry.svg)](https://pkg.go.dev/github.com/vnteamopen/goretry) [![build_pr](https://github.com/vnteamopen/goretry/actions/workflows/build.yml/badge.svg)](https://github.com/vnteamopen/goretry/actions/workflows/build.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/vnteamopen/goretry)](https://goreportcard.com/report/github.com/vnteamopen/goretry) 
[![Built with WeBuild](https://raw.githubusercontent.com/webuild-community/badge/master/svg/WeBuild.svg)](https://webuild.community) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/vnteamopen/goretry/blob/main/LICENSE)

Go Retry is Go language library that enables the code can retry to do it if it's failed.

It's good to handle transient failures when it tries to connect to a service or network resource, or try to do a required logic.

Go Retry library supports multiple strategies:
 - [Constant Backoff](#constant-backoff)
 - [Linear Backoff](#linear-backoff)
 - [Fibonacci Backoff](#fibonacci-backoff)
 - [Quadratic Backoff](#quadratic-backoff)
 - [Exponential Backoff](#exponential-backoff)
 - [Polynomial Backoff](#polynomial-backoff)
 - And [Jitter](#jitter)

## Quick starts

1. Download dependencies:
```
go get -u github.com/vnteamopen/goretry
```

2. Import and usage:

```go
package main

import (
	"fmt"
	"github.com/vnteamopen/goretry"
)

func main() {
	counter := 0
	goretry.Do(10*time.Millisecond, func() error {
		if counter == 5 {
			fmt.Println("Success")
			return nil
		}
		fmt.Println("Fail")
		counter++
		return fmt.Errorf("fake error")
	})
}
```
## Features

### Custom instance

All functions from goretry such as `goretry.Do()`, `goretry.Fibonacci()` uses default configuration instance. You can custom configuration of goretry with changes:

```go
package main

import (
	"fmt"
	"github.com/vnteamopen/goretry"
)

func main() {
	customRetry := Instance{
		MaxStopRetries: 10, // maximum number of retry times. Default: NoLimit.
		MaxStopTotalWaiting: time.Duration(5 * time.Minute), // maximum total waiting duration of retry times. Default: NoLimit
		CeilingSleep: time.Duration(1 * time.Minute), // maximum duration waiting that can increase to.
		Logger: os.Stdout, // Logger defines log output, main purpose for debug
	}
	customRetry.Do(10*time.Millisecond, func() error {
		// Do something here
	})
}
```

### Constant Backoff

Simplest case of retry strategy, it keep constant waiting duration beetween retry actions.

Constant backoff retry adds a fixed waiting duration after the first failure and between retry actions with the following formula:

```
duration(t) = <constant value>
```

![constant backoff](https://thuc.space/images/retry_strategies/retry_strategies-constant_backoff.png)

```go
package main

import (
	"github.com/vnteamopen/goretry"
)

func main() {
	goretry.Do(5*time.Second, func() error {
		// Do something here
	})
}
```

### Linear Backoff

The linear backoff retry strategy supports the waiting duration after the first failure and increases the waiting duration of the next retries. It increases constantly delta x from the previous waiting time with the formula:

```
duration(t) = duration(t-1) + x
```

![Linear backoff](https://thuc.space/images/retry_strategies/retry_strategies-linear_backoff.png)

```go
package main

import (
	"github.com/vnteamopen/goretry"
)

func main() {
	firstWaiting := 1*time.Second
	increasement := 1*time.Second
	goretry.Linear(firstWaiting, increasement, func() error {
		// Do something here
	})
}
```

### Fibonacci Backoff

The Fibonacci backoff retry strategy supports the waiting duration that calculates by using the Fibonacci sequence with the formula:

```
duration(t) = duration(t-1) + duration(t-2)
```

![Fibonacci backoff](https://thuc.space/images/retry_strategies/retry_strategies-fibonacci_backoff.png)

```go
package main

import (
	"github.com/vnteamopen/goretry"
)

func main() {
	initWaiting := 1*time.Second
	goretry.Fibonacci(initWaiting, func() error {
		// Do something here
	})
}
```

### Quadratic Backoff

The quadratic backoff retry strategy supports the waiting duration that calculates by following the quadratic curve with the formula:

```
duration(t) = attempt ^ 2 * base-time
```

![Quadratic Backoff](https://thuc.space/images/retry_strategies/retry_strategies-quadratic_backoff.png)

```go
package main

import (
	"github.com/vnteamopen/goretry"
)

func main() {
	baseDuration := time.Second
	goretry.Quadratic(baseDuration, func() error {
		// Do something here
	})
}
```

### Exponential Backoff

The exponential backoff retry strategy supports the waiting duration that calculates by following the exponential curve with the formula:

```
duration(t) = 2 ^ attempt * base-time
```

![Exponential Backoff](https://thuc.space/images/retry_strategies/retry_strategies-exponential_backoff.png)

```go
package main

import (
	"github.com/vnteamopen/goretry"
)

func main() {
	baseDuration := time.Second
	goretry.Exponential(baseDuration, func() error {
		// Do something here
	})
}
```

### Polynomial Backoff

The Polynomial backoff retry strategy supports the waiting duration that calculates by following the formula:

```
duration(t) = attempt ^ degree * base-time
```

![Polynomial Backoff](https://thuc.space/images/retry_strategies/retry_strategies-polynomial_backoff.png)

```go
package main

import (
	"github.com/vnteamopen/goretry"
)

func main() {
	baseDuration := time.Second
	degree := 3
	goretry.Polynomial(baseDuration, degree, func() error {
		// Do something here
	})
}
```

### Jitter

Suppose we have multiple retry callers do an action or send requests that collide and fail. They all decide to retry with a backoff strategy. They will all retry at the same time, which leads to colliding again.

Jitter is a technique to solve that problem. It adds or removes different random waiting durations to back off time. So each next retry will happen at a different time and help to avoid several calls next time.

![Jitter](https://thuc.space/images/retry_strategies/retry_strategies-jitter.png)

```go


package main

import (
	"github.com/vnteamopen/goretry"
)

func main() {
	customRetry := Instance{
		JitterEnabled: true,                     // Enable Jitter
		JitterFloorSleep: 10 * time.Millisecond, // Minimun waiting duration after jitter is 10ms
	}
	customRetry.Do(...)
	customRetry.Fibonacci(...)
}
```

# Links
 - https://vnteamopen.com/
 - Backoff strategy https://thuc.space/posts/retry_strategies/
