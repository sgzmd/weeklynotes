package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/araddon/dateparse"
	"github.com/jessevdk/go-flags"
)

type Options struct {
	Directory string `short:"d" long:"directory" description:"Directory to scan" default:"."`
	OutputDir string `short:"o" long:"output-dir" description:"Output directory" default:"."`
}

type Note struct {
	Date  time.Time
	Notes []string
	Name  string
}

var options Options
var parser = flags.NewParser(&options, flags.Default)

func ExtractNotesFromFile(path string) ([]Note, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	return ExtractNotes(scanner, filepath.Base(path))
}

// Opens the file and finds the position of the last markdown header.
func ExtractNotes(scanner *bufio.Scanner, name string) ([]Note, error) {
	notes := make([]Note, 0)
	note := make([]string, 0)

	var lastDate time.Time

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") {
			dateString := strings.Replace(line, "# ", "", -1)
			dt, err := dateparse.ParseAny(dateString)
			if err == nil {
				if len(note) > 0 {
					notes = append(notes, Note{Date: lastDate, Notes: note, Name: name})
					note = make([]string, 0)
				}

				lastDate = dt
			}
		} else {
			note = append(note, line)
		}
	}
	if len(note) > 0 {
		notes = append(notes, Note{Date: lastDate, Notes: note, Name: name})
	}

	if len(notes) > 0 {
		return notes, nil
	} else {
		return notes, fmt.Errorf("no notes found")
	}
}

func main() {
	if _, err := parser.Parse(); err != nil {
		switch flagsErr := err.(type) {
		case flags.ErrorType:
			if flagsErr == flags.ErrHelp {
				os.Exit(0)
			}
			os.Exit(1)
		default:
			os.Exit(1)
		}
	}

	mds := make([]string, 0)
	notes := make([]Note, 0)
	filepath.Walk(options.Directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if strings.Contains(path, "rollup.md") {
			return nil
		}
		if filepath.Ext(path) == ".md" {
			mds = append(mds, path)
			note, err := ExtractNotesFromFile(path)
			if err == nil {
				notes = append(notes, note[:]...)
			}
		}
		return nil
	})

	// log.Printf("Notes found: %+v", notes)

	sort.Slice(notes, func(i, j int) bool {
		return notes[i].Date.After(notes[j].Date)
	})

	// create file with the name set to current date
	now := time.Now()
	fileName := fmt.Sprintf("%s/%s rollup.md", options.OutputDir, now.Format("2006-01-02"))
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// write the header
	header := fmt.Sprintf("%% %s\n", time.Now().Format("2006-01-02"))
	_, err = file.WriteString(header + "\n")
	if err != nil {
		log.Fatal(err)
	}

	_, week := notes[0].Date.ISOWeek()
	_, _ = file.WriteString(fmt.Sprintf("# Week %d\n", week))
	// write the notes
	for _, note := range notes {
		_, w := note.Date.ISOWeek()
		if w != week {
			week = w
			_, err = file.WriteString(fmt.Sprintf("# Week %d\n", week))
			if err != nil {
				log.Fatal(err)
			}
		}

		file.WriteString(fmt.Sprintf("## %s on %s\n", note.Name, note.Date.Format("2006-01-02")))
		for _, line := range note.Notes {
			_, err = file.WriteString(line + "\n")
			if err != nil {
				log.Fatal(err)
			}
		}
		file.WriteString("\n")
	}
}
