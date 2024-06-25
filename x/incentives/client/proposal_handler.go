package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"

	"github.com/eniac-x-labs/dapplink/x/incentives/client/cli"
	"github.com/eniac-x-labs/dapplink/x/incentives/client/rest"
)

var (
	RegisterIncentiveProposalHandler = govclient.NewProposalHandler(cli.NewRegisterIncentiveProposalCmd, rest.RegisterIncentiveProposalRESTHandler)
	CancelIncentiveProposalHandler   = govclient.NewProposalHandler(cli.NewCancelIncentiveProposalCmd, rest.CancelIncentiveProposalRequestRESTHandler)
)
