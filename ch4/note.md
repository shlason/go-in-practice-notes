# Handling errors and panics

---

### Errors and Panics
error 表示某個特定的工作未能成功完成
panic 類似於發生例外，可能是當初 code 有東西沒寫好之類的問題，若沒有特別捕捉做處理 (defer 中 recover) 會導致程式退出

Go 官方表示不提供例外，是說到他們認為 try-catch-finally 會讓 codebase 變得錯綜複雜，也使得原本常常應該被處理的**錯誤**被視為**例外**，但其實 Go 還是可以用 panic 來達成類似拋出例外。
Go 在設計上就鼓勵大家用 error 來明確地做錯誤檢查和相關處理，不過如果是巢狀深層的呼叫使用 panic 還是比較好回傳上去，但這樣的行為在 Go 中盡量以 pakcage 為一個範疇，並且可以的話在 package 內確實捕捉來轉為一個相對明確的 error 來回傳會更好。

### 參考來源
[Why does Go not have exceptions? - Go FAQ](https://go.dev/doc/faq#exceptions)
[Go 的 error 與 panic - iThome | 林信良](https://www.ithome.com.tw/voice/103455)
