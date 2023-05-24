# golang-http 超时机制

## 总结

其实就是 write-timeout 和 read-timeout 不作用于 serverHttp 函数，只作用于 conn 读写上的超时，而是需要传入个超时值的 context 下去

## Reference
- [Go net/http 超时机制完全手册](https://colobu.com/2016/07/01/the-complete-guide-to-golang-net-http-timeouts/)