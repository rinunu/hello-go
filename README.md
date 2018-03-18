# hello go

go サンプル (RSS リーダー)

## セットアップ

- [Docker for Mac](https://docs.docker.com/docker-for-mac/) をインストールしてね。


```
docker-compose run flyway migrate
docker-compose run --no-deps api dep ensure
docker-compose run --no-deps bff npm install
docker-compose up
```

## 日常の操作

### サーバ群起動

```
docker-compose up
```

### DB migration

```
docker-compose run --no-deps flyway migrate
```

### サーバ群停止

"docker-compose up" を Ctrl+C

### ページ表示

```
open http://localhost:3000/
```

### mysql につなぐには
```
mysql -h 127.0.0.1 -P 33306 -u root -p
```

### コマンド実行

例

```
docker-compose run --no-deps api dep init
```

