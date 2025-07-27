package facade

import "fmt"

// Facade Pattern Example
// This pattern provides a simplified interface to a complex subsystem.
// In this example, we create a home theater system with a facade that simplifies the process of starting a movie.
// The HomeTheaterFacade struct provides a simple method to start the movie, which
// internally manages the interactions with various components like TV, Lights, and Speakers.
// The TV, Lights, and Speakers structs represent the complex subsystems.

// Complex subsystem
type TV struct{}

func (t *TV) On() string {
	return "TV is turned ON"
}
func (t *TV) SetInput(input string) string {
	return fmt.Sprintf("TV input set to %s", input)
}

// Complex subsystem
type Lights struct{}

func (l *Lights) Dim() string {
	return "Lights are dimmed"
}

// Complex subsystem
type Speakers struct{}

func (s *Speakers) On() string {
	return "Speakers are turned ON"
}

// Grouping the complex subsystems into a facade
type HomeTheaterFacade struct {
	tv       *TV
	lights   *Lights
	speakers *Speakers
}

func NewHomeTheaterFacade() *HomeTheaterFacade {
	return &HomeTheaterFacade{
		tv:       &TV{},
		lights:   &Lights{},
		speakers: &Speakers{},
	}
}

func (h *HomeTheaterFacade) StartMovie() string {
	output := ""
	output += h.lights.Dim() + "\n"
	output += h.tv.On() + "\n"
	output += h.tv.SetInput("HDMI") + "\n"
	output += h.speakers.On() + "\n"
	output += "Enjoy your movie!\n"
	return output
}

func Facade_DP() {
	homeTheater := NewHomeTheaterFacade()

	// Start the movie using the facade, and everything will be set up automatically
	fmt.Print(homeTheater.StartMovie())
}
