// module/register/dto/register_response.go
package dto

type RegisterResponse struct {
    ID       uint   `json:"id"`
    Username string `json:"username"`
    Email    string `json:"email"`
}
