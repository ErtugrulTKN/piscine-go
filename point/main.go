package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func PrintStr(str string) { // PrintStr fonksiyonu, bir string'i karakter karakter yazdırmak için kullanılır. Her karakteri z01.PrintRune ile yazdırır.
	for _, i := range str {
		z01.PrintRune(i)
	}
}

func setPoint(ptr *point) { // setPoint fonksiyonu, bir point yapısının "x" ve "y" alanlarını sırasıyla 42 ve 21 değerleri ile ayarlar.
	ptr.x = 42
	ptr.y = 21
}

func PrintInt(a int) { // PrintInt fonksiyonu, bir tamsayıyı karakter karakter yazdırmak için kullanılır. İlk olarak tamsayının her basamağını hesaplamak ve ardından z01.PrintRune ile karakter olarak yazdırmak için özyinelemeli bir işlem yapar.
	r := '0'
	if a/10 > 0 {
		PrintInt(a / 10) // parantez içindeki bölüm ifadesinin değerinin int olması için printınt içinde yaptık eger farklı bi şekilde tanımlayıp printrune ile bastırsak da olabilirdi
	}
	for i := 0; i < a%10; i++ {
		r++
	}
	z01.PrintRune(r)
}

func main() { // main fonksiyonunda, bir point yapısı olan points oluşturulur ve setPoint fonksiyonu ile "x" ve "y" değerleri atanır.
	var points point
	// points := &point{} pointer olduğunu belli etmek için süslü parantez kullandık
	// Ardından "x =" yazısı, "PrintStr" ile karakter karakter yazdırılır, ardından "PrintInt" ile "points.x" değeri yazdırılır. "PrintStr" ile "," ve "y =" yazıları yazdırılır ve "PrintInt" ile "points.y" değeri yazdırılır.
	setPoint(&points)

	PrintStr("x = ")
	PrintInt(points.x)
	PrintStr(", y = ")
	PrintInt(points.y)
	z01.PrintRune('\n')

	// fmt.Printf("x = %d, y = %d\n",points.x, points.y)
}
