package util

func Unshift[T any](arr []T) (popped T, new []T, overflow bool) {
	if len(arr) == 0 {
		overflow = true
		return
	}
	popped = arr[0]
	if len(arr) == 1 {
		new = []T{}
	} else {
		new = arr[1:]
	}
	return
}

func UnshiftMany[T any](arr []T, count int) (popped []T, new []T, overflow bool) {
	if count < 0 || len(arr) <= count {
		overflow = true
		return
	}

	popped = arr[:count]
	if len(arr) == count {
		new = []T{}
	} else {
		new = arr[count:]
	}
	return
}
