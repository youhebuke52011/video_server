package api

type UserCredental struct {
	Username string `json:"username"`
	Pwd string `json:"password"`
}

type SimpleSession struct {
	UserName string `json:"user_name"`
	TTL int64 `json:"ttl"`
}

type VideoInfo struct {
	Id string `json:"id"`
	AuthorId int64 `json:"author_id"`
	Name string `json:"name"`
	DisplayCtime string `json:"display_ctime"`
}

type Comment struct {
	Id string `json:"id"`
	VedioId string `json:"vedio_id"`
	AuthorId string `json:"author_id"`
	Content string `json:"content"`
}

type SignedUp struct {
	Success bool `json:"success"`
	SessionId string `json:"session_id"`
}