package models

type Art struct {
	Art           []Artist
	SearchArtist  []Artist
	Relation      []string
	Client        string
	CreationDate  []int
	MinCreateDate int
	MaxCreateDate int
	MinFirstAlbum int
	MaxFirstAlbum int
}
