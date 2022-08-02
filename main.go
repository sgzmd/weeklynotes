package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/araddon/dateparse"
	"github.com/jessevdk/go-flags"
)

type Options struct {
	Directory string `short:"d" long:"directory" description:"Directory to scan" default:"."`
}

type Notes struct {
	Date  time.Time
	Notes []string
	Name  string
}

var options Options
var parser = flags.NewParser(&options, flags.Default)

// Checks the passed date is within the last 7 days from the current date
func isDateWithinLast7Days(date time.Time) bool {
	now := time.Now()
	diff := now.Sub(date)
	return diff.Hours() < 24.0*7
}

// Opens the file and finds the position of the last markdown header.
func findLastHeaderBlockInMarkdownFile(path string) (Notes, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	dateFound := false
	date := time.Time{}
	notes := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()

		if dateFound {
			notes = append(notes, line)
		}

		if strings.HasPrefix(line, "#") {
			dateString := strings.Replace(line, "# ", "", -1)
			dt, err := dateparse.ParseAny(dateString)
			if err == nil {
				if isDateWithinLast7Days(dt) {
					log.Printf("Found header block: %s with date %+v", dateString, dt)
					dateFound = true
					date = dt
				}
			}
		}
	}
	if dateFound {
		log.Printf("Found notes: %+v", notes)
		fileName := filepath.Base(path)

		return Notes{
			Date:  date,
			Notes: notes,
			Name:  fileName,
		}, nil

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return Notes{}, fmt.Errorf("no notes found")
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
	notes := make([]Notes, 0)
	filepath.Walk(options.Directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ".md" {
			mds = append(mds, path)
			note, err := findLastHeaderBlockInMarkdownFile(path)
			if err == nil {
				notes = append(notes, note)
			}
		}
		return nil
	})

	// log.Printf("Notes found: %+v", notes)

	// create file with the name set to current date
	fileName := time.Now().Format("2006-01-02") + ".md"
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// write the header
	header := fmt.Sprintf("# %s\n", time.Now().Format("2006-01-02"))
	_, err = file.WriteString(header)
	if err != nil {
		log.Fatal(err)
	}

	// write the notes
	for _, note := range notes {
		file.WriteString(fmt.Sprintf("# %s on %s\n", note.Name, note.Date.Format("2006-01-02")))
		for _, line := range note.Notes {
			_, err = file.WriteString(line + "\n")
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
