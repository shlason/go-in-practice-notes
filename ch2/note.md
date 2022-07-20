# A solid foundation

---

### The Twelve-Factor App
> TODO: 待補充和了解
1. Codebase
2. Dependencies
3. Config
4. Backing services
5. Build, release, run
6. Processes
7. Port binding
8. Concurrency
9. Disposability
10. Dev/prod parity
11. Logs
12. Admin processes

### Environment variables
程式碼進行開源或洩露出去的話，帳號密碼也會跟著一起洩漏。預防的方法就是在程式碼裡面使用類似挖個坑，執行時再把帳號密碼填進去。

以及因為環境變數存在於環境中，所以在不同環境 (dev, stage, prod) 也不需要特別修改只要把變數填好就好。

### Graceful shutdown
Graceful shutdown 就是先停止收 request，然後把已經收下來的 request 全部執行完，再把服務關掉。

實際原理是寫一段 code 去攔截送到 OS 的關閉信號 (signal)，立好一個類似服務要關閉了的 flag，程式每次執行到某個階段 (結束)，如果檢查到這個 flag 有立的話，則進行 os.Exit 等類似的關閉方法。

### signal.Notify 使用 buffered channel
> TODO: 待補充和了解

### 參考來源
[The Twelve-Factor App](https://12factor.net/)
[為什麼 signal.Notify 要使用 buffered channel](https://blog.wu-boy.com/2021/03/why-use-buffered-channel-in-signal-notify/)
[[Go 教學] 什麼是 graceful shutdown?](https://blog.wu-boy.com/2020/02/what-is-graceful-shutdown-in-golang/)
