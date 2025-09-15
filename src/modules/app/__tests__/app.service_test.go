package app
import "testing"
func TestServiceGreet(t *testing.T) {
    s := NewService()
    got := s.Greet()
    want := "hello from Velora!"
    if got != want {
        t.Errorf("Greet() = %q; want %q", got, want)
    }
}