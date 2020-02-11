#!/bin/sh

set -e

generate_structs() {
    pkg_name=$1
    prefix=$2

    echo "_preamble(. \"${pkg_name}\", ${prefix})" \
        | m4 ./generate_test_template.m4 -

    gopath=${GOPATH:-${HOME}/go}

    ls "${gopath}/src/${pkg_name}"/*.go \
        | grep -v -e _test.go -e .generated.go \
        | xargs grep -h -e 'type [A-Z].* struct' \
        | sort  \
        | awk '{printf "_test_type(%s)\n", $2;}' \
        | sed "s|)|, ${prefix})|g" \
        | m4 ./generate_test_template.m4 -
}

#generate_structs github.com/hashicorp/nomad/nomad/structs  NomadNomad >  nomad_nomad_structs_generated_test.go
#generate_structs github.com/hashicorp/nomad/client/structs NomadClient > nomad_client_structs_generated_test.go
generate_structs github.com/hashicorp/consul/agent/structs ConsulAgent > consul_agent_structs_generated_test.go
