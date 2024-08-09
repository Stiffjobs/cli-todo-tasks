package cmd

import (
	"strconv"

	"github.com/spf13/cobra"
)

func init() {
	listCommand.Flags().BoolVarP(&showAll, "all", "a", false, "whether show completed tasks")
	rootCmd.AddCommand(listCommand)
}

var showAll bool

var listCommand = &cobra.Command{
	Use:   "list",
	Short: "Print the todo list",
	Long:  `Print the todo list`,
	RunE: func(cmd *cobra.Command, args []string) error {

		records, err := readOrCreateCSV()
		if err != nil {
			return err
		}

		updated := [][]string{}

		for i, record := range records {
			if i == 0 {
				updated = append(updated, record)
				continue
			}
			completed, err := strconv.ParseBool(record[3])
			if err != nil {
				return err
			}
			if showAll {
				updated = append(updated, record)
			} else {
				if !completed {
					updated = append(updated, record)
				}
			}
		}

		printCSV(updated)

		return nil

	},
}
