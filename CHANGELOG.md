# Changelog

## v2.0.0 — 2026-09-25

- `server.exit()` now uses Garry's Mod `engine.CloseServer()` instead of manually running `ShutDown` and terminating the process with `os.Exit(1)`.
- **Action required:** add `-allowquit` to the dedicated server launch command. The module no longer forces exit when that option is absent.
- Fixed module unload when the global `server` table has been removed.
- Removed the AMD64 v3 build requirement for broader CPU compatibility.
- Updated the installation, usage, and build documentation; added CI builds for changes to `main`.

## v1.0.0 — 2025-05-29

Initial release.
