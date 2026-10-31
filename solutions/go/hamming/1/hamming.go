package hamming
import "fmt"
func Distance(a, b string) (int, error) {
	runasA := []rune(a)
	runasB := []rune(b)
	count := 0
	if len(runasA) != len(runasB) {
		return 0, fmt.Errorf("No tienen la misma longitud las cadenas")
	}
	for i, l := range runasA {
		if l != runasB[i] {
			count++
		}
	}
	return count, nil
}
