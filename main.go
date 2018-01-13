package main

import (
  "net/http"
  "log"
  "io/ioutil"
  "encoding/json"
  "bytes"
  "strings"
)

type Request struct {
  Url    string            `json:"url"`
  Method string            `json:"method"`
  Header map[string]string `json:"header"`
  Body   string            `json:"body"`
}

type Response struct {
  Url        string            `json:"url"`
  Method     string            `json:"method"`
  Header     map[string]string `json:"header"`
  Body       string            `json:"body"`
  Request    Request           `json:"request"`
  Status     int               `json:"status"`
  StatusText string            `json:"statusText"`
}

func JSONParse(request interface{}, body []byte) error {
  return json.Unmarshal(body, &request)
}

type Mapper map[string]*Response

func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
      w.Write([]byte(err.Error()))
      return
    }

    seriesRequest := make([]Request, 0)
    concurrentRequest := *new(map[string]*Request)

    client := &http.Client{
    }

    if err := JSONParse(&seriesRequest, body); err != nil {

      // 如果传过来的是一个Object，则并发执行
      err := JSONParse(&concurrentRequest, body)

      if err != nil {
        w.Write([]byte(err.Error()))
        return
      }

      var result = Mapper{}

      count := 0
      ch := make(chan bool, 5)

      for k, q := range concurrentRequest {
        result[k] = nil
        go func(k string, q *Request) {
          req, err := http.NewRequest(q.Method, q.Url, bytes.NewReader([]byte(q.Body)))

          if err != nil {
            result[k] = nil
            count++
            return
          }

          result[k] = &Response{
            Url:     q.Url,
            Method:  q.Method,
            Request: *q,
            Header:  map[string]string{},
          }

          result[k].Request.Header = map[string]string{}
          result[k].Request.Body = string(body)

          // customer header
          for k, v := range q.Header {
            req.Header.Add(k, v)
          }

          res, err := client.Do(req)

          result[k].Status = res.StatusCode
          result[k].StatusText = res.Status

          if err != nil {
            result[k] = nil
            count++
            return
          }

          // get request head
          for key, values := range req.Header {
            result[k].Request.Header[key] = strings.Join(values, ";")
          }

          // set response head
          for key, values := range res.Header {
            result[k].Header[key] = strings.Join(values, ";")
          }

          body, err := ioutil.ReadAll(res.Body)
          if err != nil {
            result[k] = nil
            count++
            return
          }

          result[k].Body = string(body)
          count++

          if count == len(concurrentRequest) {
            close(ch)
          }
        }(k, q)

      }

      for range ch {

      }

      b, err := json.Marshal(result)

      if err != nil {
        w.Write([]byte(err.Error()))
        return
      }

      w.Write(b)

    } else {
      // 如果传过来的是数组，则按照顺序执行
      result := make([]*Response, len(seriesRequest))

      for i, q := range seriesRequest {
        req, err := http.NewRequest(q.Method, q.Url, bytes.NewReader([]byte(q.Body)))

        result[i] = &Response{
          Url:     q.Url,
          Method:  q.Method,
          Request: q,
          Header:  map[string]string{},
        }

        result[i].Request.Body = string(body)

        if err != nil {
          result[i] = nil
          continue
        }
        for k, v := range q.Header {
          req.Header.Add(k, v)
        }

        res, err := client.Do(req)

        if err != nil {
          result[i] = nil
          continue
        }

        result[i].Status = res.StatusCode
        result[i].StatusText = res.Status

        // get request head
        for key, values := range req.Header {
          result[i].Request.Header[key] = strings.Join(values, ";")
        }

        // set response head
        for key, values := range res.Header {
          result[i].Header[key] = strings.Join(values, ";")
        }



        body, err := ioutil.ReadAll(res.Body)
        if err != nil {
          result[i] = nil
          return
        }

        result[i].Body = string(body)
      }

      b, err := json.Marshal(result)

      if err != nil {
        w.Write([]byte(err.Error()))
        return
      }

      w.Write(b)
    }

  })
  log.Fatal(http.ListenAndServe(":8086", nil))
}
