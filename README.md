# gke-demo-app

勉強会で使う簡単なアプリケーション(API)

## RUN

```bash
docker run --rm -p 80:8080 -it grandcolline/gke-demo-app
```

## USAGE

```bash
GET /           : バージョン確認
GET /down       : アプリケーション終了
GET /fibo?n=5   : n番目のフィボナッチ数を返す
```
