package topksequences

import (
	"reflect"
	"strings"
	"testing"
)

func Test_getTopKSequences(t *testing.T) {
	type args struct {
		sequenceCountMap SequenceCountMap
		k                int
	}
	tests := []struct {
		name string
		args args
		want []*SequenceCount
	}{
		{
			name: "should return empty result for empty sequence count map",
			args: args{
				sequenceCountMap: SequenceCountMap{},
				k:                100,
			},
			want: []*SequenceCount{},
		},
		{
			name: "should return top k results",
			args: args{
				sequenceCountMap: SequenceCountMap{
					"a": 1,
					"b": 2,
					"c": 3,
				},
				k: 2,
			},
			want: []*SequenceCount{{"c", 3}, {"b", 2}},
		},
		{
			name: "should return all results when k less than map size",
			args: args{
				sequenceCountMap: SequenceCountMap{
					"a": 1,
					"b": 2,
					"c": 3,
				},
				k: 4,
			},
			want: []*SequenceCount{
				{"c", 3},
				{"b", 2},
				{"a", 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getTopKSequences(tt.args.sequenceCountMap, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getTopKSequences() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_processText(t *testing.T) {
	tests := []struct {
		name string
		text string
		want SequenceCountMap
	}{
		{
			name: "should return empty when map when text content is less than the required sequence size",
			text: "abc def",
			want: SequenceCountMap{},
		},
		{
			name: "should return sequence count",
			text: "Hello world example. This is a HELLO WORLD example!!",
			want: SequenceCountMap{
				"hello world example": 2,
				"world example this":  1,
				"example this is":     1,
				"this is a":           1,
				"is a hello":          1,
				"a hello world":       1,
			},
		},
		{
			name: "whitespaces and special characters at start and end",
			text: "  %% how are you ?     !!!!   ",
			want: SequenceCountMap{
				"how are you": 1,
			},
		},
		{
			name: "should support utf-8 characters",
			text: "ᚻᛖ ᚳᚹᚫᚦ ᚦᚫᛏ !!  ",
			want: SequenceCountMap{
				"ᚻᛖ ᚳᚹᚫᚦ ᚦᚫᛏ": 1,
			},
		},
	}
	for _, tt := range tests {
		r := strings.NewReader(tt.text)
		sequenceCountMapStream := make(chan SequenceCountMap)

		t.Run(tt.name, func(t *testing.T) {
			go processText(r, sequenceCountMapStream)

			got := <-sequenceCountMapStream

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getTopKSequences() = %v, want %v", got, tt.want)
			}
		})
	}
}
