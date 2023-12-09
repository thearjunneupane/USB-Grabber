# Dump all files from the inserted USB to your Computer
### Build: `go build -ldflags="-H windowsgui" -o grabber.exe`
### Default Path: `%HOMEPATH%\grabbed`  
### To change Path: `start grabber.exe -dir <Absolute Directory Path>`
### To stop: `taskkill /F /IM grabber.exe /T` or end task in TskMngr
<h3>Note:</h3> 
 The grabbed folder can be navigate using shell.
<hr>

### Inspired By [Ginray/USB-Dumper](https://github.com/Ginray/USB-Dumper)