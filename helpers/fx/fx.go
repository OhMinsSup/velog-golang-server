package fx

func Filter(iter []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, a := range iter {
		if f(a) {
			vsf = append(vsf, a)
		}
	}
	return vsf
}
