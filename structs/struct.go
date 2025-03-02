package structs

import "time"

type User struct {
	ID         int       `json:"id"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	CreatedAt  time.Time `json:"created_at"`
	CreatedBy  string    `json:"created_by"`
	ModifiedAt time.Time `json:"modified_at"`
	ModifiedBy string    `json:"modified_by"`
}

type Category struct {
	ID  		int    	`json:"id"`
	Name 		string 	`json:"name"`
}

type Book struct {
	ID			int 	`json:"id"`
	Title		string	`json:"title"`
	Author		string	`json:"author"`
	ReleaseYear	int		`json:"release_year"`
	TotalPage  	int    `json:"total_page"`
    Thickness   string `json:"thickness"`
    CategoryID  int    `json:"category_id"`

}
