# Building

The module uses [glua](https://github.com/Srlion/glua) and Go cgo. Build it on the target operating system with Go 1.24.3 or newer, Python 3, and a C compiler for the target architecture. A Go toolchain alone is not enough for `-buildmode=c-shared`.

## Linux

Install `gcc` and the 32-bit multilib toolchain for 32-bit builds. Then run:

```sh
python3 go_build.py --name serverexit --arch 64
python3 go_build.py --name serverexit --arch 32
```

## Windows

Install the matching 32-bit or 64-bit WinLibs MinGW-w64 toolchain under `C:\mingw32` or `C:\mingw64`. The build script adds its `bin` directory to `PATH` for the selected architecture.

```powershell
python go_build.py --name serverexit --arch 64
python go_build.py --name serverexit --arch 32
```

The four outputs are `bin/gmsv_serverexit_linux.dll`, `bin/gmsv_serverexit_linux64.dll`, `bin/gmsv_serverexit_win32.dll`, and `bin/gmsv_serverexit_win64.dll`. The script also accepts `--directory`, `--outdir`, `--cflags`, and `--ldflags`; run `python go_build.py --help` for details.

## CI and releases

GitHub Actions builds all four binaries on changes to `main` and on version tags. A tag beginning with `v` publishes a release only after all four builds succeed. Verify the assets attached to each release before deploying it to a server. The CI build checks compilation and packaging; behavior inside a running Garry's Mod server needs a server test.
