package admin

import "testing"

func TestServiceFindAll(t *testing.T) {
	s := NewService()
	got := s.FindAll()
	want := []string{"example"}
	if len(got) != len(want) {
		t.Fatalf("expected %v, got %v", want, got)
	}
}
