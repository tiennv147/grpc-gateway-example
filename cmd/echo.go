package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	"github.com/tiennv147/grpc-gateway-example/pb"
)

// echoCmd represents the echo command
var echoCmd = &cobra.Command{
	Use:   "echo",
	Short: "Example echo gRPC service CLI client",
	Run: func(cmd *cobra.Command, args []string) {
		var opts []grpc.DialOption
		conn, err := grpc.Dial(demoAddr, opts...)
		if err != nil {
			grpclog.Fatalf("fail to dial: %v", err)
		}
		defer conn.Close()
		client := pb.NewCrudServiceClient(conn)

		msg, err := client.Echo(context.Background(), &pb.StringMessage{Value: strings.Join(os.Args[2:], " ")})
		println(msg.Value)

	},
}

func init() {
	RootCmd.AddCommand(echoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// echoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// echoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
