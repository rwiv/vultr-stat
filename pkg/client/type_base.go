package client

type Meta struct {
	Total int32  `json:"total"` // 리스트의 아이템 개수
	Links *Links `json:"links"`
}

type Links struct {
	Next string `json:"next"`
	Prev string `json:"prev"`
}
