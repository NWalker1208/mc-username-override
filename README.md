# MC Username Override

A script to override one's username when playing Minecraft. This only works locally, but should allow multiple players to join a LAN game with the same Minecraft account.

## How It Works

First, Minecraft must be opened the typical way using the Minecraft Launcher. Then, when the script is run, it finds the arguments that were passed to the Minecraft client, alters the username value, and restarts Minecraft with the modified arguments.

## Development Environment

### Requirements
- Conda
- ImageMagick

### Setup

```sh
mkdir build
conda env create -f environment.yml
conda activate mc-username-override
```

### Build

```sh
magick -background transparent -define 'icon:auto-resize=256,128,96,64,48,32,24,16' icon.svg build/icon.ico
pyinstaller -F -n "Change Minecraft Username" --icon build/icon.ico change_username.py
```
