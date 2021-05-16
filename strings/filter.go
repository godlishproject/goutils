package strings

func Filter(items []string, filter func(string) bool) []string {
	var res []string
	for _, elm := range items {
		if filter == nil && elm != "" {
			res = append(res, elm)
		}
		if filter != nil && filter(elm) {
			res = append(res, elm)
		}
	}
	return res
}
