package main

// hold 64 cached pitch values
var notes [1 << 6]float64
var volumes [1 << 2]float64

// fill notes and volumes tables
func init() {
	i := 0
	for i < 1<<6 {
		f, n := 65.40639132514957, 0
		for n < i {
			f *= 1.0594630943592953
			n++
		}
		notes[i] = f
		i++
	}

	j := 0
	for j < 1<<2 {
		volumes[j] = float64(j) / 2
		j++
	}
}
