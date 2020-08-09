# gen-lgtm

![](https://github.com/8398a7/gen-lgtm/workflows/release/badge.svg)
![](https://img.shields.io/github/license/8398a7/gen-lgtm?color=brightgreen)
![](https://img.shields.io/github/v/release/8398a7/gen-lgtm?color=brightgreen)

Convert your favorite images to LGTM images.

![](https://user-images.githubusercontent.com/8043276/89729643-a229d980-da72-11ea-9264-5c537cd64826.gif)

```
Generate LGTM gif (source: https://github.com/8398a7/gen-lgtm)

Usage:
  gen-lgtm [flags]
  gen-lgtm [command]

Available Commands:
  help        Help about any command
  version     Print version information

Flags:
  -d, --dist string    Distribution image path (default "dist.gif")
  -h, --help           help for gen-lgtm
  -i, --image string   Overwrite the LGTM image
  -l, --loop int       Number of loops in the gif image. If 0 is specified, it is an infinite loop (default -1)
  -s, --src string     Source image path

Use "gen-lgtm [command] --help" for more information about a command.
```

## Install

macOS (x86_64)

```bash
$ curl -sLo gen-lgtm.tar.gz https://github.com/8398a7/gen-lgtm/releases/latest/download/gen-lgtm_Darwin_x86_64.tar.gz
$ tar -zxvf gen-lgtm.tar.gz
$ mv gen-lgtm /usr/local/bin
```

Linux (x86_64)

```bash
$ curl -sLo gen-lgtm.tar.gz https://github.com/8398a7/gen-lgtm/releases/latest/download/gen-lgtm_Linux_x86_64.tar.gz
$ tar -zxvf gen-lgtm.tar.gz
$ mv gen-lgtm /usr/local/bin
```

or

go get

```bash
go get -u github.com/8398a7/gen-lgtm
```

## Quick Start

You can run it through Docker.  
refs: https://hub.docker.com/repository/docker/8398a7/gen-lgtm

```bash
$ docker run -v $(pwd)/tmp:/tmp 8398a7/gen-lgtm -s /tmp/sample3.gif -d /tmp/dist.gif
$ open tmp/dist.gif
```

## Usage

```bash
$ gen-lgtm --src src.gif --dist dist.gif --loop 5
```

If 0 is set to `loop`, this creates an infinite loop.  
The default value of `loop` is -1, which means it loops once.
