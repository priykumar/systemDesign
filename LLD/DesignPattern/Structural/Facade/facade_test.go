package facade

import "testing"

func TestHomeTheaterFacade_StartMovie(t *testing.T) {
	expected := `Lights are dimmed
TV is turned ON
TV input set to HDMI
Speakers are turned ON
Enjoy your movie!
`

	home := NewHomeTheaterFacade()
	actual := home.StartMovie()

	if actual != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, actual)
	}
}
