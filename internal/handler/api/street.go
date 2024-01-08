package api

type CreateStreetRequest struct {
	Name string `json:"name" binding:"required"`
}

type DeleteStreetRequest struct {
	ID
	Name string `json:"name" binding:"required"`
}
