# Concurrency in Go

---

### Communicating Sequential Processes (CSP)
> TODO: 待補充和了解

Go 所使用的 concurrency model

### runtime.Gosched
類似讓 Go 的排程器在當下這時間點，去處理一個或多個其他的 task (goroutine)，不能用這個方法來保證當下其它的 task 一定執行完畢，因爲可能有些 task 是跑 database query 或是其它一樣需要等待的操作，所以不應該依賴該方法來確保 tasks 都執行完畢，需要使用 channel 來建立一套完善的機制或 **wait groups**。

當呼叫 goroutine 時，是和 Go 的 runtime 說如果可以的話快點幫我執行這個 function，**可能不會是立即執行**的。
如果跑在單處器的機器上，幾乎可預期它不會立即執行。

### Goroutine
非常類似於執行緒，運行是透過多工的方式並行。

每個 OS 執行緒都有固定大小的**堆疊**記憶體區塊 (通常 2 ~ 4MB)，是用來儲存進行中的函式呼叫或呼叫另一個函式而暫時停止時的區域變數的工作區域。

而 goroutine 是可成長的堆疊從 2KB 開始，如同 OS 執行緒一樣 goroutine 的堆疊保存活動中的區域變數與暫停中的函式呼叫，但與 OS 執行緒不同的是 goroutine 的堆疊不是固定的，它會隨著需求而縮放。

Goroutine 可以在 OS 執行緒之間移動，由 Go 排程器來規劃，來最好的運用運算資源，不造成某些執行緒阻塞時而另一個執行緒閒置沒事做。

Go 的執行期本身具有使用稱為 **m:n 排程** 技術的排程器，名稱來自於它在 n 個 OS 執行緒上複合 (排程) m 個 goroutine，Go 排程器的工作類似於 OS 核心排程器，但它只管理單一 Go 程式上的 goroutine。

與 OS 的執行緒排程器不同的是，Go 排程器並非由硬體的計時器週期性來調用，而是由 Go 自己內部特定的結構處理。例如 goroutine 呼叫 time.Sleep 或是受到 channel, mutex 阻斷時，排程器會讓它沈睡並執行其他 goroutine 直到該叫醒它的時候。

由於 OS 執行緒由核心排程，從一個執行緒轉移控制到另一個執行緒需要完整的**背景交換**，也就是儲存一個執行緒的狀態到記憶體中、恢復另一個執行緒的狀態、並更新排程器的資料結構。因為位置性的差異與所需記憶體存取數量，所以這個操作很慢，在存取記憶體所需的 CPU 時脈更多時還會更糟。

由於不需要交換核心背景，因此 goroutine 的重新排程比執行緒的重新排程成本更低。

Go 的排程器使用稱為 GOMAXPROCS 的參數來決定同時有多少 OS 執行緒可以執行 Go 的程式，預設值為該機器的 CPU 數量，這個參數也就是 **m:n 排程** 中的 n。

Goroutine 沒有識別名稱，是當初設計上有意而為之的，設計者認為執行緒區域儲存體常常被濫用，就像大量依賴全域變數一樣，會引發**超距作用**也就是函式的行為並非完全由所帶入的參數決定，而是執行緒的識別名稱來決定。

Never start a goroutine without knowning when it will stop.

### Channel
用來讓 goroutine 間可以互相發送值的一個通訊機制。

### 不要過度使用 Channel
> TODO: 待補充和了解詳細

- 對性能有所影響
- 會增加 codebase 的複雜度
- 記憶體佔用很多

### 參考來源
[The Go Programming Language - 2015 - Oreilly](https://www.gopl.io/)
