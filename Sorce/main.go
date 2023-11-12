package main

import "math/rand"

func randomSeleceter(keys []string) string { //Selects a string randomly in slice

	return keys[rand.Intn(len(keys))]
}

func checkAndRemoveAnswer(word string, table map[string]string) {

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
	println(randomSeleceter(keys))
}
