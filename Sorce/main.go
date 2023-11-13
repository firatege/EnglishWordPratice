package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func DeleteformArray(array []string, word string) []string { //This func is about delete every word uses
	index := -1
	for i := 0; i < len(array); i++ {

		if array[i] == word {
			index = i
			break
		}
	}
	if index == 0 {
		return array[1:]
	} else if index == len(array)-1 {

		return array[:len(array)-1]
	} else {

		a := array[index+1:]
		b := array[:(index)]
		return append(a, b...)
	}
}

func checkAndRemoveAnswer(table map[string]string, keys []string) { // THİS func is about ask the quetsion and take input from user
	word := keys[rand.Intn(len(keys))]            //RANDOM WORD
	fmt.Printf("%s kelimesinin turkcesi: ", word) //İNPUT
	// bufio.Scanner kullanarak bir satır okuma yapılır
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		answer := scanner.Text()

		// Kullanıcının girişini temizle
		cleanedAnswer := strings.ToLower(strings.TrimSpace(answer))

		if cleanedAnswer == strings.ToLower(table[word]) {
			fmt.Println("SİMDİKİ array:", keys)
			fmt.Println("Doğru Cevap")
			keys = DeleteformArray(keys, word)
			fmt.Println("SONRAKİ array:", keys)
		} else {
			fmt.Println("Yanlış Cevap!")
		}
	} else {
		fmt.Println("Giriş Hatası:", scanner.Err())
	}

	checkAndRemoveAnswer(table, keys)
}

func main() {

	var table = map[string]string{
		"Although":        "Gerçi",
		"Borrows":         "Ödünç & alır",
		"Relatives":       "Akrabalar",
		"Straightforward": "Doğrudan",
		"Initially":       "İlk olarak",
		"Purpose":         "Amaç",
		"Current":         "Şu anki & mevcut",
		"Concurrent":      "Eşzamanlı",
		"Dependencies":    "Bağımlılıklar",
		"Implementations": "Uygulamalar",
		"Announced":       "Duyurulan",
		"Environment":     "Çevre & ortam",
		"Patterns":        "Desenler",
		"Concurrency":     "Eşzamanlılık",
		"Statement":       "İfade",
		"Concise":         "Öz",
		"Embedding":       "Gömme",
		"External":        "Harici",
		"Instructions":    "Talimatlar",
		"Distribution":    "Dağılım",
		"Discuss":         "Tartışmak",
		"Consists":        "Oluşur",
		"Punctuation":     "Noktalama",
		"Throughout":      "Boyunca",
		"Immutable":       "Değişmez",
		"Assurance":       "Güvence",
		"Describe":        "Tanımlamak",
		"Expressions":     "İfadeler",
		"Refer":           "Başvurmak",
		"Enumeration":     "Numaralandırma",
		"Column":          "Sütun",
		"Compare":         "Karşılaştırma",
	}

	var keys []string
	for key := range table {
		keys = append(keys, key)
	}
	checkAndRemoveAnswer(table, keys)
}
