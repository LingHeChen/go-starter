package hello

type HelloReq struct {
	Name string `form:"name" json:"name" binding:"required"`
}

type HelloResp struct {
	Reply string `json:"reply"`
	Time  string `json:"time"`
}
