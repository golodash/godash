package slices

func Join(slice []string, separator string) (string, error) {
	var result string
	for i := 0; i < len(slice); i++ {
		if len(result) == 0 {
			result = slice[0]
		} else {
			result += separator + slice[i]
		}
	}
	return result, nil
}
