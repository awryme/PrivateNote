package notestorage

import (
	"andreyko0/PrivateNote/packages/notes"
	"time"
)

type Storage struct {
	notes map[string]*notes.Note
}

func NewStorage() *Storage {
	return &Storage{
		make(map[string]*notes.Note, 0),
	}
}

func (s *Storage) DeleteNote(id string) {
	delete(s.notes, id)
}

func (s *Storage) SelfDestroy(n *notes.Note) {
	go func() {
		<-time.After(time.Until(n.Deadline))
		s.DeleteNote(n.Id)
	}()
}

func (s *Storage) GetNote(id string) *notes.Note {
	if !s.NoteExists(id) {
		return nil
	}
	return s.notes[id]
}

func (s *Storage) AddNote(n *notes.Note) {
	s.notes[n.Id] = n
	s.SelfDestroy(n)
}

func (s *Storage) NoteExists(id string) bool {
	_, ok := s.notes[id]
	return ok
}

func (s *Storage) Len() int {
	return len(s.notes)
}
