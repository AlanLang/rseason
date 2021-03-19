package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var confirm string

var rootCmd = &cobra.Command{
	Use:   "rseason",
	Short: "rseason is a very fast to rname season files",
	Long:  `rseason is a very fast to rname season files`,
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("continue to rename? [Y/n]")
		fmt.Scanln(&confirm)
		if confirm != "Y" && confirm != "y" && confirm != "" {
			os.Exit(0)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		seasonName, _ := cmd.Flags().GetString("name")
		rname(seasonName)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("name", "n", "1", "the season name you wan't to set")
	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
}

func rname(seasonName string) {
	path, _ := os.Getwd()
	files, _ := ioutil.ReadDir(path)
	for index, f := range files {
		if f.Name() != ".DS_Store" && f.Name() != "@eaDir" {
			fst := strings.Split(f.Name(), ".")
			fileType := fst[len(fst)-1]
			newName := getSeasonName(seasonName) + getEpisodesName(index+1) + "." + fileType
			err := os.Rename(path+"/"+f.Name(), path+"/"+newName)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("文件[" + f.Name() + "]已成功命名为[" + newName + "]")
			}
		}
	}
	fmt.Println("文件重命名完毕")
}

func padStart(name string) string {
	if utf8.RuneCountInString(name) == 1 {
		return "0" + name
	}
	return name
}

func getSeasonName(seasonName string) string {
	return "S" + padStart(seasonName)
}

func getEpisodesName(index int) string {
	return "E" + padStart(strconv.Itoa(index))
}
