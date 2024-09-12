wget https://go.dev/dl/go1.20.7.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.20.7.linux-amd64.tar.gz
echo "export PATH=\$PATH:/usr/local/go/bin" >> ~/.profile
source ~/.profile
go version
sudo apt install make
sudo apt install binutils
make guest
go build .