/* Collatz Sanısı: “tüm sayıların 1’e indirgenmesi mümkündür” tezini ortaya atmıştır.
Probleme göre , herhangi bir sayı seçiyoruz, sayı tekse 3 ile çarpıp 1 ekliyoruz , sayı çiftse 2 ye bölüyoruz.
Bu işlemleri sayı üzerinde devamlı uyguluyoruz. */

package piscine

func CollatzCountdown(start int) int {
	if start <= 0 {
		return -1
	}

	out := 0
	for start != 1 {
		if out++; start%2 == 0 {
			start /= 2
		} else {
			start = 3*start + 1
		}
	}

	return out
}
