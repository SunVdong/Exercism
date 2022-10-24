package variablelengthquantity

import (
	"errors"
	"strconv"
)

func EncodeVarint(input []uint32) []byte {
	var res []byte
	for _, v := range input {
		s := strconv.FormatInt(int64(v), 2)

		var mask int64 = 0
		var tempArr []byte
		for i := len(s); i > 0; i -= 7 {
			start := i - 7
			if start <= 0 {
				start = 0
			}

			b, _ := strconv.ParseInt(s[start:i], 2, 16)
			b = b | mask
			mask = 0b10000000

			tempArr = append([]byte{byte(b)}, tempArr...)
		}

		res = append(res, tempArr...)
	}

	return res
}

func DecodeVarint(input []byte) ([]uint32, error) {
	var res []uint32

	var numTmp uint32 = 0
	for _, v := range input {
		var mask uint32 = 0b10000000
		var uv uint32 = uint32(v)
		numTmp *= 128
		numTmp += uv & ^mask
		if uv&mask != mask {
			res = append(res, numTmp)
			numTmp = 0
		}
	}
	if res == nil {
		return nil, errors.New("something is wrong")
	}

	return res, nil
}
