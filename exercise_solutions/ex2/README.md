# Exercise 2

## Question
Write a program that adds randomly generated numbers between 0 (inclusive) and 100,000,000 (exclusive) together until one of two things happen: the number 1234 is generated or 2 seconds has passed. Print out the sum, the number of iterations, and the reason for ending (timeout or number reached).

## Solution

This program demonstrates how to respect a timeout in a context.

```go
func main() {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()
	total := 0
	count := 0
	for {
		select {
		case <-ctx.Done():
			fmt.Println("total:", total, "number of iterations:", count, ctx.Err())
			return
		default:
		}
		newNum := rand.Intn(100_000_000)
		if newNum == 1_234 {
			fmt.Println("total:", total, "number of iterations:", count, "got 1,234")
			return
		}
		total += newNum
		count++
	}
}
```

The important part is to use the select-default idiom periodically to check if the timeout has been reached (or if the context has been cancelled for some other reason).

If you are feeling ambitious, you could combine the code here with the middleware from exercise 1 to make a web app.