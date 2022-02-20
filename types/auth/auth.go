package auth

type (
	AuthRequest struct {
		// ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
		Email    string `json:"email"`
		Password string `json:"password"`
		// Status   string             `json:"status" bson:"status,omitempty"`
	}
	AuthResponse struct {
		Token string `json:"token"`
	}
)
