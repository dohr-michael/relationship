package universes

type Universe struct {
	Id   string `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
}

type Universes []Universe
