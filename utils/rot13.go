package utils


func Rot13(msg string) string {
	var res string
	for _, val := range msg {
		if val >= 'A' && val <= 'Z' {
			if val + 13 > 'Z' {
				res += string((val + 13) - 26)

			} else {
				res += string(val +13)
			}
		} else if val >= 'a' && val <= 'z' {
			if val + 13 > 'z' {
				res += string((val + 13)- 26)
			} else {
				res += string (val + 13)
			}
		} else {
			res += string(val)
		}
	}
	return res
  }