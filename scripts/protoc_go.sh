#!/bin/bash

ROOT=$(cd `dirname $0`&&pwd)
OUT_PUT_dir=""

function show_help() {
    echo -e " Usage:\n" \
    "$0 [options] go_src_dir\n" \
    "Options: \n" \
    "    -f input_file_name : only generate single proto file's GO output\n" \
    "    -h                 : show this help message\n" \
    "    -p package_name    : output package (default: ${default_go_package_base})\n" >&2
}

while [[ -n "$1" ]]; do
    case "$1" in
    -f)
        shift
        input_file_name=$1
        ;;
    -p)
        shift
        go_package_base=$1
        ;;
    -h)
        show_help
        exit 0
        ;;
    *)
        break
        ;;
    esac
    shift
done

function generate_go() {
    local file_name=$1
    local pkg_name=$(grep 'package' --max-count=1 ${file_name} | awk '{print $2}' | tr -d \; )
    local path_name=${pkg_name/./\/}
    local base_name=$(basename ${path_name})
    #echo "file name: ${file_name}"
    #echo "package name: ${pkg_name}"
    #echo "path name: ${path_name}"
    #echo "base name: ${base_name}"
    local go_output_path="${go_output_base}/${path_name}"
    mkdir -p ${go_output_path}
    echo "${file_name} => ${go_output_path}"
    go_out_options="import_path=${base_name}"
    go_out_options="${go_out_options},Mucloud.proto=${go_package_base}/ucloud"
    go_out_options="${go_out_options},Mudb.56000.57000.proto=${go_package_base}/ucloud/udb"
    go_out_options="${go_out_options},Muns.55000.56000.proto=${go_package_base}/ucloud/uns"
    protoc "--go_out=${go_out_options}:${go_output_path}" -I ${mydir}/proto ${file_name}
}

if [[ -z ${input_file_name} ]]; then
    for file_name in ${mydir}/proto/*; do
        generate_go "${file_name}"
    done
else
    generate_go "${input_file_name}"
fi