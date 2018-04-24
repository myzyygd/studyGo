package library

import "errors"

type MusicEntry struct {
	Id      int
	Name    string
	ArtList string
	Genre   string
	source  string
	Type    string
}
type MusicManger struct {
	musics []MusicEntry
}

func NewMusicManger() *MusicManger {
	return &MusicManger{make([]MusicEntry, 0)}
}

func (m *MusicManger) Len() int {
	return len(m.musics)
}

func (m *MusicManger) Get(index int) (music *MusicEntry, err error) {
	if index < 0 || index > len(m.musics) {
		return nil, errors.New("index out of range")
	}
	return &m.musics[index], nil
}

func (m *MusicManger) Add(music *MusicEntry) {
	m.musics = append(m.musics, *music)
}
func (m *MusicManger) Find(name string) *MusicEntry {
	if len(m.musics) == 0 {
		return nil
	}
	for _, value := range m.musics {
		if value.Name == name {
			return &value
		}
	}
	return nil
}

func (m *MusicManger) Remove(index int) *MusicEntry {
	if index < 0 || index >= len(m.musics) {
		return nil
	}
	removeMusic := &m.musics[index]
	//中间元素
	if index < len(m.musics)-1 {
		m.musics = append(m.musics[:index], m.musics[:index+1]...)
	} else if index == 0 {
		m.musics = make([]MusicEntry, 0)
	} else {
		m.musics = append(m.musics[:index-1])
	}

	return removeMusic
}

func (m *MusicManger) RemoveByName(name string) *MusicEntry {
	var isPos int = -1
	for index, value := range m.musics {
		if value.Name == name {
			isPos = index
			break
		}
	}
	if isPos < 0 {
		return nil
	}
	removeMusic := m.Remove(isPos)
	return removeMusic
}
