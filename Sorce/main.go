package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func DeleteformArray(array []string, word string) {
	// Silmek istediğiniz eleman
	elementToDelete := word

	// Belirli bir elemanı sil
	indexToDelete := -1
	for i, value := range array {
		if value == elementToDelete {
			indexToDelete = i
			break
		}
	}

	if indexToDelete != -1 {
		array = append(array[:indexToDelete], array[indexToDelete+1:]...)
		fmt.Printf("%s elemanı silindi.\n", elementToDelete)
	} else {
		fmt.Printf("%s elemanı bulunamadı.\n", elementToDelete)
	}

	// Sonucu yazdır
	fmt.Println("Yeni string dizisi:", array)
	fmt.Println("Yeni dizi boyutui: ", len(array))
}

func checkAndRemoveAnswer(table map[string]string, keys []string) {
	word := keys[rand.Intn(len(keys))]            //RANDOM WORD
	fmt.Printf("%s kelimesinin turkcesi: ", word) //İNPUT
	// bufio.Scanner kullanarak bir satır okuma yapılır
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		answer := scanner.Text()

		// Kullanıcının girişini temizle
		cleanedAnswer := strings.ToLower(strings.TrimSpace(answer))

		if cleanedAnswer == strings.ToLower(table[word]) {
			fmt.Println("Doğru Cevap")
			DeleteformArray(keys, word)
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
		"Relatives":       " Akrabalar",
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
