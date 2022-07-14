# Getting into Go

---

### Go 是什麼？
是由 Google 所開發的一種靜態型別、編譯型、並行運算，並具有垃圾回收功能的開源程式語言。
但 Go 並非像傳統的靜態型別和編譯型語言一樣，它擁有一些類似動態型別的特性，以及編譯過後的可執行檔裡就已包含一個有 Garbage Collection 的 Runtime 了。

### Go 有什麼優點？
- 很好的發揮多核心硬體的性能
- 支援跨平台編譯
- 快速及相對簡單的部署
&emsp;&emsp;只需要將編譯出來的二進位執行檔，直接丟到伺服器啟動就好了
- 不用額外設定 Web Server
&emsp;&emsp;不需要 Apache 或 Nginx 等相關服務，Go 本身就可以處理 HTTP 連線，代表不需要伺服器設定，相對的降低維護成本以及伺服器安全性問題

### Go 的缺點
- GC 部分的優化還是比不上 JVM 的 G1
- Go 的編譯器為了編譯速度，將不少編譯時期能做的優化都省略掉了
- 目前還不支援泛型
- 錯誤處理的方式 (但看法兩極)

### 參考來源
[The Go Programming Language - 2015 - Oreilly](https://www.gopl.io/)