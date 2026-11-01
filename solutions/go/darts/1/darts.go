package darts

func Score(x, y float64) int {
	distancia := (x * x) + (y * y)
	resultado := 0
	if distancia <= 1 {
		resultado = 10
	}else if distancia <= 25 {
		resultado = 5
	}else if distancia <= 100 {
		resultado = 1
	}else {
		return 0
	}
	return resultado
}