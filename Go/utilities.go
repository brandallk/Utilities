package utilities

type IntSlice []int

func (s1 IntSlice) IsEqual(s2 IntSlice) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
