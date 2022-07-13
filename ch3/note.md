# Concurrency in Go

---

### Communicating Sequential Processes (CSP)
> TODO: 待補充和了解

Go 所使用的 concurrency model

### runtime.Gosched
類似讓 Go 的排程器在當下這時間點，去處理一個或多個其他的 task (goroutine)，不能用這個方法來保證當下其它的 task 一定執行完畢，因爲可能有些 task 是跑 database query 或是其它一樣需要等待的操作，所以不應該依賴該方法來確保 tasks 都執行完畢，需要使用 channel 來建立一套完善的機制或 **wait groups**。

當呼叫 goroutine 時，是和 Go 的 runtime 說如果可以的話快點幫我執行這個 function，**可能不會是立即執行**的。
如果跑在單處器的機器上，幾乎可預期它不會立即執行。

### 不要過度使用 Channel
> TODO: 待補充和了解詳細

- 對性能有所影響
- 會增加 codebase 的複雜度
- 記憶體佔用很多
