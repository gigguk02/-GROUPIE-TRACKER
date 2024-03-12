package models

type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int32    `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Relation     map[string][]string
}
type Relations struct {
	Index []struct {
		ID              int                 `json:"id"`
		LocationAndDate map[string][]string `json:"datesLocations"`
	} `json:"index"`
}
