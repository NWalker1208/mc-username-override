# Development Environment

## Requirements
- Conda
- ImageMagick

## Setup

```sh
mkdir build
conda env create -f environment.yml
conda activate mc-username-override
```

## Build

```sh
magick -background transparent -define 'icon:auto-resize=256,128,96,64,48,32,24,16' icon.svg build/icon.ico
pyinstaller -D -n "Change Minecraft Username" --icon build/icon.ico change_username.py
```
