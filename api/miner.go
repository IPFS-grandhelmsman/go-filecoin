package api

import (
	"context"

	"gx/ipfs/QmQsErDt8Qgw1XrsXf2BpEzDgGWtB1YLsTAARBup5b6B9W/go-libp2p-peer"
	cid "gx/ipfs/QmZFbDTY9jfSBms2MchvYM9oYRbAF19K7Pby47yDBfpPrb/go-cid"

	"github.com/filecoin-project/go-filecoin/address"
	"github.com/filecoin-project/go-filecoin/types"
)

// Miner is the interface that defines methods to manage miner operations.
type Miner interface {
	Create(ctx context.Context, fromAddr address.Address, pledge *types.BytesAmount, pid peer.ID, collateral *types.AttoFIL) (address.Address, error)
	UpdatePeerID(ctx context.Context, fromAddr, minerAddr address.Address, newPid peer.ID) (*cid.Cid, error)
	AddAsk(ctx context.Context, fromAddr, minerAddr address.Address, size *types.BytesAmount, price *types.AttoFIL) (*cid.Cid, error)
}