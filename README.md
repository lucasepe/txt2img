
# txt2img - Turn plain text into pixel-perfect PNGs

> A simple and fast CLI tool that converts any text—ASCII or Unicode—into crisp PNG images.

## Features

- Convert text into a PNG image
- Read text from stdin, a file, or terminal input
- Automatic or manual image sizing
- Customizable text color, background color, and margins
- Optionally force square images
- Handles tabs with configurable tab width

## Usage

```
txt2img [flags] [file]
```

### Input

* **Pipe from stdin**:

```bash
echo "Hello World" | txt2img -o out.png
```

* **Read from a file**:

```bash
txt2img -o out.png myfile.txt
```

* **Manual input** (if no file is provided and stdin is not piped):

```
txt2img -o out.png
Type your text and finish with CTRL+D
```

## Support

All tools are completely free to use, with every feature fully unlocked and accessible.

If you find one or more of these tool helpful, please consider supporting its development with a donation.

Your contribution, no matter the amount, helps cover the time and effort dedicated to creating and maintaining these tools, ensuring they remain free and receive continuous improvements.

Every bit of support makes a meaningful difference and allows me to focus on building more tools that solve real-world challenges.

Thank you for your generosity and for being part of this journey!

[![Donate with PayPal](https://img.shields.io/badge/💸-Tip%20me%20on%20PayPal-0070ba?style=for-the-badge&logo=paypal&logoColor=white)](https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=FV575PVWGXZBY&source=url)

## How To Install

### Using the _install.sh_ script (macOS & linux)

Simply run the following command in your terminal:

```sh
curl -sL https://raw.githubusercontent.com/lucasepe/txt2img/main/install.sh | bash
```

This script will:

- Detect your operating system and architecture
- Download the latest release binary
- Install it into _/usr/local/bin_ (requires sudo)
  - otherwise fallback to _$HOME/.local/bin_ 
- Make sure the install directory is in your _PATH_ environment variable


### Manually download the latest binaries from the [releases page](https://github.com/lucasepe/txt2img/releases/latest):

- [macOS](https://github.com/lucasepe/txt2img/releases/latest)
- [Windows](https://github.com/lucasepe/txt2img/releases/latest)
- [Linux (arm64)](https://github.com/lucasepe/txt2img/releases/latest)
- [Linux (amd64)](https://github.com/lucasepe/txt2img/releases/latest)

Unpack the binary into any directory that is part of your _PATH_.

## If you have [Go](https://go.dev/dl/) installed

You can also install using:

```bash
go install github.com/lucasepe/txt2img@latest
```

Make sure your `$GOPATH/bin` is in your PATH to run the tool from anywhere.
