package main

// These imports will be used later on the tutorial. If you save the file
// now, Go might complain they are unused, but that's fine.
// You may also need to run `go mod tidy` to download bubbletea and its
// dependencies.
import (
	//////////////////"errors"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	//"time"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	randomFact string
}

type factResponse struct {
	Text string `json:"text"`
}

func initialModel() model {
	return model{
		randomFact: getRandomFact(),
	}
}

func getRandomFact() string {

	requestUrl := "https://uselessfacts.jsph.pl/random.json?language=en"

	res, err := http.Get(requestUrl)

	if err != nil {
		fmt.Printf("Error: %v\n", err)

		os.Exit(1)
	}

	defer res.Body.Close()

	var data factResponse
	err = json.NewDecoder(res.Body).Decode(&data)

	if err != nil {
		fmt.Printf("Error: %v\n", err)

		os.Exit(1)
	}

	return data.Text
}

func (modelInstance model) Init() tea.Cmd {
	return nil
}

func (modelInstance model) View() string {
	s := "Random Fact: " + modelInstance.randomFact

	return s
}

func (modelInstance model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return modelInstance, tea.Quit
		case "r":
			modelInstance.randomFact = getRandomFact()
		}
	}
	return modelInstance, nil
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

}
