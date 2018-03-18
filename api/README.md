## 日常の操作

### ソースの編集

編集すると自動で反映する。

(ただし、ソースの追加は自動では反映しないかも?)

### 依存関係の追加

```
docker-compose run --no-deps api dep ensure -add github.com/gorilla/mux
```

## 参考

- https://dev.to/craigchilds94/hot-reloading-go-programswith-docker-glide--fresh
- https://github.com/marcusolsson/goddd
