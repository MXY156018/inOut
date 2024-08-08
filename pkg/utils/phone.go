package utils

func PhoneEncode(phone string) string {
	if phone == "" {
		return ""
	}
	if len(phone) <= 4 {
		return phone
	}
	var enc = ""
	var starCnt = 0
	var suffixVisibleCnt = 0
	for i := len(phone) - 1; i >= 0; i-- {
		if suffixVisibleCnt < 4 {
			enc = string(phone[i]) + enc
			suffixVisibleCnt++
		} else if i < 3 {
			enc = string(phone[i]) + enc
		} else {
			if starCnt < 4 {
				enc = "*" + enc
				starCnt++
			}
		}
	}
	return enc
}
