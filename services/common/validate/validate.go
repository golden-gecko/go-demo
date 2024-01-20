package validate

func NonEmptyString(value string) bool {
	return len(value) > 0
}
