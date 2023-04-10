package main

import (
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

// Define the model for the application
type model struct {
	text   string
	cursor int
}

// Define the message type for user input
type msg textinput.Msg

// Define the update function to handle user textinput
func update(msg tea.Msg, mdl model) (model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			// Exit the program on Ctrl+C or ESC
			return mdl, tea.Quit
		case "enter":
			// Save the text to a file on Enter
			err := saveText(mdl.text)
			if err != nil {
				fmt.Println("Error saving text:", err)
			} else {
				fmt.Println("Text saved to text.txt")
			}
			return mdl, nil
		default:
			// Update the model with the new text
			mdl.text, mdl.cursor = textinput.Update(msg, mdl.text, mdl.cursor)
			return mdl, nil
		}
	default:
		return mdl, nil
	}
}

// Define the view function to render the TUI
func view(mdl model) string {
	return fmt.Sprintf(
		"Enter some text:\n\n%s",
		textinput.View(mdl.text, "> ", mdl.cursor),
	)
}

// Define a function to save the text to a file
func saveText(text string) error {
	// Open a file for writing
	file, err := os.Create("text.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the text to the file
	_, err = file.WriteString(text)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	// Initialize the model with an empty string and a cursor position of 0
	initialModel := model{"", 0}

	// Initialize the program with the model and the update function
	p := tea.NewProgram(initialModel, update)

	// Start the program and render the TUI
	if err := p.Start(); err != nil {
		fmt.Println("Error starting program:", err)
	}
}
