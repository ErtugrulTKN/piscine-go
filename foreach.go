package piscine

func ForEach(a func(int), b []int) { // int için o dilimin her elemanına bir fonksiyon uygulayan fonksiyon.
	for _, value := range b {
		a(value)
	}
}
