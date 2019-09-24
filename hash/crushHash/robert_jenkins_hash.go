/*
 * Функция хэша Роберта Дженкинса.
 * http://burtleburtle.net/bob/hash/evahash.html
 *
 * Хэш функция общего назначения для ключей переменной длины могут
 * использоваться в качестве контрольной суммы для обнаружения
 * случайного повреждения данных или обнаружения идентичных записей
 * в базе данных.
 * Функция обладает лавииным эффектом.
 *
 * Используется для снятия хэшей с объектов в Ceph (в CRUSH
 * алгоритме).
 */

package crushHash

import (
	"fmt"
)

type CrushHash struct {
	CrushHashSeed int64
}

func NewCrushHash(seed int64) *CrushHash {
	if seed == 0 {
		seed = 1315423911
	}
	return &CrushHash{CrushHashSeed: seed}
}

// Переменные a и b - случайные биты.
func (ch *CrushHash) hashMix(a, b, c int64) int64 {
	a = a - b
	a = a - c
	a = a ^ (c >> 13)
	b = b - c
	b = b - a
	b = b ^ (a << 8)
	c = c - a
	c = c - b
	c = c ^ (b >> 13)

	a = a - b
	a = a - c
	a = a ^ (c >> 12)
	b = b - c
	b = b - a
	b = b ^ (a << 16)
	c = c - a
	c = c - b
	c = c ^ (b >> 5)

	a = a - b
	a = a - c
	a = a ^ (c >> 3)
	b = b - c
	b = b - a
	b = b ^ (a << 10)
	c = c - a
	c = c - b
	c = c ^ (b >> 15)

	return c
}

func (ch *CrushHash) CrushHash32RJenkins1(a int64) int64 {
	hash := ch.CrushHashSeed ^ a
	b := a

	var x, y int64
	x = 231232
	y = 1232

	r := ch.hashMix(b, x, hash)
	r = ch.hashMix(y, a, r)

	return r
}

func (ch *CrushHash) CrushHash32RJenkins2(a, b int64) int64 {
	hash := ch.CrushHashSeed ^ a ^ b

	var x, y int64
	x = 231232
	y = 1232

	r := ch.hashMix(a, b, hash)
	r = ch.hashMix(x, a, r)
	r = ch.hashMix(b, y, r)

	return r
}

func (ch *CrushHash) CrushHash32RJenkins3(a, b, c int64) int64 {
	hash := ch.CrushHashSeed ^ a ^ b ^ c

	var x, y int64
	x = 231232
	y = 1232

	r := ch.hashMix(a, b, hash)
	r = ch.hashMix(c, x, r)
	r = ch.hashMix(y, a, r)
	r = ch.hashMix(b, x, r)
	r = ch.hashMix(y, c, r)

	return r
}

func (ch *CrushHash) CrushHash32RJenkins4(a, b, c, d int64) int64 {
	hash := ch.CrushHashSeed ^ a ^ b ^ c ^ d

	var x, y int64
	x = 231232
	y = 1232

	r := ch.hashMix(a, b, hash)
	r = ch.hashMix(c, d, r)
	r = ch.hashMix(a, x, r)
	r = ch.hashMix(y, b, r)
	r = ch.hashMix(c, x, r)
	r = ch.hashMix(y, d, r)

	return r
}

func (ch *CrushHash) CrushHash32RJenkins5(a, b, c, d, e int64) int64 {
	hash := ch.CrushHashSeed ^ a ^ b ^ c ^ d ^ e

	var x, y int64
	x = 231232
	y = 1232

	r := ch.hashMix(a, b, hash)
	r = ch.hashMix(c, d, r)
	r = ch.hashMix(e, x, r)
	r = ch.hashMix(y, a, r)
	r = ch.hashMix(b, x, r)
	r = ch.hashMix(y, c, r)
	r = ch.hashMix(d, x, r)
	r = ch.hashMix(y, e, r)

	return r
}

func main() {
	c := NewCrushHash(0)

	fmt.Println(c.CrushHash32RJenkins1(2131312))
	fmt.Println(c.CrushHash32RJenkins2(2131312, 4324234))
	fmt.Println(c.CrushHash32RJenkins3(2131312, 4324234, 213321321))
	fmt.Println(c.CrushHash32RJenkins4(2131312, 4324234, 213321321, 73298))
	fmt.Println(c.CrushHash32RJenkins5(2131312, 4324234, 213321321, 73298, 12321312))
}
