package rest

type ScmLinkAdd struct {
	Url string `json:"url"`
}

type ScmLink struct {
	ScmLinkAdd
	ID int `json:"id"`
}

type ScmIDResponse struct {
	ID int `json:"id"`
}
