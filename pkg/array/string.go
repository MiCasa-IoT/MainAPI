package array

func NotStrContains(arr []string, str string) bool{
	for _, v := range arr{
		if v == str{
			return false
		}
	}
	return true
}
