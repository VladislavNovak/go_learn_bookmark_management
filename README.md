# О программе

Позволяет управлять структурой (на основе `map`), которая добавляет, содержит и удаляет записи.

Из примечательного можно отметить функции:

`printSortMap` позволяет распечатать отсортированный `map`

`makeKeysFromMap` позволяет создать массив ключей из `map`

`getKeyToDelete` использует метод slices.Contains, возвращающий true, если key содержится в sliceKeys

```go
// import sort
func printSortMap(t map[string]string) {
	sliceKeys := makeKeysFromMap(t)

	sort.Strings(sliceKeys)

	for _, v := range sliceKeys {
		fmt.Println(v, t[v])
	}
}

func makeKeysFromMap(t map[string]string) []string {
	sliceKeys := make([]string, 0, len(t))
	for k := range t {
		sliceKeys = append(sliceKeys, k)
	}

	return sliceKeys
}

// import slices
func getKeyToDelete(t map[string]string) string {
  key := getUserInput("ключ")
  sliceKeys := makeKeysFromMap(t)

  if slices.Contains(sliceKeys, key) {
    return key
  }
}
```

PS Создана в рамках учебного процесса