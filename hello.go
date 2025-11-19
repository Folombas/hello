package main

import "fmt"

func main() {
    // Создаем словарь
    dictionary := map[string]string{
        "hello": "привет",
        "world": "мир", 
        "go":    "идти",
        "cat":   "кот",
	"cow":	"корова",
	"dog":	"собака",
	"food": "еда",
	"fork": "вилка",
	"spoon": "ложка",
	"kitchen": "кухня",
	"spider": "паук",
	"apple": "яблоко",
	"pineapple": "ананас",
	"tea":	"чай",
    	"cucumber": "огурец",
	"milk":	"молоко",
	}


    // Запрашиваем слово у пользователя
    var word string
    fmt.Scan(&word)

    // Проверяем наличие слова в словаре
    translation, exists := dictionary[word]
    
    if exists {
        fmt.Println(translation)
    } else {
        fmt.Println("перевод не найден")
    }
}
