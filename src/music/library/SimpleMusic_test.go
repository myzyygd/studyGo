package library

import "testing"

func TestNewMusicMangerOps(t *testing.T) {

	m := NewMusicManger()
	if m == nil {
		t.Error("NewMusicManger Failed")
	}

	if m.Len() != 0 {
		t.Error("NewMusicManger Failed, not empty ")
	}
	m0 := &MusicEntry{
		1,
		"My Heart Will go on",
		"Celion Dion",
		"Pop",
		"http://qbox.me/24501234",
		"MP3",
	}
	m.Add(m0)
	if m.Len() != 1 {
		t.Error("MusicManager Add Failed.")
	}
	m1 := m.Find(m0.Name)
	if m1 == nil {
		t.Error("MusicManager Find Failed.")
	}
	if m1.Id != m0.Id ||
		m1.Name != m0.Name ||
		m1.ArtList != m0.ArtList ||
		m1.Genre != m0.Genre ||
		m1.Source != m0.Source ||
		m1.Type != m0.Type {
		t.Error("MusicManager Find Failed,Found items NotMatch")
	}
	m1, err := m.Get(0)
	if m1 == nil {
		t.Error("MusicManager Get Failed", err)
	}
	m1 = m.Remove(0)
	if m1 == nil {
		t.Error("MusicManager Remove Failed.")
	}
}
