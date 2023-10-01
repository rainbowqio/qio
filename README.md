# QIO

[![Build Qio CLI and push to GCR](https://github.com/rainbowqio/qio/actions/workflows/google.yaml/badge.svg)](https://github.com/rainbowqio/qio/actions/workflows/google.yaml)

Command Line Interface for RainbowQ

Uses Cobra / Viper as a framework to load _rainbow.toml_, the default data format for the RainbowQ Almanac system.

_(builds might be failing while the GCR endpoint is being worked out)_

## RainbowQ

The power in RainbowQ is what data you put in it. It's a configuration for your sociotechnical company or project, a way to help quickly link humans at the sharp end with common threads of knowledge on the command line. Yep it's damn specific and I love using it, like one of those specialized tools that is finely tooled that helps fill a specific void. Things you need on the command line like:

* A shared bookmark list
* Common commands for performing an action
* Pointers to howto pages
* Reading lists
* Endpoints for monitoring
* Account numbers or unsecure tokens
  The data is not meant to contain secrets, instead it would contain pointers to where secrets for a particular thing are kept.

This data lives by default in `rainbow.toml`. Future versions plan to allow for remote storage (e.g. Consul) or queries to other types of data endpoints (e.g. Google API) to build Plugs.

Plugs are the basic building block of the Almanac.

## Almanac System

RainbowQ organizes your data around the concept of an Almanac: a collection of data around a particular topic. This can become nested and complex if desired, or remain simple.

* The Almanac structure in RainbowQ is analogous to a database table.
* The native Almanac format is TOML.
* Almanac names can be simple or complex:
  * [meta] or [meta.last] or [endpoint.prod]
* Keys in Almanacs are unique for that Almanac:
  * [meta.last].editor is not the same as [customer.meta].editor
  * Keys should remain simple and on lines by themselves.
    That's it. Almanacs can be simple or complex according to name and contain an arbitrary set of key / value pairs, known as Plugs.

## Plugs

The basic unit of knowledge in a RainbowQ database is the *Plug*.

* Plugs are technically _key / value pairs._:
  * The _key_ is a concept.
  * The _value_ is the *pointer* for the concept.
    For instance, the `[company.corp]` *Almanac* contains the following *Plugs*:

- editor = 'maroda@rainbowq.io'
- snailmail = '1234 ILoveYou Dr., Hatachooe CA, 91111'
- phone = '+1.714.123.4567'
- account = 'RQ-BLS-0987'
  There is an `editor` Plug that points to an email address.
  There is a `snailmail` Plug that points to a mailing address.
  A `phone` Plug points to a US phone number.
  The `account` Plug points to this customer's account number.

RainbowQ does not portent to be the source of truth, but it points to sources of truth.

That is, data that ends up as part of a Plug must be programmatically useful.
It either contains authoritative data from that data's source of truth,
or it contains the pointer to where that data can be found.

There is no such thing as a default Plug value. The Plug exists because it is a pointer to a known piece of data.
That pointer can be *Hard* (a value, like a phone number) or *Soft* (a way to get the value required).

A functional piece of data, like a shell command, is considered a Soft Plug.
It won't be the result of a command, but the way to get a result.

## Usage

The `qio` command itself has no data. If run without an argument it will create its default database in `~/.config/qio/rainbow.toml`, but will not overwrite one if it already exists there. Below is the help to give some idea of how it works.

```zsh
>>> qio --help
RainbowQ ::: all your knowledgebase are belong to us

Usage:
  qio [command]

Examples:

	::: Display a Plug within an Almanac:
	$ qio ask <almanac> <plug>

	::: List what QIO knows:
	$ qio list

	::: QIO has randomizers! Toss a coin for HEADS or TAILS, get a token, or get a URL-encoded base64 token:
	$ qio coin
	$ qio coin -t
	$ qio coin -b

	::: Install shell completion (see 'qio completion --help' for more):
	$ qio completion zsh > "${fpath[1]}/_qio"

	::: Export known Rainbow to a local TOML file:
	$ qio export


Available Commands:
  ask         Ask QIO a question with: <almanac> <plug>
  coin        Toss a coin. Use the '-t|--token' flag for a randomized 32 character token.
  completion  Generate completion script for QIO
  export      Export QIO knowledge
  help        Help about any command
  list        List what QIO knows

Flags:
      --config string   config file (default is $HOME/.config/qio/rainbow.toml)
  -h, --help            help for qio
      --version         print the version

Use "qio [command] --help" for more information about a command.
```

## Write Data

Currently `rainbow.toml` is currated by the owner in its own repository.

## Download

Example setup of a Rainbow data source that lives on an internal GCP Storage Bucket along with a binary uploaded from this repo (this example is for Zsh, see qio completion --help for others):

1. `gsutil cp gs://rainbowq/qio-macos /usr/local/bin/qio`
2. `chmod 755 /usr/local/bin/qio`
3. `qio completion zsh > "${fpath[1]}/_qio"`
4. `gsutil cp gs://rainbowq/rainbow.toml ~/.config/qio/rainbow.toml` (you may need to create this directory)
