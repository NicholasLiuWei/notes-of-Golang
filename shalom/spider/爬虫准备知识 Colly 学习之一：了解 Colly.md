# 爬虫准备知识 Colly 学习之一：了解 Colly

## 简介

[Colly](https://github.com/gocolly/colly) 是一个快如闪电的优雅的开源爬虫框架，使用 Go 语言实现。Colly 提供了一个简洁的接口可用于写任何类型的爬虫（crawler/scraper/spider）。

使用 Colly 可以轻松地从网站中提取结构化的数据，这些数据可以广泛的应用于如数据挖掘、数据处理或存档等各种场景。

## Colly 主要的特性

1. 简洁的 API
2. 快速（单核上，>1k request/sec）
3. 管理每个域名的请求延迟和最大并发数
4. 自动的 cookie 和 session 处理
5. 支持同步、异步和并行爬取
6. 支持缓存
7. 自动编码非 unicode 的响应
8. Robots.txt 支持
9. 分布式爬取
10. 通过环境变量进行配置
11. 可以方便的进行扩展
12. Google App Engine 支持

## 安装

```
go get -u github.com/gocolly/colly/...
```

## 使用示例

抓取 Colly 官网：

```go
func main() {
    c := colly.NewCollector()

    // Find and visit all links
    c.OnHTML("a[href]", func(e *colly.HTMLElement) {
        e.Request.Visit(e.Attr("href"))
    })

    c.OnRequest(func(r *colly.Request) {
        fmt.Println("Visiting", r.URL)
    })

    c.Visit("http://go-colly.org/")
}
```

这么一段简单的代码就可以爬取 <http://go-colly.org/> 和外链的网站，以及外链的外链。

可见，使用 Colly，您可以构建各种复杂的 web 爬虫，从简单的爬虫到处理数百万网页的复杂异步网站爬网程序。Colly 提供了用于执行网络请求和处理接收到的内容的 API （例如，与 HTML 文档的 DOM 树交互）。

如果我们只想抓取 Colly 官网，我们只需在实例化 Collector 时加上一个选项：

```go
c := colly.NewCollector(colly.AllowedDomains("go-colly.org"))
```

