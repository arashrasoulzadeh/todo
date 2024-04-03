package requests

type Task struct {
	Title  string `json:"title"`
	Desc   string `json:"desc,omitempty"`
	Status uint8  `json:"status"`
}
