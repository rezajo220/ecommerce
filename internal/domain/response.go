package domain

type ProductResponse struct {
	Message string   `json:"message" example:"Product created successfully"`
	Data    *Product `json:"data"`
}

type ProductListResponseWrapper struct {
	Message string               `json:"message" example:"Products retrieved successfully"`
	Data    *ProductListResponse `json:"data"`
}

type BrandResponse struct {
	Message string `json:"message" example:"Brand created successfully"`
	Data    *Brand `json:"data"`
}

type BrandListResponse struct {
	Message string  `json:"message" example:"Brands retrieved successfully"`
	Data    []Brand `json:"data"`
}

type MessageResponse struct {
	Message string `json:"message" example:"Operation completed successfully"`
}

type ErrorResponse struct {
	Error string `json:"error" example:"Something went wrong"`
}
