package cmds

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// AddCmd Represent add cobra Command
var AddCmd = &cobra.Command{
	Use: "add",
	Short: "Add two integer numbers and print result to stdout",
	Run: func (cmd *cobra.Command, args []string){
		addInt(args)
	},
}



func addInt(args []string) {
	var sum int

	for _, ival := range args {
		itemp, err := strconv.Atoi(ival)

		if err != nil {
			fmt.Println(err)
		}

		sum += itemp
	}

	fmt.Printf("Sum of numbers %s is %d\n", strings.Join(args, ","), sum)
}