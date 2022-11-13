package codecs

import (
	"encoding/base64"
	"math"
)

// 这里的base64和正常的base64编码是不一样的，这里的base64并不按照普通的ABCD字母顺序来对应相应字节，而是按照以下字母表对应字节
const magicBase64Alpha = "LVoJPiCN2R8G90yg+hmFHuacZ1OWMnrsSTXkYpUq/3dlbfKwv6xztjI7DeBE45QA"

var magicBase64 = base64.NewEncoding(magicBase64Alpha)

func Encode(str string, key string) string {
	encodeResult := encode(str, key)
	base64Result := magicBase64.EncodeToString(int32ToAsciiBytes(encodeResult))

	return "{SRBX1}" + base64Result
}

// encode 直接照着原网页的js代码改的，只是能跑
// 因为各种类型转换从理论上来说性能可能会差了点，但是在这里的使用情景几乎是感觉不到的
// 难度其实主要还是在于动态类型语言和静态类型语言在运算时数据溢出的问题不好处理
func encode(str string, key string) []int32 {
	v := magicEncode(str, true)
	k := magicEncode(key, false)

	if len(k) < 4 {
		for i := 0; i < (4 - len(k)); i++ {
			k = append(k, 0)
		}
	}

	n := int32(len(v) - 1)
	z := v[n]
	y := v[0]
	c := int32(-1640531527)
	q := int(math.Floor(float64(6+52/(n+1)))) - 1
	d := int32(0)
	var e, p int32

	// 这里的m用int64（long）是因为后面+=的时候int32会溢出
	// 其他变量不用int64是因为直接用int64会导致位运算错误，和js的运算结果不一致
	var m int64

	for ; q >= 0; q-- {
		d = d + c&(-1)
		e = uRightShift(d, 2) & 3

		for p = 0; p < n; p++ {
			y = v[p+1]
			m = int64(uRightShift(z, 5) ^ y<<2)
			m += int64(uRightShift(y, 3) ^ z<<4 ^ (d ^ y))
			m += int64(k[p&3^e] ^ z)
			v[p] = v[p] + int32(m&(-1))
			z = v[p]
		}

		y = v[0]
		m = int64(uRightShift(z, 5) ^ y<<2)
		m += int64(uRightShift(y, 3) ^ z<<4 ^ (d ^ y))
		m += int64(k[p&3^e] ^ z)
		v[n] = v[n] + int32(m&(-1))
		z = v[n]
	}

	return magicDecode(v, false)
}

// 神秘的“初步”加密代码，把字符串一四个字符为一组转成神秘的int32数组
func magicEncode(source string, sizeOnLast bool) (result []int32) {
	data := []int32(source)
	dataLen := len(data)

	resultLen := dataLen / 4
	if sizeOnLast {
		result = make([]int32, resultLen+1)
		result[resultLen] = int32(dataLen)
	} else {
		result = make([]int32, resultLen)
	}

	for i := 0; i < dataLen; i += 4 {
		result[i>>2] = get(data, i, dataLen) | get(data, i+1, dataLen)<<8 | get(data, i+2, dataLen)<<16 | get(data, i+3, dataLen)<<24
	}

	return result
}

//和上面的是反过来的
func magicDecode(data []int32, sizeOnLast bool) (result []int32) {
	dataLength := len(data)

	c := dataLength - 1<<2

	if sizeOnLast {
		m := int(data[dataLength-1])
		if m < c-3 || m > c {
			return nil
		}
		c = m
	}

	for i := 0; i < dataLength; i++ {
		result = append(result, data[i]&0xff, uRightShift(data[i], 8)&0xff, uRightShift(data[i], 16)&0xff, uRightShift(data[i], 24)&0xff)
	}

	if sizeOnLast {
		return append(result, int32(c))
	} else {
		return result
	}
}

// 在go中实现无符号右移（>>>）
func uRightShift(number int32, shift int) int32 {
	return int32(uint32(number) >> shift)
}

func get(data []int32, index int, length int) int32 {
	if index >= length {
		return 0
	} else {
		return data[index]
	}
}

func int32ToAsciiBytes(data []int32) []byte {
	result := make([]byte, len(data))
	for i, number := range data {
		result[i] = byte(number)
	}

	return result
}
