<img alt="" src="https://github.com/parvez0/wabacli/raw/main/assets/whatsapp_logo.png" width="40%"/>

[![GoDoc](https://godoc.org/github.com/sirupsen/logrus?status.svg)](https://pkg.go.dev/github.com/spf13/cobra) [![Build](https://badgen.net/badge/build/sucess/green?icon=github)](https://pkg.go.dev/github.com/spf13/cobra) [![Release](https://img.shields.io/badge/release-v0.0.19-blue)](https://github.com/parvez0/wabacli/releases) [![Dependency](https://img.shields.io/badge/dependency-cobra-blueviolet)](https://pkg.go.dev/github.com/spf13/cobra)

Wabacli is a command line utility to interact with whatsapp business account APIs. It provides you a way to store context multiple accounts information which can makes development easier. You can read more about
whatsapp business account <a href="https://developers.facebook.com/docs/whatsapp/overview" target="_blank">here</a>.

You can install Wabacli on most Linux distributions (Ubuntu, Debian, CentOS, and more) and many other operating systems (FreeBSD, macOS)

Wabacli was designed for whatsapp BSP's who requires maintaining of multiple accounts.

## Menu

- [Features](#features)
- [Prerequisite](#prerequisite)  
- [Installation](#installation)
- [Documentation](#documentation)
- [Contribute](#contribute)
- [License](#license)

## Features

![Demo](./assets/whatsapp.gif)

Here's what you can expect from Wabacli:

- **Multiple Accounts:** It maintains the context of more than one account and switch between them easily.
- **Simplicity:** It's really easy to send different types of messages with a single command.
- **Zero Configuration:** It does not require any configuration to set up, you can install it and get started.

## Prerequisite

You will need a whatsapp business account set up to register your business number with facebook. After registering
you can use the provided certificate to activate your whatsapp account for complete details on how to set it up you
can follow the Get-Started guide <a href="https://developers.facebook.com/docs/whatsapp/getting-started" target="_blank">here</a>.

## Installation

To install Wabacli you can our one-line installation scripts which will verify the dependencies on your system, fetches the latest
version and installs it on your system.

```bash
bash <(curl -Ss https://raw.githubusercontent.com/parvez0/wabacli/main/install.sh)
```