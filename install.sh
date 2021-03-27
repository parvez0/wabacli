#!/usr/bin/env sh
#
# Run With:
#   bash <(curl -Ss https://raw.githubusercontent.com/parvez0/wabacli/main/install.sh)
#
# Other options:
#   --installation_path                     where to install executable default to /usr/bin
#   --update                                update the installed version to latest release
#   --version                               install a specific version (eg. v1.2.3)
#   --http_agent                            which tool utility to use curl or wget, by defaults it will search for available tool
#
# shellcheck disable=SC2039,SC2059,SC2086,SC2046,SC2006
# =================== loggers ==================

RETRIES=1
MAX_RETRIES=3

fatal() {
  printf >&2 "${TPUT_BGRED}${TPUT_WHITE}${TPUT_BOLD} ABORTED ${TPUT_RESET} ${*} \n\n" && exit
}

success() {
  printf >&2 "${TPUT_BGGREEN}${TPUT_WHITE}${TPUT_BOLD} OK ${TPUT_RESET} \n\n"
}

info(){
  printf >&2 "${*}\n"
}

warning() {
  printf >&2 "${TPUT_BGRED}${TPUT_WHITE}${TPUT_BOLD} WARNING ${TPUT_RESET} $1 \n"
  if [ "${INTERACTIVE}" = "0" ]; then
    fatal "Stopping due to non-interactive mode. Fix the issue or retry installation in an interactive mode."
  else
    if [ $RETRIES -gt $MAX_RETRIES ]; then
      fatal "Too many retries"
    fi
    read -r -p "$2"
    RETRIES=$(("$RETRIES"+1))
    verify
  fi
}

AGENT=""
OPTIONS=""
INTERACTIVE=1
UPDATE=0
VERSION="latest"
ARCH=""

verify() {
  info "verifying dependencies required for setup"
  if [ `command -v curl` ]; then
    AGENT="curl"
    OPTIONS="-Ss"
  elif [ `command -v wget` ]; then
    AGENT="wget"
    OPTIONS="-q -O -"
  else
    warning "curl or wget not found, which is required to download the setup files" "Install curl or wget and press enter >"
  fi
}

latest_release() {
  URL="https://raw.githubusercontent.com/parvez0/wabacli/main/.release"
  if [ $VERSION == "latest" ]; then
    cmd="$AGENT $OPTIONS $URL"
    info "fetching the latest release info from $URL"
    release="$(eval "$cmd")"
    VERSION="$(echo "$release" | cut -d "=" -f2)"
    if [[ "$VERSION" = v* ]]; then
      info "latest release version: $VERSION"
    else
      fatal "failed to fetch latest release: $release"
    fi
  fi
}

system_info() {
  SYSTEM="$(uname -s 2> /dev/null || uname -v)"
  OS="$(uname -o 2> /dev/null || uname -rs)"
  MACHINE="$(uname -m 2> /dev/null)"



}

download_tarball(){
  URL=""
}

verify
latest_release