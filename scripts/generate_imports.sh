#!/usr/bin/env bash

set -euo pipefail

package="$1"
target_file="${2:-}"

if [ -n "$target_file" ]; then
  exec >"$target_file"
fi

create_import_path() {
  for path in "$@"; do
    printf "\t_ \"github.com/hoshsadiq/big-fat-converter/codec/%s\"\n" "$(basename "$path")"
  done
}

export -f create_import_path

printf "package %s\n" "$package"
printf "\n"
printf "import (\n"
find ../../codec -mindepth 1 -maxdepth 1 -type d -print0 | sort -z |
  xargs -0 bash -c 'create_import_path "$@"' _
printf ")\n"
