package csv

import (
	"bytes"
	"encoding/csv"
	"errors"

	"github.com/thoohv5/common/encoding"
)

const Name = "csv"

var (
	ErrInvalid = errors.New("not csv resp")
)

func init() {
	encoding.RegisterCodec(codec{})
}

// codec is a Codec implementation with json.
type codec struct{}

func (codec) Marshal(v interface{}) ([]byte, error) {

	bytesBuffer := &bytes.Buffer{}
	// 写入UTF-8 BOM，避免使用Microsoft Excel打开乱码
	bytesBuffer.WriteString("\xEF\xBB\xBF")
	writer := csv.NewWriter(bytesBuffer)

	csvResp, ok := v.(*Response)
	if !ok {
		return nil, ErrInvalid
	}
	err := writer.Write(csvResp.GetTitle().GetCell())
	if err != nil {
		return nil, err
	}
	for _, item := range csvResp.GetContent() {
		err = writer.Write(item.GetCell())
		if err != nil {
			return nil, err
		}
	}

	writer.Flush()
	return bytesBuffer.Bytes(), nil
}

func (codec) Unmarshal(data []byte, v interface{}) error {
	return nil
}

func (codec) Name() string {
	return Name
}
