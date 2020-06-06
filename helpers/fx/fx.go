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

func Contain(iters []string, selector string) (string, bool) {
	for _, iter := range iters {
		if iter == selector {
			return iter, true
		}
	}
	return "", false
}

func ContainSelector(iters []string, selector string) (string, bool) {
	for _, iter := range iters {
		if iter == selector {
			return selector, true
		}
	}
	return "", false
}
