package reverse

func Reverse(s string) string {
	r := make([]string, len(s))
	for index, value := range s {
		r[index] = string(value)
	}
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	res := ""
	for _, value := range r {
		res += value
	}
	return res
}
