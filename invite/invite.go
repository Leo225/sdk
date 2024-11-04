package invite

import (
	"errors"
	"strings"
)

const CHARSET = "97FEMpQdLjq2ca3yGU5ZrHB84bDznYkWeRSgKoXmJh6itCuNvATsPxwVf"

var base = uint64(len(CHARSET))

type Generator struct {
	length       int
	coprime      int
	decodeFactor uint16
	maxSupport   uint64
}

func NewGenerator(length uint8) (*Generator, error) {
	return &Generator{
		length:       int(length),
		coprime:      int(minCoprime(uint64(length))),
		decodeFactor: uint16(base) * uint16(length),
		maxSupport:   pow(base, uint64(length)) - 1,
	}, nil
}

func (g *Generator) MaxSupportID() uint64 {
	return g.maxSupport
}

func (g *Generator) Encode(id uint64) (string, error) {
	if id > g.maxSupport {
		return "", errors.New("id out of range")
	}

	idx := make([]uint16, g.length)

	//diffusion
	for i := 0; i < g.length; i++ {
		idx[i] = uint16(id % base)
		idx[i] = (idx[i] + uint16(i)*idx[0]) % uint16(base)
		id /= base
	}

	//mix up
	var buf strings.Builder
	buf.Grow(g.length)
	for i := 0; i < g.length; i++ {
		n := i * g.coprime % g.length
		buf.WriteByte(CHARSET[idx[n]])
	}
	return buf.String(), nil
}

func (g *Generator) Decode(code string) uint64 {
	var idx = make([]uint16, g.length)
	for i, c := range code {
		idx[i*g.coprime%g.length] = uint16(strings.IndexRune(CHARSET, c))
	}

	var id uint64
	for i := g.length - 1; i >= 0; i-- {
		id *= base
		idx[i] = (idx[i] + g.decodeFactor - idx[0]*uint16(i)) % uint16(base)
		id += uint64(idx[i])
	}
	return id
}

func minCoprime(n uint64) uint64 {
	if n == 1 {
		return 2
	}

	for i := uint64(2); i < n; i++ {
		if isCoprime(i, n) {
			return i
		}
	}
	return n + 1
}

func isCoprime(n, m uint64) bool {
	return gcd(n, m) == 1
}

func gcd(n, m uint64) uint64 {
	if m == 0 {
		return n
	}
	return gcd(m, n%m)
}

func pow(n, m uint64) uint64 {
	sum := n
	for i := uint64(1); i < m; i++ {
		sum *= n
	}
	return sum
}
