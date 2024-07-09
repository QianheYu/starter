all: windows linux darwin

windows:
	GOOS=windows GOARCH=amd64 go build -ldflags "-w -s -extldflags '-static'" -gcflags '-trimpath -m'-o starter_$GOOS_$GOARCH.exe main.go
	GOOS=windows GOARCH=386 go build -ldflags "-w -s -extldflags '-static'" -gcflags '-trimpath -m'-o starter_$GOOS_$GOARCH.exe main.go
	GOOS=windows GOARCH=arm go build -ldflags "-w -s -extldflags '-static'" -gcflags '-trimpath -m'-o starter_$GOOS_$GOARCH.exe main.go
	GOOS=windows GOARCH=arm64 go build -ldflags "-w -s -extldflags '-static'" -gcflags '-trimpath -m'-o starter_$GOOS_$GOARCH.exe main.go
linux:
	GOOS=linux GOARCH=amd64 go build -ldflags "-w -s -extldflags '-static'" -gcflags '-trimpath -m'-o starter_$GOOS_$GOARCH main.go
	GOOS=linux GOARCH=386 go build -ldflags "-w -s -extldflags '-static'" -gcflags '-trimpath -m'-o starter_$GOOS_$GOARCH main.go
	GOOS=linux GOARCH=arm go build -ldflags "-w -s -extldflags '-static'" -gcflags '-trimpath -m'-o starter_$GOOS_$GOARCH main.go
	GOOS=linux GOARCH=arm64 go build -ldflags "-w -s -extldflags '-static'" -gcflags '-trimpath -m'-o starter_$GOOS_$GOARCH main.go
darwin:
	GOOS=darwin GOARCH=amd64 go build -ldflags "-w -s -extldflags '-static'" -gcflags '-trimpath -m'-o starter_$GOOS_$GOARCH main.go
	GOOS=darwin GOARCH=arm64 go build -ldflags "-w -s -extldflags '-static'" -gcflags '-trimpath -m'-o starter_$GOOS_$GOARCH main.go