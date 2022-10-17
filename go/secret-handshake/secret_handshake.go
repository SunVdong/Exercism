package secret

var list = []string{"wink", "double blink", "close your eyes", "jump"}

func Handshake(code uint) (res []string) {
	reverse := code&(1<<4) == (1 << 4)

	for idx, item := range list {
		if code&(1<<idx) == (1 << idx) {
			if reverse {
				res = append([]string{item}, res...)
			} else {
				res = append(res, item)
			}
		}
	}

	return
}
