#! /bin/bash
#
# helper, so you don't need to remember compiler syntax...

# build with clean cache takes long, but who cares
go clean -cache
go build .

# create package for MacOS, assuming we're building FROM MacOS
fyne package -os darwin -icon ../assets/logo.png


# !!!
# Docker needs to run for the next statements to work!
# !!!

# cross-compile for Windows
fyne-cross windows -app-id "schritt-2" -icon ../assets/logo.png

# cross-compile for Linux
fyne-cross linux -app-id "schritt-2" -icon ../assets/logo.png
