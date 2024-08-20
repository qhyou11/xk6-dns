# xk6-dns

A k6 extension for dns lookup test.

## Build

To build a `k6` binary with this plugin, first ensure you have the prerequisites:

- [Go toolchain](https://go101.org/article/go-toolchain.html)
- Git

Then:

1. Install `xk6`:

  ```shell
  go install github.com/k6io/xk6/cmd/xk6@latest
  ```

2. Build the binary:

  ```shell
  xk6 build master \
    --with github.com/qhyou11/xk6-dns
  ```

## Example

```javascript
import dns from 'k6/x/dns';
import { check } from 'k6';



export default function () {
  dns.Setup()
  _, _, err := dns.Exchange("baidu.com","10.20.88.88:53")

  check(err, {
    'no err': (r) => r == nil,
  });
}
```
