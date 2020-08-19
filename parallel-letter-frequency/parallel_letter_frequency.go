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
	var c chan FreqMap = make(chan FreqMap)

	for _, s := range input {
		go func(s string) {
			c <- Frequency(s)
		}(s)
	}

	for i := 0; i < len(input); i++ {
		m := <-c

		for k, v := range m {
			fm[k] += v
		}
	}
	return fm
}
