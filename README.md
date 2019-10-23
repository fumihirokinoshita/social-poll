# social-poll
SNS投票システム

## 環境構築(mac)

go のインストール
```
$ brew install go
```

Path の設定
```
$ mkdir $HOME/go
$ echo 'export GOPATH=$(go env GOPATH)' >> ~/.bash_profile
$ echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.bash_profile
$ source ~/.bash_profile
```

clone
```
$ mkdir $HOME/go/src
$ cd $HOME/go/src
$ git clone github.com/fumihirokinoshita/social-poll
```

[NSQ](http://nsq.io/)のインストール
```
$ brew install nsq
$ go get github.com/bitly/go-nsq // nsq driver for go
```

[MongoDB](https://docs.mongodb.com/manual/tutorial/install-mongodb-on-os-x/)のインストール
```
$ brew tap mongodb/brew
$ brew install mongodb-community@4.2
$ go get gopkg.in/mgo.v2 // mongodb driver for go
```