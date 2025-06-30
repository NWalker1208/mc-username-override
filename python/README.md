# Development Environment

## Requirements
- Python 3.13.3 or newer
- [ImageMagick](https://imagemagick.org/script/download.php)


## Setup

```sh
python -m venv .venv
./.venv/Scripts/activate
python -m pip install -r requirements.txt
```

## Build

```sh
mkdir build
magick -background transparent -define 'icon:auto-resize=256,128,96,64,48,32,24,16' icon.svg build/icon.ico
pyinstaller -F -n "Change Minecraft Username" --icon build/icon.ico change_username.py
```

The final output will appear in the `dist` directory.
