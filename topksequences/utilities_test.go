package topksequences

import (
	"reflect"
	"testing"
)

func Test_mergeMapCounts(t *testing.T) {
	type args struct {
		map1 sequenceCountMap
		map2 sequenceCountMap
	}
	tests := []struct {
		name string
		args args
		want sequenceCountMap
	}{
		{
			name: "test 1",
			args: args{
				map1: sequenceCountMap{
					"a": 1,
					"b": 2,
				},
				map2: sequenceCountMap{
					"a": 2,
					"b": 1,
					"c": 1,
				},
			},
			want: sequenceCountMap{
				"a": 3,
				"b": 3,
				"c": 1,
			},
		},
		{
			name: "test 2",
			args: args{
				map1: sequenceCountMap{},
				map2: sequenceCountMap{},
			},
			want: sequenceCountMap{},
		},
		{
			name: "empty map 2",
			args: args{
				map1: sequenceCountMap{
					"word 1": 10,
					"word 2": 20,
				},
				map2: sequenceCountMap{},
			},
			want: sequenceCountMap{
				"word 1": 10,
				"word 2": 20,
			},
		},
		{
			name: "empty map 1",
			args: args{
				map2: sequenceCountMap{
					"word 1": 10,
					"word 2": 20,
				},
				map1: sequenceCountMap{},
			},
			want: sequenceCountMap{
				"word 1": 10,
				"word 2": 20,
			},
		},
		{
			name: "add counts in maps",
			args: args{
				map2: sequenceCountMap{
					"word 1": 10,
					"word 2": 20,
				},
				map1: sequenceCountMap{
					"word 2": 20,
					"word 3": 20,
				},
			},
			want: sequenceCountMap{
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
