package link

type LinkCreateRequest struct {
	Url string `json:"url" validate:"required,url"` // validate - это go валидатор, для обработки ошибок
}

type LinkUpdateRequest struct {
	Url  string `json:"url" validate:"required,url"` // validate - это go валидатор, для обработки ошибок
	Hash string `json:"hash"`                        // validate:"required,hash"`
}
