func isPalin(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func recur(res *[][]string, tmp []string, s string) [][]string {
	if len(s) == 0 {
		copyTmp := make([]string, len(tmp))
		copy(copyTmp, tmp)
		*res = append(*res, copyTmp)
		return nil
	}

	for i := 1; i < len(s)+1; i++ {
		if isPalin(s[:i]) == true {
			x := append(tmp, s[:i])
			recur(res, x, s[i:])
		}
	}
	return nil
}

func partition(s string) [][]string {
	if len(s) == 0 {
		return [][]string{}
	}
	res := [][]string{}
	recur(&res, []string{}, s)
	return res
}