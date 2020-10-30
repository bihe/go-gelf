package gelf

import "fmt"

func sendAndRecvCallStackOpts(msgData string, compress CompressType, opts WriterOptions) (*Message, error) {
	r, err := NewReader("127.0.0.1:0")
	if err != nil {
		return nil, fmt.Errorf("NewReader: %s", err)
	}

	w, err := NewUDPWriter(r.Addr(), opts)
	if err != nil {
		return nil, fmt.Errorf("NewUDPWriter: %s", err)
	}
	w.CompressionType = compress

	if _, err = w.Write([]byte(msgData)); err != nil {
		return nil, fmt.Errorf("w.Write: %s", err)
	}

	w.Close()
	return r.ReadMessage()
}
