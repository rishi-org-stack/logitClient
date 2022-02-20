package idea

type (
	Idea struct {
		// ID           primitive.ObjectID `bson:"_id,omitempty"`
		Name         string `json:"name,omitempty"`
		Description  string `json:"description"`
		Type         string `json:"type"`
		ConcernedFor string `json:"concernedFor"`
		Problem      string `json:"problem"`
		AccessKey    string `json:"accessKey,omitempty"`
	}
)
