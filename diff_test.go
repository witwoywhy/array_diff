package finddiff

import (
	"reflect"
	"testing"
)

func TestDiff(t *testing.T) {
	t.Run("[] => [1]", func(t *testing.T) {
		var old []int
		new := []int{1}

		want := Output{
			operation: 1,
			increase:  []int{1},
			decrease:  nil,
		}

		output := Diff(old, new)
		if !reflect.DeepEqual(output, want) {
			t.Errorf("want %v but got %v", want, output)
		}
	})

	t.Run("[0] => [1]", func(t *testing.T) {
		old := []int{0}
		new := []int{1}

		want := Output{
			operation: 0,
			increase:  []int{1},
			decrease:  []int{0},
		}
		output := Diff(old, new)
		if !reflect.DeepEqual(output, want) {
			t.Errorf("want %v but got %v", want, output)
		}
	})

	t.Run("[0, 1, 2] => [1, 4, 5]", func(t *testing.T) {
		old := []int{0, 1, 2}
		new := []int{1, 4, 5}

		want := Output{
			operation: 0,
			increase:  []int{4, 5},
			decrease:  []int{0, 2},
		}
		output := Diff(old, new)
		if !reflect.DeepEqual(output, want) {
			t.Errorf("want %v but got %v", want, output)
		}
	})

	t.Run("[0] => [0, 1, 2]", func(t *testing.T) {
		old := []int{0}
		new := []int{0, 1, 2}

		want := Output{
			operation: 2,
			increase:  []int{1, 2},
			decrease:  nil,
		}
		output := Diff(old, new)
		if !reflect.DeepEqual(output, want) {
			t.Errorf("want %v but got %v", want, output)
		}
	})

	t.Run("[0] => [4, 5, 6]", func(t *testing.T) {
		old := []int{0}
		new := []int{4, 5, 6}

		want := Output{
			operation: 2,
			increase:  []int{4, 5, 6},
			decrease:  []int{0},
		}
		output := Diff(old, new)
		if !reflect.DeepEqual(output, want) {
			t.Errorf("want %v but got %v", want, output)
		}
	})

	t.Run("[4, 5, 6] => [0]", func(t *testing.T) {
		old := []int{4, 5, 6}
		new := []int{0}

		want := Output{
			operation: 2,
			increase:  []int{0},
			decrease:  []int{4, 5, 6},
		}
		output := Diff(old, new)
		if !reflect.DeepEqual(output, want) {
			t.Errorf("want %v but got %v", want, output)
		}
	})

	t.Run("[4, 5, 6] => [5]", func(t *testing.T) {
		old := []int{4, 5, 6}
		new := []int{5}

		want := Output{
			operation: 2,
			increase:  nil,
			decrease:  []int{4, 6},
		}
		output := Diff(old, new)
		if !reflect.DeepEqual(output, want) {
			t.Errorf("want %v but got %v", want, output)
		}
	})
}
