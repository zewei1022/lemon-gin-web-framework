package request

type CreateBookInfo struct {
	Name      string  `form:"name" json:"name" validate:"require"`
	Author    string  `form:"author" json:"author" validate:"require"`
	Price     float64 `form:"price" json:"price" validate:"require"`
	Page      int     `form:"page" json:"page" validate:"require"`
	Isbn      string  `form:"isbn" json:"isbn" validate:"require"`
	Publisher string  `form:"publisher" json:"publisher" validate:"require"`
}
