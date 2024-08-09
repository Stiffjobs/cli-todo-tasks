package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(completeCommand)
}

var completeCommand = &cobra.Command{
	Use:   "complete",
	Short: "Complete the task",
	Long:  "Complete your todo task",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("Only accept 1 ID for deleting")
		}

		ID := args[0]

		records, err := readOrCreateCSV()

		if err != nil {
			return err
		}

		for _, record := range records {
			if record[0] == ID {
				record[3] = strconv.FormatBool(true)
			}
		}

		err = writeCSV(records)

		return err

	},
}
