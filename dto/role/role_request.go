package roledto

type CreateRole struct {
	Id       int    `json:"id_role"`
	NameRole string `json:"name_role"`
}

type UpdateRole struct {
	Id       int    `json:"id_role"`
	NameRole string `json:"name_role"`
}
