SET CGO_ENABLED=0
SET GOOS=darwin3
SET GOARCH=amd64
go build --ldflags "-s -w" -o darwin3_amd64

SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=386
go build --ldflags "-s -w" -o linux_386

SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=arm
go build --ldflags "-s -w" -o linux_arm


SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=mips32le 
go build --ldflags "-s -w" -o linux_mips32le

SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build --ldflags "-s -w" -o linux_amd64

SET CGO_ENABLED=0
SET GOOS=windows
SET GOARCH=386
go build --ldflags "-s -w" -o windows_386.exe

SET CGO_ENABLED=0
SET GOOS=windows
SET GOARCH=amd64
go build --ldflags "-s -w" -o windows_amd64.exe