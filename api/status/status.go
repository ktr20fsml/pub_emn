package status

type Status struct {
	Message string `json:"message"`
}

type Statuses struct {
	Message []*string `json:"message"`
}
