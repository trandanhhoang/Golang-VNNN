package _weather

type Weather struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	DataInNext12H string `json:"data"`
}
