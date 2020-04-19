func minNumberOfFrogs(croakOfFrogs string) int {
	stc, str, sto, sta, stk := 0, 0, 0, 0, 0
	max := 0
	for _, v := range croakOfFrogs {
		switch v {
		case 'c':
			stc++
			if stc > max {
				max = stc
			}
		case 'r':
			str++
		case 'o':
			sto++
		case 'a':
			sta++
		case 'k':
			stc--
			str--
			sto--
			sta--
		}
		if stc < str || str < sto || sto < sta || sta < stk {
			return -1
		}
	}
	if stc != str || str != sto || sto != sta || sta != stk {
		return -1
	}
	return max
}