#!/usr/bin/env bash

set -eu -o pipefail

# constants
PROJECT_ROOT="$(git rev-parse --show-toplevel)"

declare -r PROJECT_ROOT
declare -r LOCAL_BIN="$PROJECT_ROOT/tmp/bin"

# versions
declare -r GOLANGCI_LINT_VERSION=${GOLANGCI_LINT_VERSION:-v1.64.5}

source "$PROJECT_ROOT/hack/utils.bash"

go_install() {
	local pkg="$1"
	local version="$2"
	shift 2

	info "installing $pkg version: $version"

	GOBIN=$LOCAL_BIN go install "$pkg@$version" || {
		fail "failed to install $pkg - $version"
		return 1
	}
	ok "$pkg - $version was installed successfully"
	return 0
}

validate_version() {
	local cmd="$1"
	local version_arg="$2"
	local version_regex="$3"
	shift 3

	command -v "$cmd" >/dev/null 2>&1 || return 1

	[[ "$(eval "$cmd $version_arg" | grep -o "$version_regex")" =~ $version_regex ]] || {
		return 1
	}

	ok "$cmd matching $version_regex already installed"
}

version_golangci-lint() {
	golangci-lint --version
}

install_golangci-lint() {
	local version_regex="$GOLANGCI_LINT_VERSION"
	validate_version golangci-lint --version "$version_regex" && return 0
	go_install github.com/golangci/golangci-lint/cmd/golangci-lint "$GOLANGCI_LINT_VERSION"
}

install_all() {
	info "installing all tools ..."
	local ret=0
	for tool in $(declare -F | cut -f3 -d ' ' | grep install_ | grep -v 'install_all'); do
		"$tool" || ret=1
	done
	return $ret
}

version_all() {
	header "Versions"
	for version_tool in $(declare -F | cut -f3 -d ' ' | grep version_ | grep -v 'version_all'); do
		local tool="${version_tool#version_}"
		local location=""
		location="$(command -v "$tool")"
		info "$tool -> $location"
		"$version_tool"
		echo
	done
	line "50"
}

main() {
	local op="${1:-all}"
	shift || true

	mkdir -p "$LOCAL_BIN"
	export PATH="$LOCAL_BIN:$PATH"

	# NOTE: skip installation if invocation is tools.sh version
	if [[ "$op" == "version" ]]; then
		version_all
		return $?
	fi

	install_"$op"
	version_"$op"
}

main "$@"
