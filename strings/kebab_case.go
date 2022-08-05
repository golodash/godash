package strings

func KebabCase(input string) string {
	return ScreamingDelimited(input, '-', "", false)
}
