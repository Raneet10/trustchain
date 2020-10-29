# trustchain



TrustChain - the network where people will give promises to each other and they will be recorded on the blockchain. Once a promise has been fulfilled, the predefined award of TRUST tokens for it will be given to the promise keeper. In case it wasn't fulfilled, the reward of TRUST tokens will be returned to the disillusioned side, plus a bonus amount, taken from the promise failer.

Built using Cosmos SDK and Tendermint and generated with [Starport](https://github.com/tendermint/starport).

Work in progress!

## Get started

```
starport serve
```

`serve` command installs dependencies, initializes and runs the application.

## Configure

Initialization parameters of your app are stored in `config.yml`.

### `accounts`

A list of user accounts created during genesis of your application.

| Key   | Required | Type            | Description                                       |
| ----- | -------- | --------------- | ------------------------------------------------- |
| name  | Y        | String          | Local name of the key pair                        |
| coins | Y        | List of Strings | Initial coins with denominations (e.g. "10trust") |

## Learn more

- [Starport](https://github.com/tendermint/starport)
- [Cosmos SDK documentation](https://docs.cosmos.network)
- [Cosmos Tutorials](https://tutorials.cosmos.network)
- [Channel on Discord](https://discord.gg/W8trcGV)
