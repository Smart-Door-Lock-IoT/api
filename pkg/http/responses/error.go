package httpresponses

type Error struct {
	Message string `json:"message" validate:"required"`
}
