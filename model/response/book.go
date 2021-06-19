package response

type BookInfo struct {
	ID        uint    `json:"id"`
	Name      string  `json:"name"`
	Author    string  `json:"author"`
	Price     float64 `json:"price"`
	Page      uint    `json:"page"`
	Isbn      string  `json:"isbn"`
	Publisher string  `json:"publisher"`
}
