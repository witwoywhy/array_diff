package finddiff

type Output struct {
	operation int
	increase  []int
	decrease  []int
}

func has(n int, nums []int) bool {
	for _, num := range nums {
		if n == num {
			return true
		}
	}
	return false
}

func Diff(old, new []int) Output {
	var output Output
	if len(old) == 0 {
		output.operation = 1
		output.increase = new
		return output
	}

	for _, n := range new {
		if !has(n, old) {
			output.increase = append(output.increase, n)
		}
	}

	for _, o := range old {
		if !has(o, new) {
			output.decrease = append(output.decrease, o)
		}
	}

	if len(old) == len(new) {
		output.operation = 0
	} else {
		output.operation = 2
	}

	return output
}
