package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(deleteCommand)
}

var deleteCommand = &cobra.Command{
	Use:   "delete",
	Short: "Delete specific todo with ID",
	Long:  "Delete the Todo with specific ID",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("Only accept 1 ID for deleting")
		}
		ID := args[0]
		records, err := readOrCreateCSV()

		if err != nil {
			return err
		}

		updatedRecords := [][]string{}

		for _, record := range records {
			if record[0] == ID {
				continue
			}
			if len(record) > 0 {
				updatedRecords = append(updatedRecords, record)
			}
		}

		if len(updatedRecords) == len(records) {
			return fmt.Errorf("%s not found", ID)
		}

		err = writeCSV(updatedRecords)

		return err
	},
}
