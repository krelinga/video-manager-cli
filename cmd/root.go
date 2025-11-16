/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"

	"buf.build/gen/go/krelinga/proto/connectrpc/go/krelinga/video_manager/disc/v1/discv1connect"
	discv1 "buf.build/gen/go/krelinga/proto/protocolbuffers/go/krelinga/video_manager/disc/v1"
	"connectrpc.com/connect"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "vm",
	Short: "video-manager CLI tool",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: run,
}

func run(cmd *cobra.Command, args []string) {
	fmt.Println("connecting to server address:", serverAddr)
	client := discv1connect.NewDiscServiceClient(http.DefaultClient, serverAddr)

	req := connect.NewRequest(&discv1.ListInboxRequest{})
	resp, err := client.ListInbox(cmd.Context(), req)
	if err != nil {
		fmt.Println("error listing inbox:", err)
		os.Exit(1)
		return
	}

	fmt.Println("inbox videos:")
	for _, dir := range resp.Msg.Directories {
		fmt.Printf("- %s\n", dir)
	}
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var serverAddr string

func init() {
	rootCmd.PersistentFlags().StringVar(&serverAddr, "server", "", "server address")
}
