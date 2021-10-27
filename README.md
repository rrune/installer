# Installer
simple installer script in go.  
  
Downloads a zip file defined in `config.json` and unzips it into a directory chosen by the user.
# Build
Change the URL in `config.json`
```
go build -ldflags -H=windowsgui
```
# Notes
- There are three branches, each of them using a different Dialog window framework. All of them should work on Windows, Linux and MacOS. Linux and MacOS are untested
- `filename` in `config.json` can be anything
- `temp`in `config.json`can be anything, but should be something that doesn't exist.

# Todo
- Progress Bar
- Make Temp folder not override existing folder
