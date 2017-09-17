package universes

type Universe struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Universes []Universe
