package chef

import (
	"url"
	"fmt"
	"log"
	"encoding/base64"
	"crypto/sha1"
)

const (
	HeaderAccept = "text/json"
	HeaderXChefVersion = "10.24.0"
	HeaderXOpsSign = "version=1.0"
}

const (
  defaultTimeout = 30 * time.Second
)

type Client struct {
	url url.URL
	userId string
	timeout time.Duration
}


func New(url string) Client {
	// read config
	// set up
	return Client{url: url.Parse(url), userId: "FIXME"}
}

func (client Client) Get(endpoint string) (result, err) {
	response, err := client.request(endpoint)
	// unmarshal
	return
}


func transport(netw, addr string, timeout time.Duration) (connection net.Conn, err error) {
  if timeout == 0 {
    timeout = defaultTimeout
  }
  deadline := time.Now().Add(timeout)
  connection, err = net.DialTimeout(netw, addr, timeout)
  if err == nil {
    connection.SetDeadline(deadline)
  }
  return
}


func (client Client) request(endpoint string) (response, err) {
  httpClient := http.Client{
    Transport: &http.Transport{
      Dial: func(netw, addr string) (net.Conn, error) { return transport(netw, addr, client.timeout) },
    },
  }

	url := client.url + path
	request, err := http.NewRequest("GET", url, nil)

	request.Header.Add("Accept", HeaderAccept)
	request.Header.Add("Host", client.url.Host)
	request.Header.Add("X-Chef-Version", HeaderXChefVersion)

	return httpClient.Do(client.sign(request))
}

func (client *Client) SetTimeout(timeout time.Duration) {
  client.timeout
}

func (client *Client) sign(request http.Request) (http.Request) {
	contentHash := hash(request.Body)

	request.Header.Add("X-Ops-Sign", HeaderXOpsSign)
	request.Header.Add("X-Ops-UserId", client.userId)
	request.Header.Add("X-Ops-Timestamp", client.timestamp())
	request.Header.Add("X-Ops-Content-Hash", contentHash)

	canonical_request := fmt.Sprintf(
		"Method:%s\nHashed Path:%s\nX-Ops-Content-Hash:%s\nX-Ops-Timestamp:%s\nX-Ops-UserId:%s",
		request.Method, hash(request.URL.path), contentHash, client.timestamp(), client.userId
	)

	encrypted_request = client.encrypt(canonical_request)

	signature := make([]byte, base64.StdEncoding.EncodedLen(len(encrypted_request)))
	base64.StdEncoding.Encode(signature, encrypted_request)

	for i, line =: range(signature.split...) {
		request.Header.Add(fmt.Sprintf("X-Ops-Authorization-%d", i+1)
	}
}

func (client *Client) timestamp() (string) {
	return time.Now().Format(time.RFC3339)
}

func (client *Client) encrypt(string) {
	client.privatKey....
}

