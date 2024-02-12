package iteration

const repeatCount = 5

func Repeat(character string, times int) string {
	var repeated string
	if times == 0 {
		return character
	}
	for i := 0; i < times; i++ {
		repeated += character
	}
	return repeated
}
