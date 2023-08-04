package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"srtframer/internal/services"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "srtframer",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		splitter := services.NewSplitterService()
		fileService := services.NewFileService(outPath)
		framer := services.NewSrtframer(splitter, fileService)

		err := framer.Execute(videoFilePath, srtFilePath)
		cobra.CheckErr(err)

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var videoFilePath, srtFilePath, outPath string

func init() {
	rootCmd.PersistentFlags().StringVarP(&videoFilePath, "input", "i", "", "relative path to input video file")
	rootCmd.PersistentFlags().StringVarP(&srtFilePath, "srt", "s", "", "relative path to srt file")
	rootCmd.PersistentFlags().StringVarP(&outPath, "out", "o", "./", "relative output path for frames folder")
}
