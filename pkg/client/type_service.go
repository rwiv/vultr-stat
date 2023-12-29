package client

type Meta struct {
	Total int32 `json:"total"`
	Links Links `json:"links"`
}

type Links struct {
	Next string `json:"next"`
	Prev string `json:"prev"`
}

type OsResponse struct {
	Os   []*OS `json:"os"`
	Meta *Meta `json:"meta"`
}

type OS struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Arch   string `json:"arch"`
	Family string `json:"family"`
}
