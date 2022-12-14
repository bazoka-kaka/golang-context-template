package main

import (
	"context"
	"fmt"
)

func populateContext(ctx context.Context, name string, value interface{}) context.Context {
	return context.WithValue(ctx, name, value)
}

func printContext(ctx context.Context, name string) {
	fmt.Printf("%s: %+v\n", name, ctx.Value(name))
}

func main() {
	ctx := context.Background()
	names := []string{
		"username", "password", "age", "online",
	}
	ctx = populateContext(ctx, "username", "benzion")
	ctx = populateContext(ctx, "password", "yehezkel")
	ctx = populateContext(ctx, "age", 21)
	ctx = populateContext(ctx, "online", false)
	for _, name := range names {
		printContext(ctx, name)
	}
}