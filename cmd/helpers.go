package cmd

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"syscall"
	"text/tabwriter"
	"time"

	"github.com/mergestat/timediff"
)

func loadFile(filepath string) (*os.File, error) {
	f, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("failed to open file for reading")
	}

	if err := syscall.Flock(int(f.Fd()), syscall.LOCK_EX); err != nil {
		_ = f.Close()
		return nil, err
	}

	return f, nil
}

func closeFile(f *os.File) error {
	syscall.Flock(int(f.Fd()), syscall.LOCK_UN)
	return f.Close()
}
func writeCSV(records [][]string) error {
	file, err := loadFile(filename)

	if err != nil {
		return err
	}

	defer closeFile(file)

	// Truncate the file to remove all existing data
	if _, err := file.Seek(0, io.SeekStart); err != nil {
		return err
	}
	if err := file.Truncate(0); err != nil {
		return err
	}

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if len(records) != 1 {
		printCSV(records)
	}

	return writer.WriteAll(records)
}
func readOrCreateCSV() ([][]string, error) {
	_, err := os.Stat(filename)

	if os.IsNotExist(err) {
		initialData := [][]string{
			{"ID", "Description", "CreatedAt", "IsComplete"},
		}

		err = writeCSV(initialData)
		if err != nil {
			return nil, err
		}

		return initialData, nil
	}

	file, err := loadFile(filename)
	if err != nil {
		return nil, err
	}

	defer closeFile(file)
	reader := csv.NewReader(file)
	return reader.ReadAll()

}

func printCSV(records [][]string) error {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', tabwriter.DiscardEmptyColumns)
	for i, record := range records {
		if i == 0 {
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\t\n", record[0], record[1], record[2], record[3])
			continue
		}
		timeValue, err := time.Parse(time.RFC3339, record[2])
		if err != nil {
			return err
		}
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t\n", record[0], record[1], timediff.TimeDiff(timeValue), record[3])
	}
	return w.Flush()
}
