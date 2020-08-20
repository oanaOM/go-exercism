package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequene of each run in councurrency
func ConcurrentFrequency(input []string) FreqMap {

	fm := FreqMap{}
	c := make(chan FreqMap, 10)

	for _, s := range input {
		go func(s string) {
			c <- Frequency(s)
		}(s)
	}

	for range input {
		m := <-c

		for k, v := range m {
			fm[k] += v
		}
	}
	return fm
}
