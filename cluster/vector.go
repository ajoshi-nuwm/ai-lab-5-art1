package cluster

import "sort"

type PropertyVector struct {
	vector map[string]bool
}

func NewPropertyVector(vector map[string]bool) *PropertyVector {
	return &PropertyVector{vector}
}

func (propertyVector *PropertyVector) And(another *PropertyVector) *PropertyVector {
	resultVector := make(map[string]bool)
	for k, v := range propertyVector.vector {
		resultVector[k] = v && another.vector[k]
	}
	return NewPropertyVector(resultVector)
}

func (propertyVector *PropertyVector) Len() int {
	return len(propertyVector.vector)
}

func (propertyVector *PropertyVector) weight() int {
	var c = 0
	for _, v := range propertyVector.vector {
		if v {
			c++
		}
	}
	return c
}

func (propertyVector PropertyVector) String() string {
	var result = ""
	strings := make([]string, 0)
	for k, v := range propertyVector.vector {
		if v {
			strings = append(strings, k)
		}
	}
	sort.Strings(strings)
	for _, s := range strings {
		result += s + " "
	}
	return result
}
