package web

type CustomerUpdateRequest struct {
	CustomerID string `validate:"required" json:"customer_id"`
	Name       string `validate:"required,min=1,max=100" json:"name"`
	Email      string `validate:"required,email" json:"email"`
	Phone      string `validate:"required" json:"phone"`
	Address    string `json:"address"`
	LoyaltyPts int    `json:"loyalty_points"`
}
