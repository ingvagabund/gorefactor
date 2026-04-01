header() {
	local title=" 🔆🔆🔆  $*  🔆🔆🔆 "

	local len=40
	if [[ ${#title} -gt $len ]]; then
		len=${#title}
	fi

	echo -e "\n\n  \033[1m${title}\033[0m"
	echo -n "━━━━━"
	printf '━%.0s' $(seq "$len")
	echo "━━━━━━━"

}

info() {
	echo -e " 🔔 $*" >&2
}

err() {
	echo -e " 😱 $*" >&2
}

warn() {
	echo -e "   $*" >&2
}

ok() {
	echo -e "   ✅ $*" >&2
}

skip() {
	echo -e " 🙈 SKIP: $*" >&2
}

fail() {
	echo -e " ❌ FAIL: $*" >&2
}

info_run() {
	echo -e "      $*\n" >&2
}

run() {
	echo -e " ❯ $*\n" >&2
	"$@"
}

die() {
	echo -e "\n ✋ $* "
	echo -e "──────────────────── ⛔️⛔️⛔️ ────────────────────────\n"
	exit 1
}

line() {
	local len="$1"
	local style="${2:-thin}"
	shift

	local ch='─'
	[[ "$style" == 'heavy' ]] && ch="━"

	printf "$ch%.0s" $(seq "$len") >&2
	echo
}
