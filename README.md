imageserver
===========

[![Build Status](https://travis-ci.org/tektoh/imageserver.svg?branch=master)](https://travis-ci.org/tektoh/imageserver)

こんなかんじの画像サーバーつくるよてい

GET
---

* 画像サイズをURLパラメーターで指定。

* /*IMAGE_ID*/*SIZE*.*EXT*
    * SIZE: origin, full, large, middle, small, icon ...
    * EXT: jpg, png

POST, PUT
---------

* 画像をアップロードして保存する。必要なら加工する。

* /[*IMAGE_ID*]
    * token: AUTH TOKEN
    * rpc:
    * filename:
    * file:

DELETE
------

* 画像を削除する

* /*IMAGE_ID*

Hook
----

* 処理の途中で行う処理を柔軟に追加できる

Worker
-------

* マルチスレッド
