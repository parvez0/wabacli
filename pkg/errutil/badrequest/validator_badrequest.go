package badrequest

import "fmt"

type BadRequest struct {
	Code int
	Title string
	Description string
}

func (r *BadRequest) Error() string {
	return fmt.Sprintf("%s; %s", r.Title, r.Description)
}
