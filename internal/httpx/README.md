# httpx

## Пример

```golang
res := httpx.NewRequest("GET", "https://shikimori.one/api/animes").
    Query("limit", "50"). // Указание параметра
    Header("Accept", "application/json"). // Указание заголовка 
    Do() // Выполнение запроса

if !res.Success() {
    // Обработка ошибки...
}

// Запрос выполнен успешно.
// Дальнейшая обработка ответа...
```