package commands

import (
	"github.com/theQRL/zond/chain/block/genesis"
	"github.com/theQRL/zond/cli/flags"
	"github.com/theQRL/zond/keys"
	"github.com/theQRL/zond/wallet"
	"github.com/urfave/cli/v2"
)

func getGenesisSubCommands() []*cli.Command {
	return []*cli.Command {
		{
			Name:  "generate",
			Usage: "Generates new genesis block & prestate file",
			Flags: []cli.Flag {
				flags.NetworkIDFlag,
				&cli.StringFlag {
					Name: "stake-txs-filename",
					Value: "stake_transactions.json",  // TODO: Move this to Dev Config
				},
				&cli.StringFlag {
					Name: "validators-dilithium-keys",
					Value: "dilithium_keys",
				},
				flags.WalletFile,
				flags.XMSSIndexFlag,
				flags.OTSKeyIndexFlag,
				&cli.StringFlag {
					Name: "genesisFilename",
					Value: "genesis.yml",  // TODO: Move this to Dev Config
				},
				&cli.StringFlag {
					Name: "preStateFilename",
					Value: "prestate.yml",  // TODO: Move this to Dev Config
				},
			},
			Action: func(c *cli.Context) error {
				w := wallet.NewWallet(c.String("wallet-file"))

				xmss, err := w.GetXMSSByIndex(c.Uint("xmss-index"))
				if err != nil {
					return err
				}
				xmss.SetOTSIndex(c.Uint("ots-key-index"))

				d := keys.NewDilithiumKeys(c.String("validators-dilithium-keys"))
				return genesis.GenerateGenesis(c.Uint64("network-id"),
					c.String("stake-txs-filename"),
					d.GetDilithiumGroup(),
					xmss,
					c.String("genesisFilename"),
					c.String("preStateFilename"))
			},
		},
	}
}

func AddGenesisCommand(app *cli.App) {
	app.Commands = append(app.Commands, &cli.Command{
	Name: "genesis",
		Usage: "Helps to generate a genesis block",
		Flags: []cli.Flag {

		},
		Subcommands: getGenesisSubCommands(),
	})
}
