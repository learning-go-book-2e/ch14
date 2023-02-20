package main

import (
	"context"
	"fmt"
	"math/big"
	"time"
)

func main() {
	calcWithTimeout(1_000_000)
	calcWithTimeout(10_000_000)
	calcWithTimeout(10_000_000_000)
	calcWithoutTimeout(20_000_000)
}

func calcWithoutTimeout(precision int) {
	ctx := context.Background()
	result, err := calcPi(ctx, precision)
	fmt.Println(result)
	fmt.Println(err)
}

func calcWithTimeout(precision int) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()
	result, err := calcPi(ctx, precision)
	fmt.Println(result)
	fmt.Println(err)
}

func calcPi(ctx context.Context, precision int) (string, error) {
	var sum big.Float
	sum.SetInt64(0)
	var d big.Float
	d.SetInt64(1)
	two := big.NewFloat(2)
	for i := 0; i < precision; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("cancelled after", i, "iterations")
			return sum.Text('g', 100), context.Cause(ctx)
		default:
		}
		var diff big.Float
		diff.SetInt64(4)
		diff.Quo(&diff, &d)
		if i%2 == 0 {
			sum.Add(&sum, &diff)
		} else {
			sum.Sub(&sum, &diff)
		}
		d.Add(&d, two)
	}
	fmt.Println("completed", precision, "iterations")
	return sum.Text('g', 100), nil
}
