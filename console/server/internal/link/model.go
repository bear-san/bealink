package link

type Link struct {
	ID     string `json:"id" bson:"_id"`
	Values `bson:",inline"`
}

type Values struct {
	URL         string `json:"url" bson:"url"`
	Path        string `json:"path" bson:"path"`
	Description string `json:"description" bson:"description"`
}
