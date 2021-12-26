package _weather

type Tabler interface {
	TableName() string
}

type Weather struct {
	ID            int    `json:"id" gorm:"column:id" desc:"Identify"`
	Name          string `json:"name" gorm:"column:name" desc:"Name of city"`
	DataInNext12H string `json:"data" gorm:"column:data" `
}

func (tbl Weather) TableName() string {
	return "weathers"
}
