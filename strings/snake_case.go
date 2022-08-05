package strings

func SnakeCase(input string) string {
	return ScreamingDelimited(input, '_', "", false)
}
