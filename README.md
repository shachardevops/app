# Run
docker-compose up
## GET
localhost:8080/books
## POST
localhost:8080/books/create
JSON BODY
```
    {
        "Isbn": "1",
        "Title": "Thseaaaa dddd",
        "Author": "ssss Machiavelli",
        "Price": 6.99
    }
```

## UPDATE
localhost:8080/books/{id}
JSON BODY
```
    {
        "Title": "Thseaaddddaa dddd",
        "Author": "ssss Machiavelli",
        "Price": 6.99
    }
```
## GET ONE
localhost:8080/books/{id}
## DELETE ONE
localhost:8080/books/{id}
