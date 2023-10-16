package model

type Event struct {
	Data `json:"DATA"`
}

type Data struct {
	Page       int   `json:"page"`
	PagesCount int   `json:"pages_count"`
	RowsCount  int   `json:"rows_count"`
	Rows       []Row `json:"rows"`
}

type Row struct {
	Author `json:"author"`
	Time   string `json:"time"`
	Params `json:"params"`
}

type Author struct {
	MoId     int    `json:"mo_id"`
	UserId   int    `json:"user_id"`
	UserName string `json:"user_name"`
}

type Params struct {
	IndicatorToMyId int `json:"indicator_to_my_id"`
	Period          `json:"period"`
	Platform        string `json:"platform"`
}

type Period struct {
	End     string `json:"end"`
	Start   string `json:"start"`
	TypeId  int    `json:"type_id"`
	TypeKey string `json:"type_key"`
}
