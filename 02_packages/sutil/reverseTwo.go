package sutil

func reverseTwo(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[1], r[j] = r[j], r[i]
	}
	return string(r)
}
