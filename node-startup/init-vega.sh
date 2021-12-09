#!/bin/bash
HERE="$(dirname "$(realpath "$0")")"

docker run -t -d --rm -v "${HERE}"/vegahome:/vegahome --entrypoint=/bin/bash --name core docker.pkg.github.com/vegaprotocol/vega/vega:local

# Init core and faucet
docker exec core vega init -f --home=/vegahome --nodewallet-passphrase-file=/vegahome/passphrase-file --output=json

# Make the nodewallets
docker exec core vega nodewallet generate --passphrase-file=/vegahome/passphrase-file --wallet-passphrase-file=/vegahome/passphrase-file --chain=vega --home=/vegahome  --output=json
docker exec core vega nodewallet generate --passphrase-file=/vegahome/passphrase-file --wallet-passphrase-file=/vegahome/passphrase-file --chain=ethereum --home=/vegahome --output=json

# copy good config and do faucet
cp "${HERE}"/vegahome/config.toml "${HERE}"/vegahome/config/node
docker exec core vega faucet init -f --home=/vegahome --update-in-place --passphrase-file=/vegahome/passphrase-file --output=json

# genesis file creation and update with appstate
docker exec core vega tm init --home=/vegahome
docker exec core vega genesis generate --home=/vegahome --passphrase-file=/vegahome/passphrase-file
docker exec core vega genesis update --tm-root=/vegahome --home=/vegahome --passphrase-file=/vegahome/passphrase-file

docker kill core

# Init datanode
docker run --rm -v "${HERE}"/vegahome:/vegahome docker.pkg.github.com/vegaprotocol/data-node/data-node:develop init -f --home=/vegahome 

# Init vegawallet
docker run --rm -v "${HERE}"/vegahome:/vegahome vegaprotocol/vegawallet init -f --home=/vegahome --output=json