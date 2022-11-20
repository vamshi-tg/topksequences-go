package topksequences

import (
	"github.com/gammazero/deque"
	"testing"
)

func Test_sequenceCountMapKey(t *testing.T) {
	type args struct {
		deque *deque.Deque[string]
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "should return key",
			args: args{
				deque: func() *deque.Deque[string] {
					var q deque.Deque[string]
					q.PushBack("foo")
					q.PushBack("bar")
					q.PushBack("baz")
					return &q
				}(),
			},
			want: "foo bar baz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sequenceCountMapKey(tt.args.deque); got != tt.want {
				t.Errorf("sequenceCountMapKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
