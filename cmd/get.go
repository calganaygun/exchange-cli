/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/borakasmer/exchange-cli/parser"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Bu Cli Tool ile, kur bilgileri anlık çekilir",
	Long: `
Herhangi bir tanımlama yapılmaz ise, Dolar "-d", eğer tanımlama yapılır ise 
Euro "-e" veya Sterlin "-s" kur bilgileri Doviz.com ve Wise.com anlık olarak, Parse Edilerek ekrana basılır. 

**Doviz.com'da veya Wise.com'da bir sorun olması durumunda, bu servis hizmet veremez!!'

For example:
 .exchange get => Dolar [2022-06-08 12:26:53 - Doviz.com] : 17.1448₺
 				  Dolar [2022-06-08 12:26:46 - Wise.com] : 17.156000₺
 .exchange get -e => Euro [2022-06-08 12:26:56 - Doviz.com] : 18.3792₺
 					 Euro [2022-06-08 12:26:46 - Wise.com] : 18.337999₺
`,
	/*Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get called")
	},*/
	Run: func(cmd *cobra.Command, args []string) {
		isDolar, _ := cmd.Flags().GetBool("dolar")
		isEuro, _ := cmd.Flags().GetBool("euro")
		isSterlin, _ := cmd.Flags().GetBool("sterlin")

		if isDolar {
			getDolar()
		} else if isEuro {
			getEuro()
		} else if isSterlin {
			getSterlin()
		} else {
			getDolar()
		}
	},
}

func getDolar() {
	var exchange, exchangeFlt string
	exchange = parser.ParseWeb("DOLAR")
	exchangeWise := parser.ParseWise("USD")
	exchangeFlt = strings.Replace(exchange, ",", ".", 1)
	current_time := time.Now()
	fmt.Printf("Dolar [%s - Doviz.com] : %s₺", current_time.Format("2006-01-02 15:04:05"), exchangeFlt)
	fmt.Println()
	fmt.Printf("Dolar [%s - Wise.com] : %f₺", time.UnixMilli(exchangeWise.Time).Format("2006-01-02 15:04:05"), exchangeWise.Value)
	fmt.Println()
}
func getEuro() {
	var exchange, exchangeFlt string
	exchange = parser.ParseWeb("EURO")
	exchangeWise := parser.ParseWise("EUR")
	exchangeFlt = strings.Replace(exchange, ",", ".", 1)

	current_time := time.Now()
	fmt.Printf("Euro [%s - Doviz.com] : %s₺", current_time.Format("2006-01-02 15:04:05"), exchangeFlt)
	fmt.Println()
	fmt.Printf("Euro [%s - Wise.com] : %f₺", time.UnixMilli(exchangeWise.Time).Format("2006-01-02 15:04:05"), exchangeWise.Value)
	fmt.Println()
}

func getSterlin() {
	var exchange, exchangeFlt string
	exchange = parser.ParseWeb("STERLİN")
	exchangeWise := parser.ParseWise("GBP")
	exchangeFlt = strings.Replace(exchange, ",", ".", 1)

	current_time := time.Now()
	fmt.Printf("Sterlin [%s - Doviz.com] : %s₺", current_time.Format("2006-01-02 15:04:05"), exchangeFlt)
	fmt.Println()
	fmt.Printf("Sterlin [%s - Wise.com] : %f₺", time.UnixMilli(exchangeWise.Time).Format("2006-01-02 15:04:05"), exchangeWise.Value)
	fmt.Println()
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().BoolP("dolar", "d", false, "Get Dolar Currency")
	getCmd.Flags().BoolP("euro", "e", false, "Get Euro Currency")
	getCmd.Flags().BoolP("sterlin", "s", false, "Get Sterlin Currency")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
