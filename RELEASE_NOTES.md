## v2.0.0

`server.exit()` now asks Garry's Mod to close the server through `engine.CloseServer()`. This allows the engine's normal shutdown lifecycle to run. The previous version called `ShutDown` manually and then terminated the process with `os.Exit(1)`.

**Upgrade requirement:** add `-allowquit` to the dedicated server launch command. Without it, `server.exit()` will not close the server. There is no forced-exit fallback.

The release also removes the AMD64 v3 CPU requirement, improves module unload handling, refreshes installation and build documentation, and runs four-platform builds on changes to `main`.

Choose the asset matching the server process architecture. The release binaries are build-verified; behavior in a running Garry's Mod server is not covered by CI.
