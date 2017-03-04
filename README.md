# CryptoGo ![Language Badge](https://img.shields.io/badge/Language-Go-blue.svg) ![Go Report](https://goreportcard.com/badge/github.com/isfonzar/CryptoGo) ![License Badge](https://img.shields.io/badge/License-MIT-blue.svg) ![Status Badge](https://img.shields.io/badge/Status-Beta-brightgreen.svg)

CryptoGo is a simple file encrypter for your day-to-day needs.

CryptoGo's goal is to be a simple tool to encrypt and password protect your files.

## Project Status

CryptoGo is on beta. Pull Requests [are welcome](https://github.com/isfonzar/CryptoGo#social-coding)

![](http://i.imgur.com/KimL2Xr.gif)

## Features

- STUPIDLY [EASY TO USE](https://github.com/isfonzar/CryptoGo#usage)
- Fast encryption and decryption processes
- Uses [filecrypt](https://github.com/isfonzar/filecrypt) libs
- Galois/Counter Mode (GCM) encryption (Extra secure, harder to bruteforce)

## Installation

### Option 1: Go Get

```bash
$ go get github.com/isfonzar/CryptoGo
$ CryptoGo
```

### Option 2: From source

```bash
$ git clone https://github.com/isfonzar/CryptoGo.git
$ cd CryptoGo/
$ go get -d
$ go build *.go
```
## Usage

### Encryption

```bash
# Encrypts a file
$ CryptoGo encrypt path/to/your/file
```

### Decryption

```bash
# Decrypts a file
$ CryptoGo decrypt path/to/your/file
```

### Show help

```bash
$ CryptoGo help
```

## Program Help

![](http://i.imgur.com/SLimwGt.png)

## Contributing

### Bug Reports & Feature Requests

Please use the [issue tracker](https://github.com/isfonzar/CryptoGo/issues) to report any bugs or file feature requests.

### Developing

PRs are welcome. To begin developing, do this:

```bash
$ git clone --recursive git@github.com:isfonzar/CryptoGo.git
$ cd CryptoGo/
```

## Social Coding

1. Create an issue to discuss about your idea
2. [Fork it] (https://github.com/isfonzar/CryptoGo/fork)
3. Create your feature branch (`git checkout -b my-new-feature`)
4. Commit your changes (`git commit -am 'Add some feature'`)
5. Push to the branch (`git push origin my-new-feature`)
6. Create a new Pull Request
7. Profit! :white_check_mark:

## Disclaimer

CryptoGo is still on beta. 
We will not be held responsible for any file loses that may occur due to bugs or misuse of the program distributed here.
Always keep a backup in those cases.
