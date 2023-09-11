package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"code.vegaprotocol.io/vega/core/examples/nullchain"
	config "code.vegaprotocol.io/vega/core/examples/nullchain/config"
	"code.vegaprotocol.io/vega/protos/vega"
	walletpb "code.vegaprotocol.io/vega/protos/vega/wallet/v1"
)

func randomOrder(marketID string, side vega.Side, now time.Time) *walletpb.SubmitTransactionRequest {
	rand.Seed(time.Now().UnixNano())
	price := 1000
	percent := (price / 100)
	// perturb the price a little
	perturbation := rand.Intn(2*percent) - percent

	return nullchain.OrderTxn(marketID, uint64(price+perturbation), 10, side, vega.Order_TYPE_LIMIT, now.Add(356*24*time.Hour))
}

func runScenario(conn *nullchain.Connection, w *nullchain.Wallet) error {
	fmt.Println("Creating parties...")
	parties, err := w.MakeParties(3)
	if err != nil {
		return err
	}
	// Delete when done
	defer w.DeleteParties(parties)

	// Give them funds
	if err := nullchain.FillAccounts(config.GoveranceAsset, "10000000000", parties); err != nil {
		return err
	}
	if err := nullchain.FillAccounts(config.NormalAsset, "10000000000", parties); err != nil {
		return err
	}

	now, err := conn.VegaTime()
	if err != nil {
		return err
	}

	fmt.Println("Creating a market...")
	market, err := nullchain.CreateMarketAny(w, conn, parties[0], parties[1], parties[2])
	if err != nil {
		return err
	}

	fmt.Println("")
	fmt.Println("Submitting Orders....")

	for i := 0; i < 100; i++ {
		buy := randomOrder(market.Id, vega.Side_SIDE_BUY, now)
		sell := randomOrder(market.Id, vega.Side_SIDE_SELL, now)

		// Submit a buy and a sell order
		w.SubmitTransaction(conn, parties[1], buy)
		w.SubmitTransaction(conn, parties[2], sell)

		if err := nullchain.MoveByDuration(config.BlockDuration); err != nil {
			return err
		}

	}

	return nil

	err = nullchain.SettleMarket(w, conn, parties[0])
	if err != nil {
		return err
	}

	fmt.Println()
	fmt.Println("Market Settled.")

	return nil
}

func main() {
	var err error

	w := nullchain.NewWallet(config.WalletFolder, config.Passphrase)
	conn, err := nullchain.NewConnection()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer conn.Close()

	if err := runScenario(conn, w); err != nil {
		fmt.Printf("%+v", err)
		os.Exit(1)
	}
}
