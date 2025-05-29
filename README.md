# gmsv_serverexit

A module for Garry's Mod that allows you to safely shut down the server from Lua

## Installation
1. Download the appropriate file from [Releases](https://github.com/Dobytchick/gmsv_serverexit/releases)
2. Place it in your server's `lua/bin` folder
3. Restart your server

## Usage
```lua
server.exit()
```

Before shutdown, the Lua [ShutDown](https://wiki.facepunch.com/gmod/GM:ShutDown) hook will be called for your cleanup code.