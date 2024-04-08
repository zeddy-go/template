package http

type req struct {
	A int `form:"a" json:"a"`
}

type Something struct {
	B int
}
