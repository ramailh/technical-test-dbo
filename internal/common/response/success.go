package response

type CommonSuccess struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}
