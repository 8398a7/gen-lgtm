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
  -r, --dir string     Converts gif images in a specified folder. The image name is "original-image-dist.gif"
  -d, --dist string    Distribution image path
  -h, --help           help for gen-lgtm
  -i, --image string   Overwrite the LGTM image
  -l, --loop int       Number of loops in the gif image. If 0 is specified, it is an infinite loop (default -1)
  -s, --src string     Source image path

Use "gen-lgtm [command] --help" for more information about a command.
```

## Install

### brew or linuxbrew

```bash
$ brew tap 8398a7/gen-lgtm
$ brew install gen-lgtm
```

### go get

```bash
$ go get -u github.com/8398a7/gen-lgtm
```

## Quick Start

You can run it through Docker.  
refs: https://hub.docker.com/repository/docker/8398a7/gen-lgtm

```bash
$ docker run -v $(pwd)/tmp:/tmp 8398a7/gen-lgtm -s /tmp/src.gif
$ open tmp/src-dist.gif
```

## Usage

```bash
$ gen-lgtm --src src.gif --dist dist.gif --loop 5 --image other_lgtm.png
```

If 0 is set to `loop`, this creates an infinite loop.  
The default value of `loop` is -1, which means it loops once.
