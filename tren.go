package trenlly

import "errors"

type Trening struct {
	Id        int    `json:"-"`
	Time      int    `json:"time" db:"time"`
	Weight    int    `json:"weight" db:"weight"`
	Height    int    `json:"height" db:"height"`
	MainGroup string `json:"maingroup" db:"main_group"`
}
type Result struct {
	Recomendation string `json:"recomendation,omitempty"`
	Solution      string `json:"solution,omitempty"`
	Terms         string `json:"terms,omitempty"`
	Comment       string `json:"comment,omitempty"`
}

type UpdateTrenItem struct {
	Time      *int    `json:"time"`
	Weight    *int    `json:"weight" `
	Height    *int    `json:"height" `
	MainGroup *string `json:"maingroup" `
}

func (i *UpdateTrenItem) Validate() error {
	if i.Height == nil && i.MainGroup == nil && i.Weight == nil && i.Time == nil {
		return errors.New("trening struct are null")

	}

	return nil
}
