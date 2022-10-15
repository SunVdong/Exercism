package cipher

import (
	"strings"
)

type shift struct {
	distance int
}

type vigenere struct {
	key string
}

func NewCaesar() Cipher {
	return shift{distance: 3}
}

func NewShift(distance int) Cipher {
	if (distance >= 1 && distance <= 25) || (distance >= -25 && distance <= -1) {
		return shift{distance: distance}
	}

	return nil
}

func (c shift) Encode(input string) string {
	var sb strings.Builder
	lower := strings.ToLower(input)
	for _, r := range lower {
		if r >= 'a' && r <= 'z' {
			char := r + rune(c.distance)
			if char > 'z' {
				char = char - 26
			}
			if char < 'a' {
				char = char + 26
			}
			sb.WriteRune(char)
		}
	}

	return sb.String()
}

func (c shift) Decode(input string) string {
	var sb strings.Builder

	for _, r := range input {
		char := r - rune(c.distance)
		if char > 'z' {
			char = char - 26
		}
		if char < 'a' {
			char = char + 26
		}
		sb.WriteRune(char)
	}
	return sb.String()
}

func NewVigenere(key string) Cipher {
	var valid bool = false
	for _, r := range key {
		if !(r >= 'a' && r <= 'z') {
			return nil
		}
		if r != 'a' {
			valid = true
		}
	}
	if !valid {
		return nil
	}

	return vigenere{key: key}
}

func (v vigenere) Encode(input string) string {
	var sb strings.Builder
	keyLen := len(v.key)
	lower := strings.ToLower(input)
	count := 0
	for i := 0; i < len(lower); i++ {
		if lower[i] >= 'a' && lower[i] <= 'z' {
			ki := count
			if count >= keyLen {
				ki = count % keyLen
			}

			r := lower[i] + (v.key[ki] - 'a')
			if r > 'z' {
				r = r - 26
			}
			sb.WriteRune(rune(r))
			count++
		}
	}

	return sb.String()
}

func (v vigenere) Decode(input string) string {
	var sb strings.Builder
	keyLen := len(v.key)
	for i := 0; i < len(input); i++ {
		ki := i
		if i >= keyLen {
			ki = i % keyLen
		}

		r := input[i] - (v.key[ki] - 'a')
		if r < 'a' {
			r += 26
		}

		sb.WriteRune(rune(r))
	}

	return sb.String()
}
