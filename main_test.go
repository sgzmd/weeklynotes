package main

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Tests ExtractNotes method.
func TestExtractNotes(t *testing.T) {
	const text = `# 2018-01-01
Some notes
# 2022-01-01
Some more notes
	`

	// Create bufio.Scanner from text
	scanner := bufio.NewScanner(strings.NewReader(text))

	// Extract notes
	notes, err := ExtractNotes(scanner, "test.md")

	// Check if there is an error
	if err != nil {
		t.Errorf("ExtractNotes() error = %v", err)
	}

	// Check if there are notes
	if len(notes) == 0 {
		t.Errorf("ExtractNotes() error = %v", err)
	}

	// Check if the notes are correct
	if notes[0].Date.Format("2006-01-02") != "2018-01-01" {
		t.Errorf("ExtractNotes() error = %v", err)
	}
	if notes[1].Date.Format("2006-01-02") != "2022-01-01" {
		t.Errorf("ExtractNotes() error = %v", err)
	}

	assert.Equal(t, "Some notes", notes[0].Notes[0])
	assert.Equal(t, "Some more notes", notes[1].Notes[0])
	assert.Equal(t, "test.md", notes[0].Name)
}
