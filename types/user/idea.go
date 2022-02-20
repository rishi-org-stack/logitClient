package user

type (
	User struct {
		// ID     primitive.ObjectID `json:"id" bson:"_id,omitempty"`
		Name string `json:"name"`
		// AuthID primitive.ObjectID `json:"authID" bson:"auth_id"`
		// Ideas map[string]*Status `json:"ideas" bson:"ideas"`
	}
	UserAggregate struct {
		User
		// Auth au.AuthRequest `json:"auth"`
	}
	// Idea   idea.Idea
	Status struct {
		MarkedAs  string `json:"markedAs"`
		Deadline  string `json:"deadline"`
		CreatedOn string `json:"createdOn"`
	}
)
