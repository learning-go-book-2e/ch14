# Exercise 3

## Question
3. Assume you have a simple logging function that looks like this:

```go
func Log(ctx context.Context, level Level, message string) {
    var inLevel Level
    // TODO get a logging level out of the context and assign it to inLevel
	if level == Debug && inLevel == Debug {
        fmt.Println(message)
    }
    if level == Info && (inLevel == Debug || inLevel == Info) {
        fmt.Println(message)
    }
}
```

Define a type called `Level` whose underlying type is `string`. Define two constants of this type, `Debug` and `Info`, set to `"debug"` and `"info"`, respectively.

Create functions to store the log level in the context and to extract it.

Create a middleware function to get the logging level from a query parameter called `log_level`. The valid values for `log_level` are `debug` and `info`. 

Finally, fill in the `TODO` in `Log` to properly extract the log level from the context. If the log level is not assigned or is not a valid value, nothing should be printed.

## Solution

We start by defining `Level` and its constants:

```go
type Level string

const (
	Debug Level = "debug"
	Info  Level = "info"
)
```

Next, we need to make our log level context management functions. We first define an unexported type for the key's type, and an unexported constant for the key:

```go
type logKey int

const (
	_ logKey = iota
	key
)
```

After that, we use the two types we've defined to write our context value management functions:

```go
func ContextWithLevel(ctx context.Context, level Level) context.Context {
	return context.WithValue(ctx, key, level)
}

func LevelFromContext(ctx context.Context) (Level, bool) {
	level, ok := ctx.Value(key).(Level)
	return level, ok
}
```

Now we can use the `ContextWithLevel` function and the `Level` type to write our middleware:

```go
func Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		level := r.URL.Query().Get("log_level")
		ctx := ContextWithLevel(r.Context(), Level(level))
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	})
}
```

Finally, we can fill in the missing lines in the `Log` function using `LevelFromContext`:

```go
	inLevel, ok := LevelFromContext(ctx)
	if !ok {
		return
	}
```