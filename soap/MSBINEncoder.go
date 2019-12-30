package soap

import (
	"bufio"
	"bytes"
	"encoding/xml"
	"io"

	"github.com/khoad/msbingo/nbfs"
)

type MSBINEncoder struct {
	writer *bufio.Writer
}

type MSBINDecoder struct {
	reader *bufio.Reader
}

func newMSBINEncoder(w io.Writer) *MSBINEncoder {
	return &MSBINEncoder{
		writer: bufio.NewWriter(w),
	}
}

func newMSBINDecoder(rd io.Reader) *MSBINDecoder {
	return &MSBINDecoder{
		reader: bufio.NewReader(rd),
	}
}

func (enc *MSBINEncoder) Encode(v interface{}) error {
	buf := new(bytes.Buffer)
	xmlEncoder := xml.NewEncoder(buf)
	xmlEncoder.Encode(v)
	xmlEncoder.Flush()
	data := buf.Bytes()
	dataBuf := bytes.NewBuffer(data)
	encodedXml, err := nbfs.NewEncoder().Encode(dataBuf)
	if err != nil {
		return err
	}
	_, err = enc.writer.Write(encodedXml)
	return err
}

func (enc *MSBINEncoder) Flush() error {
	return enc.writer.Flush()
}

func (dec *MSBINDecoder) Decode(v interface{}) error {
	xmlRes, err := nbfs.NewDecoder().Decode(dec.reader)
	if err != nil {
		return err
	}
	err = xml.NewDecoder(bytes.NewBufferString(xmlRes)).Decode(v)
	if err != nil {
		return err
	}
	return nil
}
