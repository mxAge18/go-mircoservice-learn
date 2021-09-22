package entity

type Product struct {
	ID    int    `json:"ID"`
	Name  string `json:"Name"`
	Stock int    `json:"Stock"`
}

type ProdRequest struct {
	Size int `form:"size" json:"Size"`
}
