package downloader

import "net/http"

func Download(url string, buffsize int, handler func(int, []byte)) error {
	b := make([]byte, buffsize)
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	for {
		n, err := response.Body.Read(b)
		if n <= 0 {
			break
		}
		if err != nil {
			return err
		}
		handler(n, b[:n])
	}

	return nil
}
