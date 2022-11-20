package topksequences

import (
	"reflect"
	"testing"
)

func Test_mergeMapCounts(t *testing.T) {
	type args struct {
		map1 SequenceCountMap
		map2 SequenceCountMap
	}
	tests := []struct {
		name string
		args args
		want SequenceCountMap
	}{
		{
			name: "test 1",
			args: args{
				map1: SequenceCountMap{
					"a": 1,
					"b": 2,
				},
				map2: SequenceCountMap{
					"a": 2,
					"b": 1,
					"c": 1,
				},
			},
			want: SequenceCountMap{
				"a": 3,
				"b": 3,
				"c": 1,
			},
		},
		{
			name: "test 2",
			args: args{
				map1: SequenceCountMap{},
				map2: SequenceCountMap{},
			},
			want: SequenceCountMap{},
		},
		{
			name: "empty map 2",
			args: args{
				map1: SequenceCountMap{
					"word 1": 10,
					"word 2": 20,
				},
				map2: SequenceCountMap{},
			},
			want: SequenceCountMap{
				"word 1": 10,
				"word 2": 20,
			},
		},
		{
			name: "empty map 1",
			args: args{
				map2: SequenceCountMap{
					"word 1": 10,
					"word 2": 20,
				},
				map1: SequenceCountMap{},
			},
			want: SequenceCountMap{
				"word 1": 10,
				"word 2": 20,
			},
		},
		{
			name: "add counts in maps",
			args: args{
				map2: SequenceCountMap{
					"word 1": 10,
					"word 2": 20,
				},
				map1: SequenceCountMap{
					"word 2": 20,
					"word 3": 20,
				},
			},
			want: SequenceCountMap{
				"word 1": 10,
				"word 2": 40,
				"word 3": 20,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeMaps(tt.args.map1, tt.args.map2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeMaps() = %v, want %v", got, tt.want)
			}
		})
	}
}
