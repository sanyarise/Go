package fibbonachi

func FindFibbonachi(N uint) uint {
	Mapa := make(map[uint]uint)
	var i uint
	for i = 0; i <= N; i++ {
		var f uint
		if i <= 2 {
			f = 1
		} else {
			f = Mapa[i-1] + Mapa[i-2]
		}
		Mapa[i] = f
	}
	return Mapa[N]
}
