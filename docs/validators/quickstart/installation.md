<!--
order: 1
-->

# Installation

Build and install the dapplink binaries from source or using Docker. {synopsis}

## Pre-requisites

- [Install Go 1.18.5+](https://golang.org/dl/) {prereq}
- [Install jq](https://stedolan.github.io/jq/download/) {prereq}

## Install Go

::: warning
dapplink is built using [Go](https://golang.org/dl/) version `1.18+`
:::

```bash
go version
```

:::tip
If the `dapplinkd: command not found` error message is returned, confirm that your [`GOPATH`](https://golang.org/doc/gopath_code#GOPATH) is correctly configured by running the following command:

```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

:::

## Install Binaries

::: tip
The latest {{ $themeConfig.project.name }} [version](https://github.com/dapplink/dapplink/releases) is `{{ $themeConfig.project.binary }} {{ $themeConfig.project.latest_version }}`
:::

### GitHub

Clone and build {{ $themeConfig.project.name }} using `git`:

```bash
git clone https://github.com/dapplink/dapplink.git
cd dapplink
make install
```

Check that the `{{ $themeConfig.project.binary }}` binaries have been successfully installed:

```bash
dapplinkd version
```

### Docker

You can build {{ $themeConfig.project.name }} using Docker by running:

```bash
make build-docker
```

The command above will create a docker container: `tharsishq/dapplink:latest`. Now you can run `dapplinkd` in the container.

```bash
docker run -it -p 26657:26657 -p 26656:26656 -v ~/.dapplinkd/:/root/.dapplinkd tharsishq/dapplink:latest dapplinkd version

# To initialize
# docker run -it -p 26657:26657 -p 26656:26656 -v ~/.dapplinkd/:/root/.dapplinkd tharsishq/dapplink:latest dapplinkd init test-chain --chain-id test_9000-2

# To run
# docker run -it -p 26657:26657 -p 26656:26656 -v ~/.dapplinkd/:/root/.dapplinkd tharsishq/dapplink:latest dapplinkd start
```

### Releases

You can also download a specific release available on the {{ $themeConfig.project.name }} [repository](https://github.com/dapplink/dapplink/releases) or via command line:

```bash
go install github.com/dapplink/dapplink@latest
```
