# QIO

Command Line Interface for RainbowQ

Currently uses Cobra as a framework to load rainbow.toml, the default data format for the RainbowQ Almanac system.

## RainbowQ

The idea behind this app is a simple method of gathering currated resources of data.

It is meant to be system agnostic, available as a Go application on many varieties of OS.

The power in RainbowQ is what data you put in it.

    * Shared bookmark list
    * Common commands for performing an action
    * Pointers to howto pages
    * Reading lists
    * Endpoints for monitoring

This data lives by default in `rainbow.toml`. Future versions plan to allow for remote storage (e.g. Consul). The data is not meant to contain secrets, instead it would contain pointers to where secrets for a particular Almanac are kept.

## Almanac System

RainbowQ is a new way to think about data.

It uses an Almanac: a collection of data around a particular topic that only exists within the context of its connections.

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
      * The _value_ is the pointer for the concept.

For instance, the `[company.corp]` *Almanac* contains the following *Plugs*:

    - editor = 'maroda@rainbowq.io'
    - snailmail = '1234 ILoveYou Dr., Fullerton CA, 92831'
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

## Display Data

## Write Data

