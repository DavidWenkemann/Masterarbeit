compile:
	echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=amd64 go build -o Monolith_AsyncCommunication/bin/main-linux-amd64 ./
	GOOS=windows GOARCH=amd64 go build -o Monolith_AsyncCommunication/bin/main-windows-amd64.exe ./
	GOOS=darwin GOARCH=arm64 go build -o Monolith_AsyncCommunication/bin/main-mac-m1-arm64 ./