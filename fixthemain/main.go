package main

import "github.com/01-edu/z01"

const ( // Kapı duruları değişkenleri (ya açık ya kapalı)
	OPEN  = true
	CLOSE = false
)

type Door struct { // bool değişkeni tanımlandı.
	State bool
}

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func OpenDoor(ptrDoor *Door) bool { // Kapıyı açtık.
	PrintStr("Door Opening...")
	ptrDoor.State = OPEN
	return true
}

func CloseDoor(ptrDoor *Door) bool { // Kapattık.
	PrintStr("Door Closing...")
	ptrDoor.State = CLOSE
	return true
}

func IsDoorOpen(door Door) bool { // Kapı açık mı ?
	PrintStr("is the Door opened ?")
	return door.State
}

func IsDoorClose(ptrDoor *Door) bool { // Kapı kapalı mı ?
	PrintStr("is the Door closed ?")
	return !ptrDoor.State
}

func main() {
	door := &Door{}

	OpenDoor(door)
	if IsDoorClose(door) {
		OpenDoor(door)
	}
	if IsDoorOpen(*door) {
		CloseDoor(door)
	}
}
