package main

func Merge(a, b []int64) []int64 {
	merged := a[:]
	merged = append(merged, b...)

	ret := make([]int64, 0)
	m := make(map[int64]struct{}, 0)
	for _, el := range merged {
		if _, ok := m[el]; ok == false {
			m[el] = struct{}{}
			ret = append(ret, el)
		}
	}
	return ret
}
