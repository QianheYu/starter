all: windows linux darwin

windows:
	@GOOS=windows GOARCH=amd64 go build -ldflags "-w -s -extldflags '-static'" -gcflags '-trimpath -m' -o bin/starter_windows_amd64.exe main.go
	@GOOS=windows GOARCH=386 go build -ldflags "-w -s -extldflags '-static'" -gcflags '-trimpath -m' -o bin/starter_windows_i386.exe main.go
	@GOOS=windows GOARCH=arm go build -ldflags "-w -s -extldflags '-static'" -gcflags '-trimpath -m' -o bin/starter_windows_arm.exe main.go
	@GOOS=windows GOARCH=arm64 go build -ldflags "-w -s -extldflags '-static'" -gcflags '-trimpath -m' -o bin/starter_windows_arm64.exe main.go
linux:
	@GOOS=linux GOARCH=amd64 go build -ldflags "-w -s -extldflags '-static'" -gcflags '-trimpath -m' -o bin/starter_linux_amd64 main.go
	@GOOS=linux GOARCH=386 go build -ldflags "-w -s -extldflags '-static'" -gcflags '-trimpath -m' -o bin/starter_linux_i386 main.go
	@GOOS=linux GOARCH=arm go build -ldflags "-w -s -extldflags '-static'" -gcflags '-trimpath -m' -o bin/starter_linux_arm main.go
	@GOOS=linux GOARCH=arm64 go build -ldflags "-w -s -extldflags '-static'" -gcflags '-trimpath -m' -o bin/starter_linux_arm64 main.go
darwin:
	@GOOS=darwin GOARCH=amd64 go build -ldflags "-w -s -extldflags '-static'" -gcflags '-trimpath -m' -o bin/starter_darwin_amd64 main.go
	@GOOS=darwin GOARCH=arm64 go build -ldflags "-w -s -extldflags '-static'" -gcflags '-trimpath -m' -o bin/starter_darwin_arm64 main.go
