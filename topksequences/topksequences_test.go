package topksequences

import (
	"reflect"
	"strings"
	"testing"
)

func Test_getTopKSequences(t *testing.T) {
	type args struct {
		sequenceCountMap sequenceCountMap
		k                int
	}
	tests := []struct {
		name string
		args args
		want []*sequenceCount
	}{
		{
			name: "should return empty result for empty sequence count map",
			args: args{
				sequenceCountMap: sequenceCountMap{},
				k:                100,
			},
			want: []*sequenceCount{},
		},
		{
			name: "should return top k results",
			args: args{
				sequenceCountMap: sequenceCountMap{
					"a": 1,
					"b": 2,
					"c": 3,
				},
				k: 2,
			},
			want: []*sequenceCount{{"c", 3}, {"b", 2}},
		},
		{
			name: "should return all results when k less than map size",
			args: args{
				sequenceCountMap: sequenceCountMap{
					"a": 1,
					"b": 2,
					"c": 3,
				},
				k: 4,
			},
			want: []*sequenceCount{
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
		want sequenceCountMap
	}{
		{
			name: "should return empty when map when text content is less than the required sequence size",
			text: "abc def",
			want: sequenceCountMap{},
		},
		{
			name: "should return sequence count",
			text: "Hello world example. This is a HELLO WORLD example time!!",
			want: sequenceCountMap{
				"hello world example": 2,
				"world example this":  1,
				"example this is":     1,
				"this is a":           1,
				"is a hello":          1,
				"a hello world":       1,
				"world example time":  1,
			},
		},
		{
			name: "whitespaces and special characters at start and end",
			text: "  %% how are you ?     !!!!   ",
			want: sequenceCountMap{
				"how are you": 1,
			},
		},
		{
			name: "should support utf-8 characters",
			text: "ᚻᛖ ᚳᚹᚫᚦ ᚦᚫᛏ !!  ",
			want: sequenceCountMap{
				"ᚻᛖ ᚳᚹᚫᚦ ᚦᚫᛏ": 1,
			},
		},
	}
	for _, tt := range tests {
		r := strings.NewReader(tt.text)
		sequenceCountMapStream := make(chan sequenceCountMap)

		t.Run(tt.name, func(t *testing.T) {
			go processText(r, sequenceCountMapStream)

			got := <-sequenceCountMapStream

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getTopKSequences() = %v, want %v", got, tt.want)
			}
		})
	}
}
