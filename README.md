# gmsv_serverexit

[![Build module](https://github.com/Dobytchick/gmsv_serverexit/actions/workflows/build.yml/badge.svg)](https://github.com/Dobytchick/gmsv_serverexit/actions/workflows/build.yml)
[![Latest release](https://img.shields.io/github/v/release/Dobytchick/gmsv_serverexit)](https://github.com/Dobytchick/gmsv_serverexit/releases/latest)

A server-side Garry's Mod binary module that exposes `server.exit()` to Lua. It asks the engine to close the server, so the normal shutdown lifecycle can run.

> **Required:** start the dedicated server with `-allowquit`. Without that launch option, `engine.CloseServer()` will not exit the server. This module does not force a process exit.

## Install

1. Download the binary for your server's operating system and architecture from [Releases](https://github.com/Dobytchick/gmsv_serverexit/releases/latest).
2. Copy the `.dll` file into the server's `garrysmod/lua/bin/` directory. Keep the filename unchanged.
3. Add `-allowquit` to the server launch command and restart the server.
4. Load the module from server-side Lua, for example in `garrysmod/lua/autorun/server/serverexit.lua`:

   ```lua
   require("serverexit")
   ```

| Release file | Server platform |
| --- | --- |
| `gmsv_serverexit_linux.dll` | Linux 32-bit |
| `gmsv_serverexit_linux64.dll` | Linux 64-bit |
| `gmsv_serverexit_win32.dll` | Windows 32-bit |
| `gmsv_serverexit_win64.dll` | Windows 64-bit |

Garry's Mod uses the `.dll` module suffix on Linux as well. Choose the binary matching the *server process*, not merely the host operating system.

## Use

Call this only from server-side Lua when you intend to stop the server:

```lua
server.exit()
```

The module calls [`engine.CloseServer()`](https://wiki.facepunch.com/gmod/engine.CloseServer). Garry's Mod invokes its [`ShutDown` hook](https://wiki.facepunch.com/gmod/GM:ShutDown) as part of the normal shutdown. The module does not run that hook manually or use a timed `os.Exit` fallback.

```lua
hook.Add("ShutDown", "SaveBeforeExit", function()
    -- Save state synchronously here.
end)
```

The `ShutDown` hook also runs on other Lua shutdown events, such as map changes. Do not treat it as a signal that `server.exit()` was called specifically.

## Build from source

See [Building](docs/BUILDING.md) for toolchains, commands, output names, and the release workflow. Changes are recorded in the [changelog](CHANGELOG.md).
