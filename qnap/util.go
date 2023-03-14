package qnap

var c = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")

/*
encodePwd: from QNAPTool.ezEncode:
*/
func encodePwd(in string) string {
	var C = []rune(in)
	var y []rune

	e := len(C)
	A := 0
	B := rune(0)
	for A < e {
		B = C[A] & 255
		A++
		if A == e {
			y = append(y, c[B>>2], c[(B&3)<<4])
			y = append(y, '=', '=')
			break
		}
		z := C[A]
		A++
		if A == e {
			y = append(y, c[B>>2], c[((B&3)<<4)|((z&240)>>4)], c[(z&15)<<2], '=')
			break
		}
		x := C[A]
		A++
		y = append(y, c[B>>2], c[((B&3)<<4)|((z&240)>>4)], c[((z&15)<<2)|((x&192)>>6)], c[x&63])
	}
	return string(y)
}
