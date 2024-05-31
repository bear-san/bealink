package link

type Link struct {
	ID     string `json:"id" bson:"_id"`
	Values `bson:",inline"`
}

type Values struct {
	URL         string `json:"url" bson:"url"`
	ShortURL    string `json:"short_url" bson:"short_url"`
	Description string `json:"description" bson:"description"`
}
