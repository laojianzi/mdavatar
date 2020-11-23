package util

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"io"
	"io/ioutil"
	"math/big"
	"os"
	"strconv"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// IsFile returns true if given path exists as a file (i.e. not a directory).
func IsFile(path string) bool {
	f, e := os.Stat(path)
	if e != nil {
		return false
	}

	return !f.IsDir()
}

// OrdByte convert byte to a value between 0 and 255
func OrdByte(b byte) int {
	return int(b)
}

// OrdRune convert rune to a value between 0 and 255
func OrdRune(r rune) int {
	return int(r)
}

// Ord convert the first byte of a string to a value between 0 and 255
func Ord(s string) int {
	return OrdRune([]rune(s)[0])
}

// Iconv transform rune by former
//
// e.g: `simplifiedchinese.GBK.NewEncoder()` and `simplifiedchinese.GBK.NewDecoder()`
func Iconv(s string, former transform.Transformer) ([]byte, error) {
	return ioutil.ReadAll(transform.NewReader(bytes.NewBufferString(s), former))
}

// FirstLetter returns the uppercase first letter of a rune
// no changes should be made if there is no suitable code
func FirstLetter(s string) (string, error) {
	cnByte, err := Iconv(s, simplifiedchinese.GBK.NewEncoder())
	if err != nil {
		return s, err
	}

	firstCharHex := hex.EncodeToString(cnByte)
	firstCharDec, err := strconv.ParseInt(firstCharHex, 16, 0)
	if err != nil {
		return s, err
	}

	code := firstCharDec - 65536
	switch {
	case code >= -20319 && code <= -20284:
		return "A", nil
	case code >= -20283 && code <= -19776:
		return "B", nil
	case code >= -19775 && code <= -19219:
		return "C", nil
	case code >= -19218 && code <= -18711:
		return "D", nil
	case code >= -18710 && code <= -18527:
		return "E", nil
	case code >= -18526 && code <= -18240:
		return "F", nil
	case code >= -18239 && code <= -17923:
		return "G", nil
	case code >= -17922 && code <= -17418:
		return "H", nil
	case code >= -17417 && code <= -16475:
		return "J", nil
	case code >= -16474 && code <= -16213:
		return "K", nil
	case code >= -16212 && code <= -15641:
		return "L", nil
	case code >= -15640 && code <= -15166:
		return "M", nil
	case code >= -15165 && code <= -14923:
		return "N", nil
	case code >= -14922 && code <= -14915:
		return "O", nil
	case code >= -14914 && code <= -14631:
		return "P", nil
	case code >= -14630 && code <= -14150:
		return "Q", nil
	case code >= -14149 && code <= -14091:
		return "R", nil
	case code >= -14090 && code <= -13319:
		return "S", nil
	case code >= -13318 && code <= -12839:
		return "T", nil
	case code >= -12838 && code <= -12557:
		return "W", nil
	case code >= -12556 && code <= -11848:
		return "X", nil
	case code >= -11847 && code <= -11056:
		return "Y", nil
	case code >= -11055 && code <= -10247:
		return "Z", nil
	}

	return s, nil
}

// MtRand crypto rand a number
func MtRand(min, max int64) (int64, error) {
	v, err := Int(rand.Reader, big.NewInt(min), big.NewInt(max))
	if err != nil {
		return 0, err
	}

	return v.Int64(), nil
}

// Int returns a uniform random value in [min, max). It panics if max <= min.
// Ref crypto/rand
func Int(rand io.Reader, min, max *big.Int) (n *big.Int, err error) {
	if max.Sign() <= 0 {
		panic("util.Int: argument to Int is <= 0")
	}

	if min.Sign() >= max.Sign() {
		panic("util.Int: argument to Int is >= max")
	}

	n = new(big.Int)
	n.Sub(max, min)
	// bitLen is the maximum bit length needed to encode a value < max.
	bitLen := n.BitLen()
	if bitLen == 0 {
		// the only valid result is 0
		return
	}
	// k is the maximum byte length needed to encode a value < max.
	k := (bitLen + 7) / 8
	// b is the number of bits in the most significant byte of max-1.
	b := uint(bitLen % 8)
	if b == 0 {
		b = 8
	}

	buf := make([]byte, k)

	for {
		_, err = io.ReadFull(rand, buf)
		if err != nil {
			return nil, err
		}

		// Clear bits in the first byte to increase the probability
		// that the candidate is < max.
		buf[0] &= uint8(int(1<<b) - 1)

		n.SetBytes(buf)
		if n.Cmp(max) < 0 && n.Cmp(min) > 0 {
			return
		}
	}
}
