package _01

func changeArrayString(arr []string) []string {
	for i, value := range arr {
		if value == "stupid" {
			arr[i] = "smart"
		}
		if value == "weak" {
			arr[i] = "strong"
		}
	}
	return arr
}
