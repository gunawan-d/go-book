## 🚀 Introduction  

Welcome to the **go-books** project! This is a REST API for managing books and categories, secured with authentication.  

### A. 🔐 Authentication  

This project supports two authentication methods:  

1. **Basic Auth Middleware**  
   - Uses a username and password in the HTTP Authorization header.  
   - Simple and useful for internal APIs.  

2. **JSON Web Token (JWT) Middleware**  
   - Requires user login via:  
     ```
     POST /api/users/login
     ```
   - On successful login, a JWT token is returned and must be included in requests.  
   - Ideal for secure token-based authentication.  

---

### 🛠 User Authentication Routes  

These routes do **not** require authentication:  

#### 1. **User Login**  
- `POST /api/users/login` → User login  
```bash
curl --location 'http://localhost:8080/api/users/login' \
--header 'Content-Type: application/json' \
--data '{
    "username": "account",
    "password": "pass"
}'
```
**Response:**
```json
{
  "token": "eyJhbGcitokaOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

#### 2. **User Registration**  
- `POST /api/users/register` → User registration  
```bash
curl --location 'http://localhost:8080/api/users/register' \
--header 'Content-Type: application/json' \
--data '{
    "username": "account",
    "password": "pass"
}'
```
**Response:**
```json
{   
    "message":"User registered successfully"
}
```

---

### B. 📚 Category Routes (Requires Authentication)  

These routes manage book categories and require authentication:  

#### 1. **Get All Categories**  
- `GET /api/categories` → Get all categories  
```bash
curl --location 'http://localhost:8080/api/categories' \
--header 'Authorization: Bearer **********' \
--header 'Content-Type: application/json'
```
**Response:**
```json
{
    {
        "id": 1,
        "name": "Technology"
    },
    {
        "id": 2,
        "name": "Technology"
    },
    {
        "id": 3,
        "name": "Sains"
    }
}
```

#### 2. **Create New Category**  
- `POST /api/categories` → Create a new category  
```bash
curl --location 'http://localhost:8080/api/categories' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer **********' \
--data '{
    "name": "Math"
}'
```
**Response:**
```json
{
    "message": "Category created successfully"

}
```

#### 3. **Get Category Details by ID**  
- `GET /api/categories/:id` → Get category details by ID  
```bash
curl --location 'http://localhost:8080/api/categories/4' \
--header 'Authorization: Bearer **********' \
--header 'Content-Type: application/json'
```
**Response:**
```json
{
    "category": {
        "id": 2,
        "name": "Technology"
    }
}
```

#### 4. **Delete Category by ID**  
- `DELETE /api/categories/:id` → Delete a category by ID  
```bash
curl --location --request DELETE 'http://localhost:8080/api/categories/4' \
--header 'Authorization: Bearer **********'
```
**Response:**
```json
{
  "message": "Category deleted successfully"
}
```

#### 5. **Get Books by Category**  
- `GET /api/categories/:id/books` → Get books by category  
```bash
curl --location 'http://localhost:8080/api/categories/2/books' \
--header 'Authorization: Bearer **********' \
--header 'Content-Type: application/json'
```
**Response:**
```json
{
    "books": [
        "Terbitlah terang"
    ]
}
```

### C. 📚 Books API

#### 🔐 Authentication
All book-related endpoints require authentication using JSON Web Token (JWT). Include the token in the request header as follows:
```
Authorization: Bearer ******
```

---

## ✅ Validation Rules
1. **`release_year` must be between 1980 and 2024**
2. **Title, author, total_page, and category_id are required fields**
3. **Total pages must be a positive number**

---

#### 📖 Endpoints

#### 1️. Get All Books
Retrieve a list of all books.
```
GET /api/books
```
##### Request:
```sh
curl --location 'http://localhost:8080/api/books' \
--header 'Authorization: Bearer ******'
```
##### Response:
```json
[
    {
        "id": 5,
        "title": "The Hest Lini Slicer",
        "author": "Gunawan Baskara",
        "release_year": 2022,
        "total_page": 150,
        "thickness": "tebal",
        "category_id": 2
    },
    {
        "id": 6,
        "title": "Terbitlah terang",
        "author": "Husain Bagaskara",
        "release_year": 1999,
        "total_page": 10,
        "thickness": "tipis",
        "category_id": 3
    }
]
```

#### 2️. Create a New Book
Add a new book to the collection.
```
POST /api/books
```
##### Request:
```sh
curl --location 'http://localhost:8080/api/books' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer ******' \
--data '{
    "title": "Terbitlah terang",
    "author": "Husain Bagaskara",
    "release_year": 1999,
    "total_page": 10,
    "category_id": 3
}'
```
##### Response:
```json
{"message":"Book created successfully"}
```

**Validation Errors:**
```json
{
    "error": "release_year must be between 1980 and 2024"
}
```
```json
{
    "error": "total_page must be a positive number"
}
```

#### 3. Get Book by ID
Retrieve details of a specific book by its ID.
```
GET /api/books/{id}
```
##### Request:
```sh
curl --location 'http://localhost:8080/api/books/5' \
--header 'Authorization: Bearer ******'
```
##### Response:
```json
{"book":{
    "id":5,
    "title":"The Hest Lini Slicer",
    "author":"Gunawan Baskara",
    "release_year":2022,
    "total_page":150,
    "thickness":"tebal",
    "category_id":2
}}
```

**Error Response (Book Not Found):**
```json
{
    "error": "Book not found"
}
```

**Success Response:**
```json
{
    "message": "Book deleted successfully"
}
```


#### 4️. Delete Book by ID
Remove a book from the collection.
```
DELETE /api/books/{id}
```
##### Request:
```sh
curl --location --request DELETE 'http://localhost:8080/api/books/1' \
--header 'Authorization: Bearer ******'
```

**Error Response (Book Not Found):**
```json
{
    "error": "Book not found"
}
```

**Validation Errors:**
```json
{
    "error": "release_year must be between 1980 and 2024"
}
```

**Success Response:**
```json
{
    "message": "Book updated successfully"
}
```

---
Ensure you replace `******` with your valid JWT token in all requests.


