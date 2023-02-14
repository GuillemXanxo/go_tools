##Time Advice

#### Prefer `30 * time.Second` instead of `time.Duration(30) * time.Second`

You don't need to wrap untyped const in a type, compiler will figure it out. Also prefer to move const to the first place:
```go
// BAD
delay := time.Second * 60 * 24 * 60

// VERY BAD
delay := 60 * time.Second * 60 * 24

// GOOD
delay := 24 * 60 * 60 * time.Second

// EVEN BETTER
delay := 24 * time.Hour
```

#### Use `time.Duration` instead of `int64` + variable name

```go
// BAD
var delayMillis int64 = 15000

// GOOD
var delay time.Duration = 15 * time.Second
```