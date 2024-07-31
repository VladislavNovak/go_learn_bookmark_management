package main

import (
	"fmt"
	"slices"
	"sort"
)

// --- Общие функции ---

// Создаём массив ключей из map
// Сортируем их.
// Затем - итерируемся по ключам, извлекая инф из map
func printSortMap(t map[string]string, title string) {
	sliceKeys := makeKeysFromMap(t)

	sort.Strings(sliceKeys)

	fmt.Printf("-> %s:\n", title)
	for _, v := range sliceKeys {
		fmt.Println("----", v, t[v])
	}
}

// Создаём массив ключей из map
func makeKeysFromMap(t map[string]string) []string {
	sliceKeys := make([]string, 0, len(t))
	for k := range t {
		sliceKeys = append(sliceKeys, k)
	}

	return sliceKeys
}

// --- Программа ---

type bookmarkType = map[string]string

func getMenuItem(t bookmarkType) string {
	for {
		userInput := getUserInput("пункт меню")
		for k := range t {
			if userInput == k {
				return userInput
			}
		}
		fmt.Println("Внимание, введите 0, 1, 2 или 3")
	}
}

func getUserInput(prompt string) (userInput string) {
	fmt.Printf("Введите %s: ", prompt)
	fmt.Scan(&userInput)
	return
}

func getKeyToAdd(t map[string]string) string {
	for {
		key := getUserInput("ключ")
		sliceKeys := makeKeysFromMap(t)

		if !slices.Contains(sliceKeys, key) {
			return key
		}

		fmt.Println("-> Введённый ключ уже существует. Попробуйте снова")
	}
}

func getKeyToDelete(t bookmarkType) string {
	for {
		key := getUserInput("ключ")
		sliceKeys := makeKeysFromMap(t)

		if slices.Contains(sliceKeys, key) {
			return key
		}

		fmt.Println("-> Введённого ключа не существует. Попробуйте снова")
	}
}

func addBookmark(t bookmarkType) {
	key := getKeyToAdd(t)
	value := getUserInput("значение")
	t[key] = value
	fmt.Printf("-> Ключ (%s) со значением (%s) добавлены успешно\n", key, value)
}

func deleteBookmark(t bookmarkType) {
	if len(t) == 0 {
		fmt.Println("-> Закладок нет")
		return
	}

	key := getKeyToDelete(t)
	delete(t, key)
	fmt.Printf("-> Запись с ключом (%s) успешно удалена\n", key)
}

func infoBookmark(t bookmarkType) {
	if len(t) == 0 {
		fmt.Println("Закладок нет")
		return
	}

	printSortMap(t, "Список закладок")
}

func main() {
	mainMenu := map[string]string{
		"0": "- Завершение",
		"1": "- Посмотреть закладки",
		"2": "- Добавить закладку",
		"3": "- Удалить закладку",
	}
	bookmarks := map[string]string{}

	for {
		printSortMap(mainMenu, "Меню")

		menuItem := getMenuItem(mainMenu)

	Switch:
		switch menuItem {
		case "1":
			infoBookmark(bookmarks)
		case "2":
			addBookmark(bookmarks)
		case "3":
			deleteBookmark(bookmarks)
		case "0":
			fallthrough
		default:
			break Switch
		}

		if menuItem == "0" {
			break
		}
	}

	fmt.Println("Программа завершена")
}
