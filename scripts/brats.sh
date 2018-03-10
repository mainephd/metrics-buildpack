#!/usr/bin/env bash
set -euo pipefail

export ROOT=$(dirname $(greadlink -f ${BASH_SOURCE%/*}))
$ROOT/scripts/install_tools.sh

GINKGO_NODES=${GINKGO_NODES:-3}
GINKGO_ATTEMPTS=${GINKGO_ATTEMPTS:-1}

cd $ROOT/src/binary/brats
ginkgo -r --flakeAttempts=$GINKGO_ATTEMPTS -nodes $GINKGO_NODES --slowSpecThreshold=60
