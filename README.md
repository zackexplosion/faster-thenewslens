# 跑比較快的，關鍵評論網

覺得關鍵評論網太慢嗎？來試試看! <http://www.thenewslens.xyz>

本專案公開原始碼，任何人都可以修改與複製。

## Why

最近發現關鍵評論網的內容不錯，很喜歡他們的內容。但是載入的速度實在慢到讓人受不了，於是打開網頁開發工具檢查了一下。發現裡面的圖片都是直接從圖床要原圖下來的版本。

其實圖床本身就有提供壓縮的參數，只是不知道開發人員為什麼沒注意到。於是製作了這個 APP 可以讓使用者下載壓縮過的圖片，節省大家的網路費用，又同時可以欣賞到優質的內容。

## Compare

<http://www.thenewslens.com/post/82643/>

14.8M transfered, window onload in 33.67s, finished in 4.9min

![Imgur](http://i.imgur.com/La5eeMO.jpg)


<http://www.thenewslens.xyz/post/82643/>

5.2M transfered, window onload in 9.3s, finished in 1.9min

![Imgur](http://i.imgur.com/a2gHbFi.png)


## Contribute

git clone https://github.com/zackexplosion/the-faster-thenewslens.git && go run app.go

## How it works?

<http://zackexplosion.com/p/d73d>

## Read more
* <http://evanbyrne.com/blog/go-production-server-ubuntu-nginx>

