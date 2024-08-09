package cmd

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCommand)
}

var addCommand = &cobra.Command{
	Use:   "add",
	Short: "Add new todo",
	Long:  "Add new todo to list",
	RunE: func(cmd *cobra.Command, args []string) error {

		if len(args) == 0 {
			return fmt.Errorf("Require a description to add todo")
		}

		description := strings.Join(args, " ")

		records, err := readOrCreateCSV()

		if err != nil {
			return err
		}

		records = append(records, []string{
			strconv.Itoa(len(records)),
			description,
			time.Now().Format("2006-01-02T15:04:05-07:00"),
			strconv.FormatBool(false),
		})

		err = writeCSV(records)

		return err
	},
}
