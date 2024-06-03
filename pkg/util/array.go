package util

import "golang.org/x/exp/constraints"

func InArray[T constraints.Ordered](label T, labels []T) bool {
	for _, v := range labels {
		if label == v {
			return true
		}
	}
	return false
}
