package request

type CreateCategoryRequest struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}

type UpdateCategoryRequest struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}
