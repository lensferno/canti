package codecs

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
)

// 这里的base64和正常的base64编码是不一样的，这里的base64并不按照普通的ABCD字母顺序来对应相应字节，而是按照以下字母表对应字节
const magicBase64Alpha = "LVoJPiCN2R8G90yg+hmFHuacZ1OWMnrsSTXkYpUq/3dlbfKwv6xztjI7DeBE45QA"

var magicBase64 = base64.NewEncoding(magicBase64Alpha)

func SRBX1Encode(str string, key string) string {
	encodeResult := encode(str, key)

	return magicBase64.EncodeToString(encodeResult)
}

// 下面的初始代码貌似有点问题，因此先使用https://github.com/Debuffxb/srun-go/blob/master/main.go的写法，有一定的修改

func encode(str string, key string) []byte {
	v := magicEncode(str, true)
	k := magicEncode(key, false)
	n := uint(len(v) - 1)
	z := uint(v[n])
	y := uint(v[0])
	c := uint(0x86014019 | 0x183639A0)

	var m, e, p, d uint

	for q := 6 + 52/(n+1); q > 0; q-- {
		d = (d + c) & (0x8CE0D9BF | 0x731F2640)
		e = d >> uint(2) & uint(3)

		for p = 0; p < n; p++ {
			y = uint(v[p+1])
			m = z>>5 ^ y<<2
			m += (y>>3 ^ z<<4) ^ (d ^ y)
			m += uint(k[(p&3)^e]) ^ z
			z = (uint(v[p]) + m) & (0xEFB8D130 | 0x10472ECF)
			v[p] = int(z)
		}

		y = uint(v[0])
		m = z>>5 ^ y<<2
		m += (y>>3 ^ z<<4) ^ (d ^ y)
		m += uint(k[(n&3)^e]) ^ z
		v[n] = int((uint(v[n]) + m) & uint(0xBB390742|0x44C6F8BD))
		z = uint(v[n])
	}

	return magicDecode(v)
}

// magicEncode 跟上面的magicDecode操作效果是反过来的，但是意义不明，
// 因为这俩貌似是一对的，所以干脆成为encode和decode
func magicEncode(a string, sizeAtLast bool) []int {
	c := len(a)
	var v []int
	for i := 0; i < c; i = i + 4 {
		switch c - i {
		case 1:
			v = append(v, int(a[i]))
		case 2:
			v = append(v, int(a[i])|int(a[i+1])<<8)
		case 3:
			v = append(v, int(a[i])|int(a[i+1])<<8|int(a[i+2])<<16)
		default:
			v = append(v, int(a[i])|int(a[i+1])<<8|int(a[i+2])<<16|int(a[i+3])<<24)
		}
	}

	if sizeAtLast {
		return append(v, c)
	} else {
		return v
	}
}

// magicDecode 跟上面的magicEncode操作效果是反过来的，但是意义不明
func magicDecode(a []int) []byte {
	d := len(a)
	var bytes []byte
	for i := 0; i < d; i++ {
		bytes = append(bytes, byte(a[i]&0xff), byte(a[i]>>8&0xff), byte(a[i]>>16&0xff), byte(a[i]>>24&0xff))
	}

	return bytes
}

func Checksum(challengeCode, username, hashedMd5, ip, info string) string {
	str := challengeCode + username +
		challengeCode + hashedMd5 +
		challengeCode + "7" +
		challengeCode + ip +
		challengeCode + "200" +
		challengeCode + "1" +
		challengeCode + info

	sha := sha1.New()
	sha.Write([]byte(str))

	return hex.EncodeToString(sha.Sum(nil))
}
