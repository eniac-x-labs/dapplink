<!-- This file is auto-generated. Please do not modify it yourself. -->
# Protobuf Documentation
<a name="top"></a>

## Table of Contents

- [dapplink/epochs/v1/genesis.proto](#dapplink/epochs/v1/genesis.proto)
    - [EpochInfo](#dapplink.epochs.v1.EpochInfo)
    - [GenesisState](#dapplink.epochs.v1.GenesisState)
  
- [dapplink/epochs/v1/query.proto](#dapplink/epochs/v1/query.proto)
    - [QueryCurrentEpochRequest](#dapplink.epochs.v1.QueryCurrentEpochRequest)
    - [QueryCurrentEpochResponse](#dapplink.epochs.v1.QueryCurrentEpochResponse)
    - [QueryEpochsInfoRequest](#dapplink.epochs.v1.QueryEpochsInfoRequest)
    - [QueryEpochsInfoResponse](#dapplink.epochs.v1.QueryEpochsInfoResponse)
  
    - [Query](#dapplink.epochs.v1.Query)
  
- [dapplink/erc20/v1/erc20.proto](#dapplink/erc20/v1/erc20.proto)
    - [RegisterCoinProposal](#dapplink.erc20.v1.RegisterCoinProposal)
    - [RegisterERC20Proposal](#dapplink.erc20.v1.RegisterERC20Proposal)
    - [ToggleTokenConversionProposal](#dapplink.erc20.v1.ToggleTokenConversionProposal)
    - [TokenPair](#dapplink.erc20.v1.TokenPair)
  
    - [Owner](#dapplink.erc20.v1.Owner)
  
- [dapplink/erc20/v1/genesis.proto](#dapplink/erc20/v1/genesis.proto)
    - [GenesisState](#dapplink.erc20.v1.GenesisState)
    - [Params](#dapplink.erc20.v1.Params)
  
- [dapplink/erc20/v1/query.proto](#dapplink/erc20/v1/query.proto)
    - [QueryParamsRequest](#dapplink.erc20.v1.QueryParamsRequest)
    - [QueryParamsResponse](#dapplink.erc20.v1.QueryParamsResponse)
    - [QueryTokenPairRequest](#dapplink.erc20.v1.QueryTokenPairRequest)
    - [QueryTokenPairResponse](#dapplink.erc20.v1.QueryTokenPairResponse)
    - [QueryTokenPairsRequest](#dapplink.erc20.v1.QueryTokenPairsRequest)
    - [QueryTokenPairsResponse](#dapplink.erc20.v1.QueryTokenPairsResponse)
  
    - [Query](#dapplink.erc20.v1.Query)
  
- [dapplink/erc20/v1/tx.proto](#dapplink/erc20/v1/tx.proto)
    - [MsgConvertCoin](#dapplink.erc20.v1.MsgConvertCoin)
    - [MsgConvertCoinResponse](#dapplink.erc20.v1.MsgConvertCoinResponse)
    - [MsgConvertERC20](#dapplink.erc20.v1.MsgConvertERC20)
    - [MsgConvertERC20Response](#dapplink.erc20.v1.MsgConvertERC20Response)
  
    - [Msg](#dapplink.erc20.v1.Msg)
  
- [dapplink/incentives/v1/incentives.proto](#dapplink/incentives/v1/incentives.proto)
    - [CancelIncentiveProposal](#dapplink.incentives.v1.CancelIncentiveProposal)
    - [GasMeter](#dapplink.incentives.v1.GasMeter)
    - [Incentive](#dapplink.incentives.v1.Incentive)
    - [RegisterIncentiveProposal](#dapplink.incentives.v1.RegisterIncentiveProposal)
  
- [dapplink/incentives/v1/genesis.proto](#dapplink/incentives/v1/genesis.proto)
    - [GenesisState](#dapplink.incentives.v1.GenesisState)
    - [Params](#dapplink.incentives.v1.Params)
  
- [dapplink/incentives/v1/query.proto](#dapplink/incentives/v1/query.proto)
    - [QueryAllocationMeterRequest](#dapplink.incentives.v1.QueryAllocationMeterRequest)
    - [QueryAllocationMeterResponse](#dapplink.incentives.v1.QueryAllocationMeterResponse)
    - [QueryAllocationMetersRequest](#dapplink.incentives.v1.QueryAllocationMetersRequest)
    - [QueryAllocationMetersResponse](#dapplink.incentives.v1.QueryAllocationMetersResponse)
    - [QueryGasMeterRequest](#dapplink.incentives.v1.QueryGasMeterRequest)
    - [QueryGasMeterResponse](#dapplink.incentives.v1.QueryGasMeterResponse)
    - [QueryGasMetersRequest](#dapplink.incentives.v1.QueryGasMetersRequest)
    - [QueryGasMetersResponse](#dapplink.incentives.v1.QueryGasMetersResponse)
    - [QueryIncentiveRequest](#dapplink.incentives.v1.QueryIncentiveRequest)
    - [QueryIncentiveResponse](#dapplink.incentives.v1.QueryIncentiveResponse)
    - [QueryIncentivesRequest](#dapplink.incentives.v1.QueryIncentivesRequest)
    - [QueryIncentivesResponse](#dapplink.incentives.v1.QueryIncentivesResponse)
    - [QueryParamsRequest](#dapplink.incentives.v1.QueryParamsRequest)
    - [QueryParamsResponse](#dapplink.incentives.v1.QueryParamsResponse)
  
    - [Query](#dapplink.incentives.v1.Query)
  
- [dapplink/inflation/v1/inflation.proto](#dapplink/inflation/v1/inflation.proto)
    - [ExponentialCalculation](#dapplink.inflation.v1.ExponentialCalculation)
    - [InflationDistribution](#dapplink.inflation.v1.InflationDistribution)
  
- [dapplink/inflation/v1/genesis.proto](#dapplink/inflation/v1/genesis.proto)
    - [GenesisState](#dapplink.inflation.v1.GenesisState)
    - [Params](#dapplink.inflation.v1.Params)
  
- [dapplink/inflation/v1/query.proto](#dapplink/inflation/v1/query.proto)
    - [QueryCirculatingSupplyRequest](#dapplink.inflation.v1.QueryCirculatingSupplyRequest)
    - [QueryCirculatingSupplyResponse](#dapplink.inflation.v1.QueryCirculatingSupplyResponse)
    - [QueryEpochMintProvisionRequest](#dapplink.inflation.v1.QueryEpochMintProvisionRequest)
    - [QueryEpochMintProvisionResponse](#dapplink.inflation.v1.QueryEpochMintProvisionResponse)
    - [QueryInflationRateRequest](#dapplink.inflation.v1.QueryInflationRateRequest)
    - [QueryInflationRateResponse](#dapplink.inflation.v1.QueryInflationRateResponse)
    - [QueryParamsRequest](#dapplink.inflation.v1.QueryParamsRequest)
    - [QueryParamsResponse](#dapplink.inflation.v1.QueryParamsResponse)
    - [QueryPeriodRequest](#dapplink.inflation.v1.QueryPeriodRequest)
    - [QueryPeriodResponse](#dapplink.inflation.v1.QueryPeriodResponse)
    - [QuerySkippedEpochsRequest](#dapplink.inflation.v1.QuerySkippedEpochsRequest)
    - [QuerySkippedEpochsResponse](#dapplink.inflation.v1.QuerySkippedEpochsResponse)
  
    - [Query](#dapplink.inflation.v1.Query)
  
- [dapplink/vesting/v1/query.proto](#dapplink/vesting/v1/query.proto)
    - [QueryBalancesRequest](#dapplink.vesting.v1.QueryBalancesRequest)
    - [QueryBalancesResponse](#dapplink.vesting.v1.QueryBalancesResponse)
  
    - [Query](#dapplink.vesting.v1.Query)
  
- [dapplink/vesting/v1/tx.proto](#dapplink/vesting/v1/tx.proto)
    - [MsgClawback](#dapplink.vesting.v1.MsgClawback)
    - [MsgClawbackResponse](#dapplink.vesting.v1.MsgClawbackResponse)
    - [MsgCreateClawbackVestingAccount](#dapplink.vesting.v1.MsgCreateClawbackVestingAccount)
    - [MsgCreateClawbackVestingAccountResponse](#dapplink.vesting.v1.MsgCreateClawbackVestingAccountResponse)
  
    - [Msg](#dapplink.vesting.v1.Msg)
  
- [dapplink/vesting/v1/vesting.proto](#dapplink/vesting/v1/vesting.proto)
    - [ClawbackVestingAccount](#dapplink.vesting.v1.ClawbackVestingAccount)
  
- [Scalar Value Types](#scalar-value-types)



<a name="dapplink/epochs/v1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## dapplink/epochs/v1/genesis.proto



<a name="dapplink.epochs.v1.EpochInfo"></a>

### EpochInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `identifier` | [string](#string) |  |  |
| `start_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `duration` | [google.protobuf.Duration](#google.protobuf.Duration) |  |  |
| `current_epoch` | [int64](#int64) |  |  |
| `current_epoch_start_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `epoch_counting_started` | [bool](#bool) |  |  |
| `current_epoch_start_height` | [int64](#int64) |  |  |






<a name="dapplink.epochs.v1.GenesisState"></a>

### GenesisState
GenesisState defines the epochs module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `epochs` | [EpochInfo](#dapplink.epochs.v1.EpochInfo) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="dapplink/epochs/v1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## dapplink/epochs/v1/query.proto



<a name="dapplink.epochs.v1.QueryCurrentEpochRequest"></a>

### QueryCurrentEpochRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `identifier` | [string](#string) |  |  |






<a name="dapplink.epochs.v1.QueryCurrentEpochResponse"></a>

### QueryCurrentEpochResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `current_epoch` | [int64](#int64) |  |  |






<a name="dapplink.epochs.v1.QueryEpochsInfoRequest"></a>

### QueryEpochsInfoRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="dapplink.epochs.v1.QueryEpochsInfoResponse"></a>

### QueryEpochsInfoResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `epochs` | [EpochInfo](#dapplink.epochs.v1.EpochInfo) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="dapplink.epochs.v1.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `EpochInfos` | [QueryEpochsInfoRequest](#dapplink.epochs.v1.QueryEpochsInfoRequest) | [QueryEpochsInfoResponse](#dapplink.epochs.v1.QueryEpochsInfoResponse) | EpochInfos provide running epochInfos | GET|/dapplink/epochs/v1/epochs|
| `CurrentEpoch` | [QueryCurrentEpochRequest](#dapplink.epochs.v1.QueryCurrentEpochRequest) | [QueryCurrentEpochResponse](#dapplink.epochs.v1.QueryCurrentEpochResponse) | CurrentEpoch provide current epoch of specified identifier | GET|/dapplink/epochs/v1/current_epoch|

 <!-- end services -->



<a name="dapplink/erc20/v1/erc20.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## dapplink/erc20/v1/erc20.proto



<a name="dapplink.erc20.v1.RegisterCoinProposal"></a>

### RegisterCoinProposal
RegisterCoinProposal is a gov Content type to register a token pair for a
native Cosmos coin.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `title` | [string](#string) |  | title of the proposal |
| `description` | [string](#string) |  | proposal description |
| `metadata` | [cosmos.bank.v1beta1.Metadata](#cosmos.bank.v1beta1.Metadata) |  | metadata of the native Cosmos coin |






<a name="dapplink.erc20.v1.RegisterERC20Proposal"></a>

### RegisterERC20Proposal
RegisterERC20Proposal is a gov Content type to register a token pair for an
ERC20 token


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `title` | [string](#string) |  | title of the proposal |
| `description` | [string](#string) |  | proposal description |
| `erc20address` | [string](#string) |  | contract address of ERC20 token |






<a name="dapplink.erc20.v1.ToggleTokenConversionProposal"></a>

### ToggleTokenConversionProposal
ToggleTokenConversionProposal is a gov Content type to toggle the conversion
of a token pair.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `title` | [string](#string) |  | title of the proposal |
| `description` | [string](#string) |  | proposal description |
| `token` | [string](#string) |  | token identifier can be either the hex contract address of the ERC20 or the Cosmos base denomination |






<a name="dapplink.erc20.v1.TokenPair"></a>

### TokenPair
TokenPair defines an instance that records a pairing consisting of a native
 Cosmos Coin and an ERC20 token address.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `erc20_address` | [string](#string) |  | address of ERC20 contract token |
| `denom` | [string](#string) |  | cosmos base denomination to be mapped to |
| `enabled` | [bool](#bool) |  | shows token mapping enable status |
| `contract_owner` | [Owner](#dapplink.erc20.v1.Owner) |  | ERC20 owner address ENUM (0 invalid, 1 ModuleAccount, 2 external address) |





 <!-- end messages -->


<a name="dapplink.erc20.v1.Owner"></a>

### Owner
Owner enumerates the ownership of a ERC20 contract.

| Name | Number | Description |
| ---- | ------ | ----------- |
| OWNER_UNSPECIFIED | 0 | OWNER_UNSPECIFIED defines an invalid/undefined owner. |
| OWNER_MODULE | 1 | OWNER_MODULE erc20 is owned by the erc20 module account. |
| OWNER_EXTERNAL | 2 | EXTERNAL erc20 is owned by an external account. |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="dapplink/erc20/v1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## dapplink/erc20/v1/genesis.proto



<a name="dapplink.erc20.v1.GenesisState"></a>

### GenesisState
GenesisState defines the module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#dapplink.erc20.v1.Params) |  | module parameters |
| `token_pairs` | [TokenPair](#dapplink.erc20.v1.TokenPair) | repeated | registered token pairs |






<a name="dapplink.erc20.v1.Params"></a>

### Params
Params defines the erc20 module params


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `enable_erc20` | [bool](#bool) |  | parameter to enable the conversion of Cosmos coins <--> ERC20 tokens. |
| `enable_evm_hook` | [bool](#bool) |  | parameter to enable the EVM hook that converts an ERC20 token to a Cosmos Coin by transferring the Tokens through a MsgEthereumTx to the ModuleAddress Ethereum address. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="dapplink/erc20/v1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## dapplink/erc20/v1/query.proto



<a name="dapplink.erc20.v1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest is the request type for the Query/Params RPC method.






<a name="dapplink.erc20.v1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse is the response type for the Query/Params RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#dapplink.erc20.v1.Params) |  |  |






<a name="dapplink.erc20.v1.QueryTokenPairRequest"></a>

### QueryTokenPairRequest
QueryTokenPairRequest is the request type for the Query/TokenPair RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `token` | [string](#string) |  | token identifier can be either the hex contract address of the ERC20 or the Cosmos base denomination |






<a name="dapplink.erc20.v1.QueryTokenPairResponse"></a>

### QueryTokenPairResponse
QueryTokenPairResponse is the response type for the Query/TokenPair RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `token_pair` | [TokenPair](#dapplink.erc20.v1.TokenPair) |  |  |






<a name="dapplink.erc20.v1.QueryTokenPairsRequest"></a>

### QueryTokenPairsRequest
QueryTokenPairsRequest is the request type for the Query/TokenPairs RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="dapplink.erc20.v1.QueryTokenPairsResponse"></a>

### QueryTokenPairsResponse
QueryTokenPairsResponse is the response type for the Query/TokenPairs RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `token_pairs` | [TokenPair](#dapplink.erc20.v1.TokenPair) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="dapplink.erc20.v1.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `TokenPairs` | [QueryTokenPairsRequest](#dapplink.erc20.v1.QueryTokenPairsRequest) | [QueryTokenPairsResponse](#dapplink.erc20.v1.QueryTokenPairsResponse) | TokenPairs retrieves registered token pairs | GET|/dapplink/erc20/v1/token_pairs|
| `TokenPair` | [QueryTokenPairRequest](#dapplink.erc20.v1.QueryTokenPairRequest) | [QueryTokenPairResponse](#dapplink.erc20.v1.QueryTokenPairResponse) | TokenPair retrieves a registered token pair | GET|/dapplink/erc20/v1/token_pairs/{token}|
| `Params` | [QueryParamsRequest](#dapplink.erc20.v1.QueryParamsRequest) | [QueryParamsResponse](#dapplink.erc20.v1.QueryParamsResponse) | Params retrieves the erc20 module params | GET|/dapplink/erc20/v1/params|

 <!-- end services -->



<a name="dapplink/erc20/v1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## dapplink/erc20/v1/tx.proto



<a name="dapplink.erc20.v1.MsgConvertCoin"></a>

### MsgConvertCoin
MsgConvertCoin defines a Msg to convert a native Cosmos coin to a ERC20 token


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `coin` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | Cosmos coin which denomination is registered in a token pair. The coin amount defines the amount of coins to convert. |
| `receiver` | [string](#string) |  | recipient hex address to receive ERC20 token |
| `sender` | [string](#string) |  | cosmos bech32 address from the owner of the given Cosmos coins |






<a name="dapplink.erc20.v1.MsgConvertCoinResponse"></a>

### MsgConvertCoinResponse
MsgConvertCoinResponse returns no fields






<a name="dapplink.erc20.v1.MsgConvertERC20"></a>

### MsgConvertERC20
MsgConvertERC20 defines a Msg to convert a ERC20 token to a native Cosmos
coin.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `contract_address` | [string](#string) |  | ERC20 token contract address registered in a token pair |
| `amount` | [string](#string) |  | amount of ERC20 tokens to convert |
| `receiver` | [string](#string) |  | bech32 address to receive native Cosmos coins |
| `sender` | [string](#string) |  | sender hex address from the owner of the given ERC20 tokens |






<a name="dapplink.erc20.v1.MsgConvertERC20Response"></a>

### MsgConvertERC20Response
MsgConvertERC20Response returns no fields





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="dapplink.erc20.v1.Msg"></a>

### Msg
Msg defines the erc20 Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `ConvertCoin` | [MsgConvertCoin](#dapplink.erc20.v1.MsgConvertCoin) | [MsgConvertCoinResponse](#dapplink.erc20.v1.MsgConvertCoinResponse) | ConvertCoin mints a ERC20 representation of the native Cosmos coin denom that is registered on the token mapping. | GET|/dapplink/erc20/v1/tx/convert_coin|
| `ConvertERC20` | [MsgConvertERC20](#dapplink.erc20.v1.MsgConvertERC20) | [MsgConvertERC20Response](#dapplink.erc20.v1.MsgConvertERC20Response) | ConvertERC20 mints a native Cosmos coin representation of the ERC20 token contract that is registered on the token mapping. | GET|/dapplink/erc20/v1/tx/convert_erc20|

 <!-- end services -->



<a name="dapplink/incentives/v1/incentives.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## dapplink/incentives/v1/incentives.proto



<a name="dapplink.incentives.v1.CancelIncentiveProposal"></a>

### CancelIncentiveProposal
CancelIncentiveProposal is a gov Content type to cancel an incentive


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `title` | [string](#string) |  | title of the proposal |
| `description` | [string](#string) |  | proposal description |
| `contract` | [string](#string) |  | contract address |






<a name="dapplink.incentives.v1.GasMeter"></a>

### GasMeter
GasMeter tracks the cumulative gas spent per participant in one epoch


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `contract` | [string](#string) |  | hex address of the incentivized contract |
| `participant` | [string](#string) |  | participant address that interacts with the incentive |
| `cumulative_gas` | [uint64](#uint64) |  | cumulative gas spent during the epoch |






<a name="dapplink.incentives.v1.Incentive"></a>

### Incentive
Incentive defines an instance that organizes distribution conditions for a
given smart contract


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `contract` | [string](#string) |  | contract address |
| `allocations` | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) | repeated | denoms and percentage of rewards to be allocated |
| `epochs` | [uint32](#uint32) |  | number of remaining epochs |
| `start_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | distribution start time |
| `total_gas` | [uint64](#uint64) |  | cumulative gas spent by all gasmeters of the incentive during the epoch |






<a name="dapplink.incentives.v1.RegisterIncentiveProposal"></a>

### RegisterIncentiveProposal
RegisterIncentiveProposal is a gov Content type to register an incentive


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `title` | [string](#string) |  | title of the proposal |
| `description` | [string](#string) |  | proposal description |
| `contract` | [string](#string) |  | contract address |
| `allocations` | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) | repeated | denoms and percentage of rewards to be allocated |
| `epochs` | [uint32](#uint32) |  | number of remaining epochs |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="dapplink/incentives/v1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## dapplink/incentives/v1/genesis.proto



<a name="dapplink.incentives.v1.GenesisState"></a>

### GenesisState
GenesisState defines the module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#dapplink.incentives.v1.Params) |  | module parameters |
| `incentives` | [Incentive](#dapplink.incentives.v1.Incentive) | repeated | active incentives |
| `gas_meters` | [GasMeter](#dapplink.incentives.v1.GasMeter) | repeated | active Gasmeters |






<a name="dapplink.incentives.v1.Params"></a>

### Params
Params defines the incentives module params


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `enable_incentives` | [bool](#bool) |  | parameter to enable incentives |
| `allocation_limit` | [string](#string) |  | maximum percentage an incentive can allocate per denomination |
| `incentives_epoch_identifier` | [string](#string) |  | identifier for the epochs module hooks |
| `reward_scaler` | [string](#string) |  | scaling factor for capping rewards |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="dapplink/incentives/v1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## dapplink/incentives/v1/query.proto



<a name="dapplink.incentives.v1.QueryAllocationMeterRequest"></a>

### QueryAllocationMeterRequest
QueryAllocationMeterRequest is the request type for the Query/AllocationMeter
RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  | denom is the coin denom to query an allocation meter for. |






<a name="dapplink.incentives.v1.QueryAllocationMeterResponse"></a>

### QueryAllocationMeterResponse
QueryAllocationMeterResponse is the response type for the
Query/AllocationMeter RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `allocation_meter` | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) |  |  |






<a name="dapplink.incentives.v1.QueryAllocationMetersRequest"></a>

### QueryAllocationMetersRequest
QueryAllocationMetersRequest is the request type for the
Query/AllocationMeters RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="dapplink.incentives.v1.QueryAllocationMetersResponse"></a>

### QueryAllocationMetersResponse
QueryAllocationMetersResponse is the response type for the
Query/AllocationMeters RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `allocation_meters` | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |






<a name="dapplink.incentives.v1.QueryGasMeterRequest"></a>

### QueryGasMeterRequest
QueryGasMeterRequest is the request type for the Query/Incentive RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `contract` | [string](#string) |  | contract identifier is the hex contract address of a contract |
| `participant` | [string](#string) |  | participant identifier is the hex address of a user |






<a name="dapplink.incentives.v1.QueryGasMeterResponse"></a>

### QueryGasMeterResponse
QueryGasMeterResponse is the response type for the Query/Incentive RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `gas_meter` | [uint64](#uint64) |  |  |






<a name="dapplink.incentives.v1.QueryGasMetersRequest"></a>

### QueryGasMetersRequest
QueryGasMetersRequest is the request type for the Query/Incentives RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `contract` | [string](#string) |  | contract is the hex contract address of a incentivized smart contract |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="dapplink.incentives.v1.QueryGasMetersResponse"></a>

### QueryGasMetersResponse
QueryGasMetersResponse is the response type for the Query/Incentives RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `gas_meters` | [GasMeter](#dapplink.incentives.v1.GasMeter) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |






<a name="dapplink.incentives.v1.QueryIncentiveRequest"></a>

### QueryIncentiveRequest
QueryIncentiveRequest is the request type for the Query/Incentive RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `contract` | [string](#string) |  | contract identifier is the hex contract address of a contract |






<a name="dapplink.incentives.v1.QueryIncentiveResponse"></a>

### QueryIncentiveResponse
QueryIncentiveResponse is the response type for the Query/Incentive RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `incentive` | [Incentive](#dapplink.incentives.v1.Incentive) |  |  |






<a name="dapplink.incentives.v1.QueryIncentivesRequest"></a>

### QueryIncentivesRequest
QueryIncentivesRequest is the request type for the Query/Incentives RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="dapplink.incentives.v1.QueryIncentivesResponse"></a>

### QueryIncentivesResponse
QueryIncentivesResponse is the response type for the Query/Incentives RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `incentives` | [Incentive](#dapplink.incentives.v1.Incentive) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |






<a name="dapplink.incentives.v1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest is the request type for the Query/Params RPC method.






<a name="dapplink.incentives.v1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse is the response type for the Query/Params RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#dapplink.incentives.v1.Params) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="dapplink.incentives.v1.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Incentives` | [QueryIncentivesRequest](#dapplink.incentives.v1.QueryIncentivesRequest) | [QueryIncentivesResponse](#dapplink.incentives.v1.QueryIncentivesResponse) | Incentives retrieves registered incentives | GET|/dapplink/incentives/v1/incentives|
| `Incentive` | [QueryIncentiveRequest](#dapplink.incentives.v1.QueryIncentiveRequest) | [QueryIncentiveResponse](#dapplink.incentives.v1.QueryIncentiveResponse) | Incentive retrieves a registered incentive | GET|/dapplink/incentives/v1/incentives/{contract}|
| `GasMeters` | [QueryGasMetersRequest](#dapplink.incentives.v1.QueryGasMetersRequest) | [QueryGasMetersResponse](#dapplink.incentives.v1.QueryGasMetersResponse) | GasMeters retrieves active gas meters for a given contract | GET|/dapplink/incentives/v1/gas_meters/{contract}|
| `GasMeter` | [QueryGasMeterRequest](#dapplink.incentives.v1.QueryGasMeterRequest) | [QueryGasMeterResponse](#dapplink.incentives.v1.QueryGasMeterResponse) | GasMeter Retrieves a active gas meter | GET|/dapplink/incentives/v1/gas_meters/{contract}/{participant}|
| `AllocationMeters` | [QueryAllocationMetersRequest](#dapplink.incentives.v1.QueryAllocationMetersRequest) | [QueryAllocationMetersResponse](#dapplink.incentives.v1.QueryAllocationMetersResponse) | AllocationMeters retrieves active allocation meters for a given denomination | GET|/dapplink/incentives/v1/allocation_meters|
| `AllocationMeter` | [QueryAllocationMeterRequest](#dapplink.incentives.v1.QueryAllocationMeterRequest) | [QueryAllocationMeterResponse](#dapplink.incentives.v1.QueryAllocationMeterResponse) | AllocationMeter Retrieves a active gas meter | GET|/dapplink/incentives/v1/allocation_meters/{denom}|
| `Params` | [QueryParamsRequest](#dapplink.incentives.v1.QueryParamsRequest) | [QueryParamsResponse](#dapplink.incentives.v1.QueryParamsResponse) | Params retrieves the incentives module params | GET|/dapplink/incentives/v1/params|

 <!-- end services -->



<a name="dapplink/inflation/v1/inflation.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## dapplink/inflation/v1/inflation.proto



<a name="dapplink.inflation.v1.ExponentialCalculation"></a>

### ExponentialCalculation
ExponentialCalculation holds factors to calculate exponential inflation on
each period. Calculation reference:
periodProvision = exponentialDecay       *  bondingIncentive
f(x)            = (a * (1 - r) ^ x + c)  *  (1 + max_variance - bondedRatio *
(max_variance / bonding_target))


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `a` | [string](#string) |  | initial value |
| `r` | [string](#string) |  | reduction factor |
| `c` | [string](#string) |  | long term inflation |
| `bonding_target` | [string](#string) |  | bonding target |
| `max_variance` | [string](#string) |  | max variance |






<a name="dapplink.inflation.v1.InflationDistribution"></a>

### InflationDistribution
InflationDistribution defines the distribution in which inflation is
allocated through minting on each epoch (staking, incentives, community). It
excludes the team vesting distribution, as this is minted once at genesis.
The initial InflationDistribution can be calculated from the DappLink Token
Model like this:
mintDistribution1 = distribution1 / (1 - teamVestingDistribution)
0.5333333         = 40%           / (1 - 25%)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `staking_rewards` | [string](#string) |  | staking_rewards defines the proportion of the minted minted_denom that is to be allocated as staking rewards |
| `usage_incentives` | [string](#string) |  | usage_incentives defines the proportion of the minted minted_denom that is to be allocated to the incentives module address |
| `community_pool` | [string](#string) |  | community_pool defines the proportion of the minted minted_denom that is to be allocated to the community pool |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="dapplink/inflation/v1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## dapplink/inflation/v1/genesis.proto



<a name="dapplink.inflation.v1.GenesisState"></a>

### GenesisState
GenesisState defines the inflation module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#dapplink.inflation.v1.Params) |  | params defines all the paramaters of the module. |
| `period` | [uint64](#uint64) |  | amount of past periods, based on the epochs per period param |
| `epoch_identifier` | [string](#string) |  | inflation epoch identifier |
| `epochs_per_period` | [int64](#int64) |  | number of epochs after which inflation is recalculated |
| `skipped_epochs` | [uint64](#uint64) |  | number of epochs that have passed while inflation is disabled |






<a name="dapplink.inflation.v1.Params"></a>

### Params
Params holds parameters for the inflation module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `mint_denom` | [string](#string) |  | type of coin to mint |
| `exponential_calculation` | [ExponentialCalculation](#dapplink.inflation.v1.ExponentialCalculation) |  | variables to calculate exponential inflation |
| `inflation_distribution` | [InflationDistribution](#dapplink.inflation.v1.InflationDistribution) |  | inflation distribution of the minted denom |
| `enable_inflation` | [bool](#bool) |  | parameter to enable inflation and halt increasing the skipped_epochs |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="dapplink/inflation/v1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## dapplink/inflation/v1/query.proto



<a name="dapplink.inflation.v1.QueryCirculatingSupplyRequest"></a>

### QueryCirculatingSupplyRequest
QueryCirculatingSupplyRequest is the request type for the
Query/CirculatingSupply RPC method.






<a name="dapplink.inflation.v1.QueryCirculatingSupplyResponse"></a>

### QueryCirculatingSupplyResponse
QueryCirculatingSupplyResponse is the response type for the
Query/CirculatingSupply RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `circulating_supply` | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) |  | total amount of coins in circulation |






<a name="dapplink.inflation.v1.QueryEpochMintProvisionRequest"></a>

### QueryEpochMintProvisionRequest
QueryEpochMintProvisionRequest is the request type for the
Query/EpochMintProvision RPC method.






<a name="dapplink.inflation.v1.QueryEpochMintProvisionResponse"></a>

### QueryEpochMintProvisionResponse
QueryEpochMintProvisionResponse is the response type for the
Query/EpochMintProvision RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `epoch_mint_provision` | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) |  | epoch_mint_provision is the current minting per epoch provision value. |






<a name="dapplink.inflation.v1.QueryInflationRateRequest"></a>

### QueryInflationRateRequest
QueryInflationRateRequest is the request type for the Query/InflationRate RPC
method.






<a name="dapplink.inflation.v1.QueryInflationRateResponse"></a>

### QueryInflationRateResponse
QueryInflationRateResponse is the response type for the Query/InflationRate
RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `inflation_rate` | [string](#string) |  | rate by which the total supply increases within one period |






<a name="dapplink.inflation.v1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest is the request type for the Query/Params RPC method.






<a name="dapplink.inflation.v1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse is the response type for the Query/Params RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#dapplink.inflation.v1.Params) |  | params defines the parameters of the module. |






<a name="dapplink.inflation.v1.QueryPeriodRequest"></a>

### QueryPeriodRequest
QueryPeriodRequest is the request type for the Query/Period RPC method.






<a name="dapplink.inflation.v1.QueryPeriodResponse"></a>

### QueryPeriodResponse
QueryPeriodResponse is the response type for the Query/Period RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `period` | [uint64](#uint64) |  | period is the current minting per epoch provision value. |






<a name="dapplink.inflation.v1.QuerySkippedEpochsRequest"></a>

### QuerySkippedEpochsRequest
QuerySkippedEpochsRequest is the request type for the Query/SkippedEpochs RPC
method.






<a name="dapplink.inflation.v1.QuerySkippedEpochsResponse"></a>

### QuerySkippedEpochsResponse
QuerySkippedEpochsResponse is the response type for the Query/SkippedEpochs
RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `skipped_epochs` | [uint64](#uint64) |  | number of epochs that the inflation module has been disabled. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="dapplink.inflation.v1.Query"></a>

### Query
Query provides defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Period` | [QueryPeriodRequest](#dapplink.inflation.v1.QueryPeriodRequest) | [QueryPeriodResponse](#dapplink.inflation.v1.QueryPeriodResponse) | Period retrieves current period. | GET|/dapplink/inflation/v1/period|
| `EpochMintProvision` | [QueryEpochMintProvisionRequest](#dapplink.inflation.v1.QueryEpochMintProvisionRequest) | [QueryEpochMintProvisionResponse](#dapplink.inflation.v1.QueryEpochMintProvisionResponse) | EpochMintProvision retrieves current minting epoch provision value. | GET|/dapplink/inflation/v1/epoch_mint_provision|
| `SkippedEpochs` | [QuerySkippedEpochsRequest](#dapplink.inflation.v1.QuerySkippedEpochsRequest) | [QuerySkippedEpochsResponse](#dapplink.inflation.v1.QuerySkippedEpochsResponse) | SkippedEpochs retrieves the total number of skipped epochs. | GET|/dapplink/inflation/v1/skipped_epochs|
| `CirculatingSupply` | [QueryCirculatingSupplyRequest](#dapplink.inflation.v1.QueryCirculatingSupplyRequest) | [QueryCirculatingSupplyResponse](#dapplink.inflation.v1.QueryCirculatingSupplyResponse) | CirculatingSupply retrieves the total number of tokens that are in circulation (i.e. excluding unvested tokens). | GET|/dapplink/inflation/v1/circulating_supply|
| `InflationRate` | [QueryInflationRateRequest](#dapplink.inflation.v1.QueryInflationRateRequest) | [QueryInflationRateResponse](#dapplink.inflation.v1.QueryInflationRateResponse) | InflationRate retrieves the inflation rate of the current period. | GET|/dapplink/inflation/v1/inflation_rate|
| `Params` | [QueryParamsRequest](#dapplink.inflation.v1.QueryParamsRequest) | [QueryParamsResponse](#dapplink.inflation.v1.QueryParamsResponse) | Params retrieves the total set of minting parameters. | GET|/dapplink/inflation/v1/params|

 <!-- end services -->



<a name="dapplink/vesting/v1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## dapplink/vesting/v1/query.proto



<a name="dapplink.vesting.v1.QueryBalancesRequest"></a>

### QueryBalancesRequest
QueryBalancesRequest is the request type for the Query/Balances RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  | address of the clawback vesting account |






<a name="dapplink.vesting.v1.QueryBalancesResponse"></a>

### QueryBalancesResponse
QueryBalancesResponse is the response type for the Query/Balances RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `locked` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | current amount of locked tokens |
| `unvested` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | current amount of unvested tokens |
| `vested` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | current amount of vested tokens |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="dapplink.vesting.v1.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Balances` | [QueryBalancesRequest](#dapplink.vesting.v1.QueryBalancesRequest) | [QueryBalancesResponse](#dapplink.vesting.v1.QueryBalancesResponse) | Retrieves the unvested, vested and locked tokens for a vesting account | GET|/dapplink/vesting/v1/balances/{address}|

 <!-- end services -->



<a name="dapplink/vesting/v1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## dapplink/vesting/v1/tx.proto



<a name="dapplink.vesting.v1.MsgClawback"></a>

### MsgClawback
MsgClawback defines a message that removes unvested tokens from a
ClawbackVestingAccount.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `funder_address` | [string](#string) |  | funder_address is the address which funded the account |
| `account_address` | [string](#string) |  | account_address is the address of the ClawbackVestingAccount to claw back from. |
| `dest_address` | [string](#string) |  | dest_address specifies where the clawed-back tokens should be transferred to. If empty, the tokens will be transferred back to the original funder of the account. |






<a name="dapplink.vesting.v1.MsgClawbackResponse"></a>

### MsgClawbackResponse
MsgClawbackResponse defines the MsgClawback response type.






<a name="dapplink.vesting.v1.MsgCreateClawbackVestingAccount"></a>

### MsgCreateClawbackVestingAccount
MsgCreateClawbackVestingAccount defines a message that enables creating a
ClawbackVestingAccount.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `from_address` | [string](#string) |  | from_address specifies the account to provide the funds and sign the clawback request |
| `to_address` | [string](#string) |  | to_address specifies the account to receive the funds |
| `start_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | start_time defines the time at which the vesting period begins |
| `lockup_periods` | [cosmos.vesting.v1beta1.Period](#cosmos.vesting.v1beta1.Period) | repeated | lockup_periods defines the unlocking schedule relative to the start_time |
| `vesting_periods` | [cosmos.vesting.v1beta1.Period](#cosmos.vesting.v1beta1.Period) | repeated | vesting_periods defines thevesting schedule relative to the start_time |
| `merge` | [bool](#bool) |  | merge specifies a the creation mechanism for existing ClawbackVestingAccounts. If true, merge this new grant into an existing ClawbackVestingAccount, or create it if it does not exist. If false, creates a new account. New grants to an existing account must be from the same from_address. |






<a name="dapplink.vesting.v1.MsgCreateClawbackVestingAccountResponse"></a>

### MsgCreateClawbackVestingAccountResponse
MsgCreateClawbackVestingAccountResponse defines the
MsgCreateClawbackVestingAccount response type.





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="dapplink.vesting.v1.Msg"></a>

### Msg
Msg defines the vesting Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `CreateClawbackVestingAccount` | [MsgCreateClawbackVestingAccount](#dapplink.vesting.v1.MsgCreateClawbackVestingAccount) | [MsgCreateClawbackVestingAccountResponse](#dapplink.vesting.v1.MsgCreateClawbackVestingAccountResponse) | CreateClawbackVestingAccount creats a vesting account that is subject to clawback and the configuration of vesting and lockup schedules. | GET|/dapplink/vesting/v1/tx/create_clawback_vesting_account|
| `Clawback` | [MsgClawback](#dapplink.vesting.v1.MsgClawback) | [MsgClawbackResponse](#dapplink.vesting.v1.MsgClawbackResponse) | Clawback removes the unvested tokens from a ClawbackVestingAccount. | GET|/dapplink/vesting/v1/tx/clawback|

 <!-- end services -->



<a name="dapplink/vesting/v1/vesting.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## dapplink/vesting/v1/vesting.proto



<a name="dapplink.vesting.v1.ClawbackVestingAccount"></a>

### ClawbackVestingAccount
ClawbackVestingAccount implements the VestingAccount interface. It provides
an account that can hold contributions subject to "lockup" (like a
PeriodicVestingAccount), or vesting which is subject to clawback
of unvested tokens, or a combination (tokens vest, but are still locked).


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `base_vesting_account` | [cosmos.vesting.v1beta1.BaseVestingAccount](#cosmos.vesting.v1beta1.BaseVestingAccount) |  | base_vesting_account implements the VestingAccount interface. It contains all the necessary fields needed for any vesting account implementation |
| `funder_address` | [string](#string) |  | funder_address specifies the account which can perform clawback |
| `start_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | start_time defines the time at which the vesting period begins |
| `lockup_periods` | [cosmos.vesting.v1beta1.Period](#cosmos.vesting.v1beta1.Period) | repeated | lockup_periods defines the unlocking schedule relative to the start_time |
| `vesting_periods` | [cosmos.vesting.v1beta1.Period](#cosmos.vesting.v1beta1.Period) | repeated | vesting_periods defines the vesting schedule relative to the start_time |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

