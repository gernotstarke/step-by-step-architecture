#!/bin/bash
fyne bundle -name Arc42LogoPNG -package resources arc42-logo.png  > bundled.go
fyne bundle -name PDFminionLogoPNG -append PDFminion-logo.png  >> bundled.go
# splash image still missing