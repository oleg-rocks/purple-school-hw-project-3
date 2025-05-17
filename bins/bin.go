package bins

import (
	"encoding/json"
	"time"
)

type Bin struct {
	Id        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
}

type BinList struct {
	bins []Bin
}

func NewBin(id string, private bool, createdAt time.Time, name string) *Bin {
	return &Bin{id, private, createdAt, name}
}

func NewBinList(bins []Bin) *BinList {
	return &BinList{bins}
}

func (bin *Bin) ToBytes() ([]byte, error) {
	data, err := json.Marshal(bin)
	if err != nil {
		return nil, err
	}

	return data, nil
}
