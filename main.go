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
	"github.com/sgzmd/weeklynotes/calutils"
)

type Options struct {
	Directory     string `short:"d" long:"directory" description:"Directory to scan" default:"."`
	OutputDir     string `short:"o" long:"output-dir" description:"Output directory" default:"."`
	ScanBackWeeks int    `short:"w" long:"weeks-back" description:"Scan back this many weeks" default:"1"`
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

// Computes the date of monday as indicated by the week number and a year
func CreateDateFromWeekNumberAndYear(week, year int) time.Time {
	return time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC).AddDate(0, 0, week*7)
}

// Groups all notes by the week, starting on Monday.
func GroupNotesByCalendarWeek(notes []Note) map[time.Time][]Note {

	groupedNotes := make(map[time.Time][]Note)
	for _, note := range notes {
		y, week := note.Date.ISOWeek()
		monday := calutils.MondayOfTheWeek(y, week)
		groupedNotes[monday] = append(groupedNotes[monday], note)
	}

	return groupedNotes
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
		if strings.Contains(path, "rollup") {
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

	groupedNotes := GroupNotesByCalendarWeek(notes)
	keys := make([]time.Time, 0)
	for k := range groupedNotes {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i].After(keys[j])
	})

	for idx, key := range keys {
		if idx > options.ScanBackWeeks {
			break
		}
		notes := groupedNotes[key]
		if len(notes) > 0 {
			fileName := fmt.Sprintf("%s/%s-notes.md", options.OutputDir, key.Format("2006 Jan 02"))
			log.Printf("Writing %s", fileName)
			file, err := os.Create(fileName)
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()

			file.WriteString(fmt.Sprintf("=== Note for the week of %s === \n", key.Format("Monday, January 2, 2006")))
			for _, note := range notes {
				file.WriteString(fmt.Sprintf("\n## %s on %s\n", note.Name, note.Date.Format("2006 Jan 2")))
				for _, line := range note.Notes {
					file.WriteString(fmt.Sprintf("%s\n", line))
				}
			}
		}
	}
}
