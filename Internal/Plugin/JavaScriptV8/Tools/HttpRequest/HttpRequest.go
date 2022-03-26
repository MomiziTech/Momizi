/*
 * @Author: NyanCatda
 * @Date: 2022-03-26 10:21:35
 * @LastEditTime: 2022-03-26 20:43:15
 * @LastEditors: NyanCatda
 * @Description: HttpRequest函数注册
 * @FilePath: \Momizi\Internal\Plugin\JavaScriptV8\Tools\HttpRequest\HttpRequest.go
 */
package HttpRequest

import (
	"net/http"

	"rogchap.com/v8go"
)

func Register(Isolate *v8go.Isolate, Context *v8go.Context) *v8go.Object {
	HttpRequest, _ := v8go.NewObjectTemplate(Isolate)

	Get := Get(Isolate, Context)
	HttpRequest.Set("Get", Get)

	ConsoleObject, _ := HttpRequest.NewInstance(Context)
	return ConsoleObject
}

/**
 * @description: Http返回体
 */
type HttpResponse struct {
	Status     string // e.g. "200 OK"
	StatusCode int    // e.g. 200
	Proto      string // e.g. "HTTP/1.0"
	ProtoMajor int    // e.g. 1
	ProtoMinor int    // e.g. 0

	// Header maps header keys to values. If the response had multiple
	// headers with the same key, they may be concatenated, with comma
	// delimiters.  (RFC 7230, section 3.2.2 requires that multiple headers
	// be semantically equivalent to a comma-delimited sequence.) When
	// Header values are duplicated by other fields in this struct (e.g.,
	// ContentLength, TransferEncoding, Trailer), the field values are
	// authoritative.
	//
	// Keys in the map are canonicalized (see CanonicalHeaderKey).
	Header map[string][]string

	// ContentLength records the length of the associated content. The
	// value -1 indicates that the length is unknown. Unless Request.Method
	// is "HEAD", values >= 0 indicate that the given number of bytes may
	// be read from Body.
	ContentLength int64

	// Contains transfer encodings from outer-most to inner-most. Value is
	// nil, means that "identity" encoding is used.
	TransferEncoding []string

	// Close records whether the header directed that the connection be
	// closed after reading Body. The value is advice for clients: neither
	// ReadResponse nor Response.Write ever closes a connection.
	Close bool

	// Uncompressed reports whether the response was sent compressed but
	// was decompressed by the http package. When true, reading from
	// Body yields the uncompressed content instead of the compressed
	// content actually set from the server, ContentLength is set to -1,
	// and the "Content-Length" and "Content-Encoding" fields are deleted
	// from the responseHeader. To get the original response from
	// the server, set Transport.DisableCompression to true.
	Uncompressed bool

	// Trailer maps trailer keys to values in the same
	// format as Header.
	//
	// The Trailer initially contains only nil values, one for
	// each key specified in the server's "Trailer" header
	// value. Those values are not added to Header.
	//
	// Trailer must not be accessed concurrently with Read calls
	// on the Body.
	//
	// After Body.Read has returned io.EOF, Trailer will contain
	// any trailer values sent by the server.
	Trailer map[string][]string
}

/**
 * @description: 指针Http返回结构体转换
 * @param {*http.Response} HttpResponseValue
 * @return {*}
 */
func PointerHttpResponseToHttpResponse(HttpResponseValue *http.Response) HttpResponse {
	HttpResponse := HttpResponse{
		Status:           HttpResponseValue.Status,
		StatusCode:       HttpResponseValue.StatusCode,
		Proto:            HttpResponseValue.Proto,
		ProtoMajor:       HttpResponseValue.ProtoMajor,
		ProtoMinor:       HttpResponseValue.ProtoMinor,
		Header:           HttpResponseValue.Header,
		ContentLength:    HttpResponseValue.ContentLength,
		TransferEncoding: HttpResponseValue.TransferEncoding,
		Close:            HttpResponseValue.Close,
		Uncompressed:     HttpResponseValue.Uncompressed,
		Trailer:          HttpResponseValue.Trailer,
	}
	return HttpResponse
}
