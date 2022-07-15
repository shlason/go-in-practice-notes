# Handling errors and panics

---

### Errors and Panics
error 表示某個特定的工作未能成功完成

panic 類似於發生例外，可能是當初 code 有東西沒寫好之類的問題，若沒有特別捕捉做處理 (defer 中 recover) 會導致程式退出

Go 官方表示不提供例外，是說到他們認為 try-catch-finally 會讓 codebase 變得錯綜複雜，也使得原本常常應該被處理的**錯誤**被視為**例外**，但其實 Go 還是可以用 panic 來達成類似拋出例外。

Go 在設計上就鼓勵大家用 error 來明確地做錯誤檢查和相關處理，不過如果是巢狀深層的呼叫使用 panic 還是比較好回傳上去，但這樣的行為在 Go 中盡量以 pakcage 為一個範疇，並且可以的話在 package 內確實捕捉來轉為一個相對明確的 error 來回傳會更好。

### Deferred functions
>以下所指的“使用外部變數“皆是代表使用 deferred function 這個 scope 以外的變數
- 一個 function 內雖然可以宣告多個 deferred function 但非必要情況下盡量不要
- Deferred function 的呼叫順序是 LIFO
- 盡可能放在 function 內相對的最上方
- 若使用外部變數時，而那個變數的宣告又不是在 deferred function 之前，則會在編譯階段噴錯
- 一些釋放資源的操作 (close files, network connnections 之類的) 盡量使用 deferred function 來處理最好，可確保當處理發生 errors 或 panic 時，也有正確處理到以釋放資源
- Deferred function 內會形成一個閉包，當使用外部變數時會以參考的方式參考那個外部變數，所以等到 deferred function 開始執行時，它會使用到那個變數目前最新的狀態的值，而不是當初宣告時當下的值

### Panics on goroutines
呼叫 goroutine 時它們的 function call stack 會斷掉，goroutine 會自成一路變為最根層的那個，因此當在 goroutine 內發生 panic 時需要在裡面自己處理，不能依賴呼叫它的 function 來處理

例如有個 call stack: func A -> func B -> go func C -> func D (panic)

當 func D 發生 panic 時若自身又沒有特別處理 recovery，則會開始往回拋錯到 func C 而這時因為 goroutine 呼叫時自成一個新的 call stack，因此 func C 便是該 call stack 的 root 不會再繼續拋回給 func B，若 func C 也沒有處理 recovery 則會將整個程式退出 crash 掉

### Handle panics
若要寫一個 library 比較好的方式是將不管是 user 傳送的 callback function 或自身 library code 的相關 panic 問題，就在 library 內處理掉然後 log 出足夠豐富且有用的資訊以便後續修改或 debug，不要預期每個使用者都要知道或是他們可能不應該需要自己去處理這些例外

### 參考來源
[Why does Go not have exceptions? - Go FAQ](https://go.dev/doc/faq#exceptions)

[Go 的 error 與 panic - iThome | 林信良](https://www.ithome.com.tw/voice/103455)
