package models

type Todo struct {
	Page        int               `json:"page"`
	Per_page    int               `json:"per_page"`
	Total       string            `json:"total"`
	Total_pages bool              `json:"total_pages"`
	Data        []Data            `json:"data"`
	Support     map[string]string `json:"support"`
}

type Data struct {
	Id         int    `json:"id"`
	Email      string `json:"email"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Avatar     string `json:"avatar"`
}
