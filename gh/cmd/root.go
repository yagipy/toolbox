package cmd

import (
  "fmt"
  "github.com/mitchellh/go-homedir"
  "github.com/spf13/cobra"
  "github.com/spf13/viper"
  "os"
  "os/exec"
)


var cfgFile string


var rootCmd = &cobra.Command{
  Use:   "gh",
  Short: "A brief description of your application",
  Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

  	Run: func(cmd *cobra.Command, args []string) {
      branchName, _:= cmd.Flags().GetString("cop")
      if branchName != "" {
        checkout_out, _ := exec.Command("git", "checkout", "-b", branchName).CombinedOutput()
        fmt.Println(string(checkout_out))
        push_out, _ := exec.Command("git", "push", "-u", "origin", branchName).CombinedOutput()
        fmt.Println(string(push_out))
      }

      addFileName, _:= cmd.Flags().GetString("add")
      if addFileName != "" {
        add_out, _:= exec.Command("git", "add", addFileName).CombinedOutput()
        fmt.Println(string(add_out))
      }

      commitMsg, _:= cmd.Flags().GetString("cm")
      if commitMsg != "" {
        commit_out, _:= exec.Command("git", "commit", "-m", commitMsg).CombinedOutput()
        fmt.Println(string(commit_out))
      }

      prBranch, _:= cmd.Flags().GetString("pr")
      if prBranch == "o" {
          exec.Command("git", "push").Start()

          exec.Command("hub", "browse").Run()
      }
    },
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

func init() {
  cobra.OnInitialize(initConfig)

  rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gh.yaml)")

  rootCmd.Flags().StringP("cop", "b", "", "checkout -b and push")
  rootCmd.Flags().StringP("add", "a", "", "add")
  rootCmd.Flags().StringP("cm", "c", "", "commit -m")
  rootCmd.Flags().StringP("pr", "p", "", "pull-request")
}
func initConfig() {
  if cfgFile != "" {
    viper.SetConfigFile(cfgFile)
  } else {
    home, err := homedir.Dir()
    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }

    viper.AddConfigPath(home)
    viper.SetConfigName(".gh")
  }

  viper.AutomaticEnv()

  if err := viper.ReadInConfig(); err == nil {
    fmt.Println("Using config file:", viper.ConfigFileUsed())
  }
}

