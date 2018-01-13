## A server combine Http request and return

![License](https://img.shields.io/badge/license-Apache-green.svg)

## Usage

Copy and Paste into Chrome Dev Tool

#### Concurrent

```javascript
var xhr = new XMLHttpRequest();

xhr.open("POST", "https://go-http-combiner.herokuapp.com");

xhr.onreadystatechange = function(e) {
  console.log(e);
};

xhr.send(
  JSON.stringify({
    "baidu.com": {
      method: "GET",
      url: "https://www.baidu.com"
    },
    "cn.bing.com": {
      method: "GET",
      url: "https://cn.bing.com"
    },
    "sina.com": {
      method: "GET",
      url: "http://sina.com"
    }
  })
);
```

#### Series

```javascript
var xhr = new XMLHttpRequest();

xhr.open("POST", "https://go-http-combiner.herokuapp.com");

xhr.onreadystatechange = function(e) {
  console.log(e);
};

xhr.send(
  JSON.stringify([
    {
      method: "GET",
      url: "https://www.baidu.com"
    },
    {
      method: "GET",
      url: "https://cn.bing.com"
    },
    {
      method: "GET",
      url: "http://sina.com"
    }
  ])
);
```

## Contributing

[Contributing Guid](https://github.com/axetroy/http-combiner.go/blob/master/CONTRIBUTING.md)

## Contributors

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
| [<img src="https://avatars1.githubusercontent.com/u/9758711?v=3" width="100px;"/><br /><sub>Axetroy</sub>](http://axetroy.github.io)<br />[💻](https://github.com/axetroy/http-combiner.go/commits?author=axetroy) [🐛](https://github.com/axetroy/http-combiner.go/issues?q=author%3Aaxetroy) 🎨 |
| :---: |
<!-- ALL-CONTRIBUTORS-LIST:END -->

## License

[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Faxetroy%2Fhttp-combiner.go.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Faxetroy%2Fhttp-combiner.g?ref=badge_large)
