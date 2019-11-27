#!/usr/bin/env bash

# Pre-requisites: Git, Nix

set -e

GIT_ROOT=$(cd "${BASH_SOURCE%/*}" && git rev-parse --show-toplevel)

# NOTE: To use a local Nimbus repository, uncomment and edit the following line
#nimbus_dir=~/src/github.com/status-im/nimbus

target_dir="${GIT_ROOT}/vendor/github.com/status-im/status-go/eth-node/bridge/nimbus"

if [ -z "$nimbus_dir" ]; then
  # The git ref of Nimbus to fetch and build. This should represent a commit SHA or a tag, for reproducible builds
  nimbus_ref='master' # TODO: Use a tag once

  nimbus_src='https://github.com/status-im/nimbus/'
  nimbus_dir="${GIT_ROOT}/vendor/github.com/status-im/nimbus"

  trap "rm -rf $nimbus_dir" ERR INT QUIT

  # Clone nimbus repo into vendor directory, if necessary
  if [ -d "$nimbus_dir" ]; then
    cd $nimbus_dir && git reset --hard $nimbus_ref; cd -
  else
    # List fetched from vendorDeps array in https://github.com/status-im/nimbus/blob/master/nix/nimbus-wrappers.nix#L9-L12
    vendor_paths=( nim-chronicles nim-faststreams nim-json-serialization nim-chronos nim-eth nim-json nim-metrics nim-secp256k1 nim-serialization nim-stew nim-stint nimcrypto )
    vendor_path_opts="${vendor_paths[@]/#/--recurse-submodules=vendor/}"
    git clone $nimbus_src --progress ${vendor_path_opts} --depth 1 -j8 -b $nimbus_ref $nimbus_dir
  fi
fi

# Build Nimbus wrappers and copy them into the Nimbus bridge in status-eth-node
build_dir=$(cd $nimbus_dir && nix-build --pure --no-out-link -A wrappers)
# Ideally we'd use the static version of the Nimbus library (.a),
# however that causes link errors due to duplicate symbols:
# ${target_dir}/libnimbus.a(secp256k1.c.o): In function `secp256k1_context_create':
# (.text+0xca80): multiple definition of `secp256k1_context_create'
# /tmp/go-link-476687730/000014.o:${GIT_ROOT}/vendor/github.com/ethereum/go-ethereum/crypto/secp256k1/./libsecp256k1/src/secp256k1.c:56: first defined here
rm -f ${target_dir}/libnimbus.*
mkdir -p ${target_dir}
cp -f ${build_dir}/include/* ${build_dir}/lib/libnimbus.a \
      ${target_dir}/
chmod +w ${target_dir}/libnimbus.{a,h}
objcopy -L secp256k1_context_create -L secp256k1_context_clone -L secp256k1_context_destroy -L secp256k1_context_set_illegal_callback -L secp256k1_context_set_error_callback -L secp256k1_ec_pubkey_parse -L secp256k1_ec_pubkey_serialize -L secp256k1_ecdsa_signature_parse_der -L secp256k1_ecdsa_signature_parse_compact -L secp256k1_ecdsa_signature_serialize_der -L secp256k1_ecdsa_signature_serialize_compact -L secp256k1_ecdsa_signature_normalize -L secp256k1_ecdsa_verify -L secp256k1_ecdsa_sign -L secp256k1_ec_seckey_verify -L secp256k1_ec_pubkey_create -L secp256k1_ec_privkey_tweak_add -L secp256k1_ec_pubkey_tweak_add -L secp256k1_ec_privkey_tweak_mul -L secp256k1_ec_pubkey_tweak_mul -L secp256k1_context_randomize -L secp256k1_ec_pubkey_combine -L secp256k1_ecdsa_recoverable_signature_parse_compact -L secp256k1_ecdsa_recoverable_signature_serialize_compact -L secp256k1_ecdsa_recoverable_signature_convert -L secp256k1_ecdsa_sign_recoverable -L secp256k1_ecdsa_recover -L secp256k1_nonce_function_default -L secp256k1_nonce_function_rfc6979 -L CURVE_B \
  ${target_dir}/libnimbus.a
