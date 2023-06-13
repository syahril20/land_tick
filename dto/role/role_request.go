package roledto

type CreateRole struct {
	NameRole string `json:"name_role" form:"name_role" validate:"required"`
}

type UpdateRole struct {
	NameRole string `json:"name_role" form:"name_role"`
}
