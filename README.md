imageserver
===========

[![Build Status](https://travis-ci.org/tektoh/imageserver.svg?branch=master)](https://travis-ci.org/tektoh/imageserver)

こんなかんじの画像サーバーつくるよてい

Architecture
-------------

* GoogleAppengine
* GoogleCloudStorage

GET
---

* 画像サイズ・形式をURLパラメーターで指定。
* 動的にサムネイルを作成・キャッシュする。
* サムネイルの元画像はPOST時に作成したベースサムネイルから最適なものを使用する。


```
GET /IMAGE_ID/SIZE.EXT
   SIZE: origin, full, large, middle, small, icon ...
   EXT: jpg, png
```

POST
---------

* 画像をアップロードして保存する。
* GET時のサムネイル作成時に使用するベースサムネイルをマルチスレッドで作成する。
* ベースサムネイルはEXIFを削除する
* オリジナル画像、ベースサムネイル画像をS3にPOSTする。
* 外部のAPIサーバーと認証チェックを行う。認証情報はキャッシュする。
* 外部のAPIサーバーに対して画像のメタ情報をPOSTする

```
POST /[IMAGE_ID]
   token: AUTH TOKEN
   filename:
   file:
```

DELETE
------

* 画像を削除する

```
DELETE /IMAGE_ID
```
