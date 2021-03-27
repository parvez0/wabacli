#!/usr/bin/env sh
#
# Run With:
#   bash <(curl -Ss https://raw.githubusercontent.com/parvez0/wabacli/main/install.sh)
#
# Other options:
#   --installation_path                     where to install executable default to /usr/bin
#   --version                               install a specific version (eg. v1.2.3)
#
# shellcheck disable=SC2039,SC2059,SC2086,SC2046,SC2006,SC2120
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
VERSION="latest"
ARCH=""
INSTTALLATION_PATH="/usr/bin/"

# verifying dependencies
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

# fetching the system info for architecture details
system_info() {
  SYSTEM="$(uname -s 2> /dev/null || uname -v)"
  OS="$(uname -o 2> /dev/null || uname -rs)"
  MACHINE="$(uname -m 2> /dev/null)"
  ARCH=$(echo "${SYSTEM}_${MACHINE}" | tr '[:upper:]' '[:lower:]')

  info "SYSTEM       :   $SYSTEM"
  info "MACHINE      :   $MACHINE"
  info "ARCHITECTURE :   $ARCH"
}

# get the latest release version from the server
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

# download the latest released binaries for installation
download_tarball(){
  URL="https://github.com/parvez0/wabacli/releases/download/$VERSION/wabacli_$VERSION_$ARCH.tar.gz"
  info "downloading latest release binaries from $URL"
  cmd="$AGENT -O wabacli_$VERSION_$ARCH.tar.gz $URL"
  eval "$cmd"
  echo "extracting tar ball to wabacli_$VERSION_$ARCH"
  tar -C "wabacli_$VERSION_$ARCH" -xf "wabacli_$VERSION_$ARCH.tar.gz"
}

# copy the binary to installation path
install(){
  echo "installing wabacli"
  sudo cp "/tmp/wabacli_$VERSION_$ARCH/wabacli" $INSTTALLATION_PATH
}

#verify
#latest_release
#system_info

# ----------------------------------------------------------------------------------
validate_input() {
    if [[ $2 == "" || $2 == --* ]]; then
        fatal "$1 requires input"
    fi
}

while [ -n "${1}" ]; do
    arg="$1"
    case "${1}" in
    "--installation_path")
        shift 1
        INSTTALLATION_PATH="${1}"
        validate_input "$arg" "$1"
        if [ ! -d "$1" ]; then
          fatal "specified directory \"$1\" does not exits"
        fi
        info "setting up installation path to $INSTTALLATION_PATH"
        ;;
    "--version")
        shift 1
        VERSION="${1}"
        validate_input "$arg" "$1"
        info "installing application with version \"$VERSION\""
        ;;
    esac
    shift 1
done

# ------------------- #
verify
system_info
latest_release
download_tarball
install