package response

type BookInfo struct {
	Name      string  `json:"name"`
	Author    string  `json:"author"`
	Price     float64 `json:"price"`
	Page      int     `json:"page"`
	Isbn      string  `json:"isbn"`
	Publisher string  `json:"publisher"`
}
