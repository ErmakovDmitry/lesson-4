package homework

import (
	"strings"
	"unicode"
)

// Убирает лишние пробелы
func Step1(in <-chan string, out chan<- string) {
	defer close(out)
	for str := range in {
		// fmt.Println("str", str)
		trimmed := strings.Join(strings.Fields(str), " ")
		// fmt.Println("trimmed", trimmed)
		out <- trimmed
	}
}

// Делит строки на предложения
func Step2(in <-chan string, out chan<- string) {
	defer close(out)
	for str := range in {
		sentences := strings.Split(str, ".")
		for _, sentence := range sentences {
			trimmed := strings.TrimSpace(sentence)
			if len(trimmed) > 0 {
				out <- trimmed
			}
		}
	}
}

// Делает первые символы предложений заглавными
// step3 должен вернуть канал, в который будет записывать. Это значит, что внутри функции нужно запустить отдельную горутину, читающую in.
func Step3(in <-chan string) <-chan string {
	output := make(chan string)

	go func() {
		defer close(output)
		for sentence := range in {
			// fmt.Println("sentence", sentence)
			// title := strings.Title(strings.ToLower(sentence))
			// title := cases.Title(language.Russian, cases.NoLower).String(sentence)
			// title := sentence[0].toUpperCase() + sentence.slice(1).toLowerCase();
			title := []rune(sentence)
			title[0] = unicode.ToUpper(title[0])
			// fmt.Println("title", string(title))
			output <- string(title)
		}
	}()

	return output
}
