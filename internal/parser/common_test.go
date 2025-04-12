package parser

import (
	"net/http"
	"net/http/httptest"
)

const testHtmlPage = `
<html>
	<body>
	<h1>test page</h1>
	<img src="cat.png">
	<p>
		<img src="big_cat.png">
	</p>
	</body>
</html>
`

func testHtmlPageServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testHtmlPage))
	}))
}
