package _weather

type Tabler interface {
	TableName() string
}

type Weather struct {
	Name          string `json:"name" gorm:"column:name" desc:"Name of city"`
	DataInNext12H string `json:"data" gorm:"column:data" `
}

type WeatherBaseResponse struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

func (tbl Weather) TableName() string {
	return "weathers"
}

func (tbl WeatherBaseResponse) TableName() string {
	return "weathers"
}
