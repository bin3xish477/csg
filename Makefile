windows:
	GOOS=windows GOARCH=amd64 go build -o ./csg.exe ./main.go
windows-release:
	GOOS=windows GOARCH=amd64 go build -o ./csg.exe ./main.go
	tar -czvf csg-windows-amd64.tar.gz ./csg.exe ./README.md
	sha256sum csg-windows-amd64.tar.gz > csg-windows-amd64.tar.gz.sha256sum
	GOOS=windows GOARCH=386 go build -o ./csg.exe ./main.go
	tar -czvf csg-windows-386.tar.gz ./csg.exe ./README.md
	sha256sum csg-windows-386.tar.gz > csg-windows-386.tar.gz.sha256sum

darwin:
	GOOS=darwin GOARCH=amd64 go build -o ./csg ./main.go
darwin-release:
	GOOS=darwin GOARCH=amd64 go build -o ./csg ./main.go
	tar -czvf csg-darwin-amd64.tar.gz ./csg ./README.md
	sha256sum csg-darwin-amd64.tar.gz > csg-darwin-amd64.tar.gz.sha256sum

linux:
	GOOS=linux GOARCH=amd64 go build -o ./csg ./main.go
linux-release:
	GOOS=linux GOARCH=amd64 go build -o ./csg ./main.go
	tar -czvf csg-linux-amd64.tar.gz ./csg ./README.md
	sha256sum csg-linux-amd64.tar.gz > csg-linux-amd64.tar.gz.sha256sum
	GOOS=linux GOARCH=386 go build -o ./csg ./main.go
	tar -czvf csg-linux-386.tar.gz ./csg ./README.md
	sha256sum csg-linux-386.tar.gz > csg-linux-386.tar.gz.sha256sum

releases: linux-release windows-release darwin-release

clean:
	@if [ -f csg ]; then rm csg; fi
	@if [ -f csg.exe ]; then rm csg.exe; fi

