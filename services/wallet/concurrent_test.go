package wallet

import (
	"context"
	"errors"
	"math/big"
	"sort"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestConcurrentErrorInterrupts(t *testing.T) {
	concurrent := NewConcurrentDownloader(context.Background())
	var interrupted bool
	concurrent.Add(func(ctx context.Context) error {
		select {
		case <-ctx.Done():
			interrupted = true
		case <-time.After(10 * time.Second):
		}
		return nil
	})
	err := errors.New("interrupt")
	concurrent.Add(func(ctx context.Context) error {
		return err
	})
	concurrent.Wait()
	require.True(t, interrupted)
	require.Equal(t, err, concurrent.Error())
}

func TestConcurrentCollectsTransfers(t *testing.T) {
	concurrent := NewConcurrentDownloader(context.Background())
	concurrent.Add(func(context.Context) error {
		concurrent.Push(Transfer{})
		return nil
	})
	concurrent.Add(func(context.Context) error {
		concurrent.Push(Transfer{})
		return nil
	})
	concurrent.Wait()
	require.Len(t, concurrent.Get(), 2)
}

type balancesFixture []*big.Int

func (f balancesFixture) BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	index := int(blockNumber.Int64())
	if index > len(f)-1 {
		return nil, errors.New("balance unknown")
	}
	return f[index], nil
}

type batchesFixture [][]Transfer

func (f batchesFixture) GetTransfersByNumber(ctx context.Context, number *big.Int) (rst []Transfer, err error) {
	index := int(number.Int64())
	if index > len(f)-1 {
		return nil, errors.New("unknown block")
	}
	return f[index], nil
}

func TestConcurrentEthDownloader(t *testing.T) {
	type options struct {
		balances balancesFixture
		batches  batchesFixture
		result   []Transfer
		last     *big.Int
	}
	type testCase struct {
		desc    string
		options options
	}
	for _, tc := range []testCase{
		{
			desc: "NoBalances",
			options: options{
				last:     big.NewInt(3),
				balances: balancesFixture{big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0)},
			},
		},
		{
			desc: "LastBlock",
			options: options{
				last:     big.NewInt(3),
				balances: balancesFixture{big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(10)},
				batches:  batchesFixture{{}, {}, {}, {{BlockNumber: big.NewInt(3)}, {BlockNumber: big.NewInt(3)}}},
				result:   []Transfer{{BlockNumber: big.NewInt(3)}, {BlockNumber: big.NewInt(3)}},
			},
		},
		{
			desc: "ChangesInEveryBlock",
			options: options{
				last:     big.NewInt(3),
				balances: balancesFixture{big.NewInt(0), big.NewInt(3), big.NewInt(7), big.NewInt(10)},
				batches:  batchesFixture{{}, {{BlockNumber: big.NewInt(1)}}, {{BlockNumber: big.NewInt(2)}}, {{BlockNumber: big.NewInt(3)}}},
				result:   []Transfer{{BlockNumber: big.NewInt(1)}, {BlockNumber: big.NewInt(2)}, {BlockNumber: big.NewInt(3)}},
			},
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			concurrent := NewConcurrentDownloader(ctx)
			downloadEthConcurrently(
				concurrent, tc.options.balances, tc.options.batches,
				common.Address{}, zero, tc.options.last)
			concurrent.Wait()
			require.NoError(t, concurrent.Error())
			rst := concurrent.Get()
			require.Len(t, rst, len(tc.options.result))
			sort.Slice(rst, func(i, j int) bool {
				return rst[i].BlockNumber.Cmp(rst[j].BlockNumber) < 0
			})
			for i := range rst {
				require.Equal(t, tc.options.result[i].BlockNumber, rst[i].BlockNumber)
			}
		})
	}
}
