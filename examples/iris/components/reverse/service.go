package reverse

import (
	"context"
	"github.com/ServiceWeaver/weaver"
)

type Reverser interface {
	Reverse(context.Context, string) (string, error)
}
type reverse struct {
	weaver.Implements[Reverser]
}

func (r *reverse) Reverse(_ context.Context, s string) (string, error) {
	runes := []rune(s)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-i-1] = runes[n-i-1], runes[i]
	}
	return string(runes), nil
}
