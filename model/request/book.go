package request

type CreateBookInfo struct {
	Name      string  `form:"name" json:"name" validate:"required"`
	Author    string  `form:"author" json:"author" validate:"required"`
	Price     float64 `form:"price" json:"price" validate:"required,gt=0"`
	Page      int     `form:"page" json:"page" validate:"required,min=1"`
	Isbn      string  `form:"isbn" json:"isbn" validate:"required,len=13"`
	Publisher string  `form:"publisher" json:"publisher" validate:"required"`
}
