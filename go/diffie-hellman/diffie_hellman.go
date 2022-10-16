package diffiehellman

import (
	"math/big"
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))
var two = big.NewInt(2)

func PrivateKey(p *big.Int) *big.Int {
	private := new(big.Int).Sub(p, two)
	return private.Add(two, private.Rand(r, private))
}

// "crypto/rand"
// random, err := rand.Int(rand.Reader, max) // [0, max)

// var m = map[string]bool{}

// func PrivateKey(p *big.Int) *big.Int {
// 	rand := rand.New(rand.NewSource(time.Now().UnixNano()))

// 	for {
// 		c := new(big.Int).Add(new(big.Int).Rand(rand, new(big.Int).Sub(p, big.NewInt(2))), big.NewInt(2))
// 		_, ok := m[c.String()]
// 		if !ok {
// 			m[c.String()] = true
// 			return c
// 		}

// 		if big.NewInt(int64(len(m))).Cmp(new(big.Int).Sub(p, big.NewInt(2))) == 1 {
// 			return c
// 		}
// 	}

// }

func PublicKey(private, p *big.Int, g int64) *big.Int {
	return new(big.Int).Exp(big.NewInt(g), private, p)
}

func NewPair(p *big.Int, g int64) (priKey *big.Int, pubKey *big.Int) {
	priKey = PrivateKey(p)
	pubKey = PublicKey(priKey, p, g)

	return
}

func SecretKey(private1, public2, p *big.Int) *big.Int {
	return new(big.Int).Exp(public2, private1, p)
}
