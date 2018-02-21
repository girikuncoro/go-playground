package cmd

import (
	"fmt"
	"os"
	"strconv"

	bk "github.com/girikuncoro/go-playground/bookshelf/pkg"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Cli struct {
	Cmd  *cobra.Command
	Args []string

	rootCmd *cobra.Command
	v       *viper.Viper
	addCmd  *addCmd
	listCmd *listCmd

	db         *bk.BookDatabase
	dbHost     string
	dbUser     string
	dbPassword string
	dbPort     string
}

const (
	CliProgram    = "shelf"
	dbHostKey     = "DB_HOST"
	dbUserKey     = "DB_USER"
	dbPasswordKey = "DB_PASSWORD"
	dbPortKey     = "DB_PORT"
)

func NewCli() *Cli {
	cli := &Cli{v: viper.New()}
	cli.setDefaultConfig()
	cli.rootCmd = &cobra.Command{
		Use:   CliProgram,
		Short: "Shelf that contains books",
		Long:  "Shelf that contains books from various writers and titles",
		// Uncomment if bare app has an action associated with it:
		Run: func(cmd *cobra.Command, args []string) {},
	}

	cli.rootCmd.PersistentFlags().StringVar(&cli.dbHost, "host", cli.dbHost, "DB_HOST")
	cli.rootCmd.PersistentFlags().StringVar(&cli.dbUser, "user", cli.dbHost, "DB_USER")
	cli.rootCmd.PersistentFlags().StringVar(&cli.dbPassword, "password", cli.dbHost, "DB_PASSWORD")
	cli.rootCmd.PersistentFlags().StringVar(&cli.dbPort, "port", cli.dbHost, "DB_PORT")

	cli.v.BindPFlag(dbHostKey, cli.rootCmd.PersistentFlags().Lookup("host"))
	cli.v.BindPFlag(dbUserKey, cli.rootCmd.PersistentFlags().Lookup("user"))
	cli.v.BindPFlag(dbPasswordKey, cli.rootCmd.PersistentFlags().Lookup("password"))
	cli.v.BindPFlag(dbPortKey, cli.rootCmd.PersistentFlags().Lookup("port"))

	registerAddCmds(cli)
	registerListCmds(cli)
	return cli
}

func (cli *Cli) setDefaultConfig() {
	cli.v.SetDefault(dbHostKey, "")
	cli.v.SetDefault(dbUserKey, "")
	cli.v.SetDefault(dbPasswordKey, "")
	cli.v.SetDefault(dbPortKey, "")
	cli.setDefaultEnvironmentValue()
}

func (cli *Cli) setDefaultEnvironmentValue() {
	cli.dbHost = cli.v.GetString(dbHostKey)
	cli.dbUser = cli.v.GetString(dbUserKey)
	cli.dbPassword = cli.v.GetString(dbPasswordKey)
	cli.dbPort = cli.v.GetString(dbHostKey)
}

func (cli *Cli) Run() {
	cli.v.ReadInConfig()
	if err := cli.rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (cli *Cli) runner(runner func(*Cli) error) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		cli.Cmd = cmd
		cli.Args = args
		return runner(cli)
	}
}

func (cli *Cli) dbClient(user, password, host, port string) (bk.BookDatabase, error) {
	p, _ := strconv.Atoi(port)
	db, err := bk.NewMySQLDB(bk.MySQLConfig{
		Username: user,
		Password: password,
		Host:     host,
		Port:     p,
	})
	return db, err
}
