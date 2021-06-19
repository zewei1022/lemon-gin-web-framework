package request

type CreateBookInfo struct {
	Name      string  `form:"name" json:"name" validate:"required"`
	Author    string  `form:"author" json:"author" validate:"required"`
	Price     float64 `form:"price" json:"price" validate:"required,gt=0"`
	Page      uint    `form:"page" json:"page" validate:"required,min=1"`
	Isbn      string  `form:"isbn" json:"isbn" validate:"required,len=13"`
	Publisher string  `form:"publisher" json:"publisher" validate:"required"`
}

type UpdateBookInfo struct {
	ID        uint    `form:"id" json:"id" validate:"required"`
	Name      string  `form:"name" json:"name" validate:"omitempty"`
	Author    string  `form:"author" json:"author" validate:"omitempty"`
	Price     float64 `form:"price" json:"price" validate:"omitempty,gt=0"`
	Page      uint    `form:"page" json:"page" validate:"omitempty,min=1"`
	Isbn      string  `form:"isbn" json:"isbn" validate:"omitempty,len=13"`
	Publisher string  `form:"publisher" json:"publisher" validate:"omitempty"`
}
