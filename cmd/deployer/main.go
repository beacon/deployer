package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/beacon/deployer/pkg/config"
	"github.com/beacon/deployer/pkg/render"
	"github.com/beacon/deployer/pkg/server"

	"github.com/spf13/cobra"
)

var cfg *config.Config

func addRunCmd(root *cobra.Command) {
	var configFile string
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Run deployer in server/worker mode",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := config.New(configFile)
			if err != nil {
				return err
			}
			flags := cmd.Flags()
			flags.StringVar(&cfg.Addr, "addr", "", "Listen port")
			if err := flags.Parse(args); err != nil {
				return err
			}
			if err := config.Validate(cfg); err != nil {
				return err
			}

			ch := make(chan os.Signal)
			signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
			switch cfg.Mode {
			case "server":
				srv := server.New(cfg)
				go func() {
					if srv.ListenAndServe(cfg); err != nil {
						log.Println("Server closed:", err)
					}
				}()
				defer srv.Shutdown()
			case "worker":
				// TODO: run worker
			default:
				return fmt.Errorf("unknown mode %s", cfg.Mode)
			}
			<-ch
			return nil
		},
	}
	flags := cmd.Flags()
	flags.StringVarP(&configFile, "config", "c", "", "Config file for deployer")
	root.AddCommand(cmd)
}

func addRenderCmd(root *cobra.Command) {
	var output string
	var input []string
	var file string
	cmd := &cobra.Command{
		Use:       "render",
		Short:     "Render some files with given grammer",
		ValidArgs: []string{"output", "input", "file"},
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("Args:", args)
			return render.Execute(nil, file, output, input...)
		},
	}
	flags := cmd.Flags()
	flags.StringArrayVarP(&input, "input", "i", nil, "Input files, can be either file or directory")
	flags.StringVarP(&output, "output", "o", "", "Output dir for rendered files")
	flags.StringVarP(&file, "file", "f", "", "File containing input values, should be either in JSON/YAML format")
	root.AddCommand(cmd)
}

func main() {
	rootCmd := &cobra.Command{
		Use: "deployer can run either as server/worker",
	}
	addRunCmd(rootCmd)
	addRenderCmd(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatalln("Failed to execute deployer:", err)
	}
}
