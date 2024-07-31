package utils

func IsValidURL(URL string) bool {
	return "http" == URL[:4]
}
