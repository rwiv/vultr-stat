package array

func Concat(a []string, b []string) []string {
	newArr := make([]string, 0)
	for _, elem := range a {
		newArr = append(newArr, elem)
	}
	for _, elem := range b {
		newArr = append(newArr, elem)
	}
	return newArr
}
