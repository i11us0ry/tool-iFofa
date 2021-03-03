windres.exe -i resDefine.rc  -o defaultRes_windows_amd64.syso -F pe-x86-64
go build -o out\tmp.exe -i  -trimpath -buildmode=exe -tags tempdll -ldflags="-s -w -H windowsgui"
upx64 out\tmp.exe -f -o out\ifofa.exe
del /f /q out\tmp.exe
pause
::-H windowsgui