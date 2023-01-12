[![Moov Banner Logo](https://user-images.githubusercontent.com/20115216/104214617-885b3c80-53ec-11eb-8ce0-9fc745fb5bfc.png)](https://github.com/moov-io)

<p align="center">
  <a href="https://github.com/moov-io/rtp20022/tree/master/docs">Project Documentation</a>
  ·
  <a href="https://moov-io.github.io/rtp20022/api/">API Endpoints</a>
  ·
  <a href="https://slack.moov.io/">Community</a>
  ·
  <a href="https://moov.io/blog/">Blog</a>
  <br>
  <br>
</p>

[![GoDoc](https://godoc.org/github.com/moov-io/rtp20022?status.svg)](https://godoc.org/github.com/moov-io/rtp20022)
[![Build Status](https://github.com/moov-io/rtp20022/workflows/Go/badge.svg)](https://github.com/moov-io/rtp20022/actions)
[![Coverage Status](https://codecov.io/gh/moov-io/rtp20022/branch/master/graph/badge.svg)](https://codecov.io/gh/moov-io/rtp20022)
[![Go Report Card](https://goreportcard.com/badge/github.com/moov-io/rtp20022)](https://goreportcard.com/report/github.com/moov-io/rtp20022)
[![Repo Size](https://img.shields.io/github/languages/code-size/moov-io/rtp20022?label=project%20size)](https://github.com/moov-io/rtp20022)
[![Apache 2 License](https://img.shields.io/badge/license-Apache2-blue.svg)](https://raw.githubusercontent.com/moov-io/rtp20022/master/LICENSE)
[![Slack Channel](https://slack.moov.io/badge.svg?bg=e01563&fgColor=fffff)](https://slack.moov.io/)
[![Docker Pulls](https://img.shields.io/docker/pulls/moov/rtp20022)](https://hub.docker.com/r/moov/rtp20022)
[![GitHub Stars](https://img.shields.io/github/stars/moov-io/rtp20022)](https://github.com/moov-io/rtp20022)
[![Twitter](https://img.shields.io/twitter/follow/moov?style=social)](https://twitter.com/moov?lang=en)

## moov-io/rtp20022

Moov's mission is to give developers an easy way to create and integrate bank processing into their own software products. Our open source projects are each focused on solving a single responsibility in financial services and designed around performance, scalability, and ease of use.

This repository contains a subset of ISO 20022 messages for RTP payments. ISO 20022 is a standard for electronic data interchange between financial institutions. It describes a metadata repository containing descriptions of messages and business processes, and a maintenance process for the repository content. The standard covers financial information transferred between financial institutions that includes payment transactions, securities trading and settlement information, credit and debit card transactions, and other financial information.

## Project Status

Go code is [generated with `xsdgen`](https://pkg.go.dev/aqwari.net/xml/cmd/xsdgen) inside the `gen/` folder. Please star the project if you are interested in its progress. Feedback on this early version of ISO 20022 is appreciated and vital to its success. Please let us know if you encounter any bugs/unclear documentation or have feature suggestions by opening up an issue. Thanks!

## Usage

### Go Library

This project offers Go structures which can read and write XML encoded ISO 20022 messages. Go has several [known limitations around XML namespace prefixes](https://github.com/golang/go/issues/13400) which are required for RTP messages. Encoding messages can include namespace prefixes (`xml:"ns:Name"`) but decoding messages cannot (`xml:"Name"`).

## Getting Help

 channel | info
 ------- | -------
[Project Documentation](https://github.com/moov-io/rtp20022/tree/master/docs) | Our project documentation available online.
Twitter [@moov](https://twitter.com/moov) | You can follow Moov.io's Twitter feed to get updates on our project(s). You can also tweet us questions or just share blogs or stories.
[GitHub Issue](https://github.com/moov-io/rtp20022/issues/new) | If you are able to reproduce a problem please open a GitHub Issue under the specific project that caused the error.
[moov-io slack](https://slack.moov.io/) | Join our slack channel (`#rtp20022`) to have an interactive discussion about the development of the project.

## Supported and Tested Platforms

- 64-bit Linux (Ubuntu, Debian), macOS, and Windows

## License

Apache License 2.0 - See [LICENSE](LICENSE) for details.
