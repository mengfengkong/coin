package cmd

import (
	"fmt"
	"github.com/mengfengkong/coin/domain"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

var client = domain.NewClient(os.Getenv("COINMARKETCAP_API_KEY"))

var RootCommand = &cobra.Command{
	Use:  "coin",
	Long: "Coin is a tool of coin. \nnotice: you should set the environment variable[COINMARKETCAP_API_KEY] first.",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var BuildGetCommand = &cobra.Command{
	Use:   "get",
	Short: "Get information about the specified coin",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Get failed, please set the symbol of coin, eg: get BTC")
			return
		}
		res, err := client.Get(args[0])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Symbol", "Price", "Volume_Change_24h"})
		for _, coin := range res.Data {
			for _, price := range coin.Quote {
				table.Append([]string{
					coin.Symbol,
					strconv.FormatFloat(price.Price, 'f', 10, 64),
					strconv.FormatFloat(price.VolumeChange24h, 'f', 10, 64),
				})
			}
		}
		table.Render() // Send output
	},
}

func init() {
	RootCommand.AddCommand(BuildGetCommand)
}
