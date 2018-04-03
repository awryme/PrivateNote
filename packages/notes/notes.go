package notes

import (
	"encoding/json"
	"time"
)

type Note struct {
	Id       string    `json:"id"`
	Content  string    `json:"content"`
	Deadline time.Time `json:"deadline"`
}

func (n *Note) SetDeadline(mins int) {
	n.Deadline = time.Now().Add(time.Minute * time.Duration(mins))
}

func NewNote(id string, cont string, ttl int) *Note {
	n := &Note{
		Id:      id,
		Content: cont,
	}
	n.SetDeadline(ttl)
	return n
}

func (n *Note) ToJSON() []byte {
	bs, err := json.Marshal(n)
	if err != nil {
		return nil
	}
	return bs
}
