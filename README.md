# gen-lgtm

![](https://github.com/8398a7/gen-lgtm/workflows/release/badge.svg)
![](https://img.shields.io/github/license/8398a7/gen-lgtm?color=brightgreen)
![](https://img.shields.io/github/v/release/8398a7/gen-lgtm?color=brightgreen)

Convert your favorite images to LGTM images.

![](https://user-images.githubusercontent.com/8043276/89729643-a229d980-da72-11ea-9264-5c537cd64826.gif)

## Install

TBD

## Quick Start

You can run it through Docker.

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
