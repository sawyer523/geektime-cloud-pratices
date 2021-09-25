package main

func main() {
	arr := []string{"I", "am", "stupid", "and", "weak"}
	for i := range arr {
		if i == 2 {
			arr[i] = "smart"
		}
		if i == 4 {
			arr[i] = "strong"
		}
	}
}
