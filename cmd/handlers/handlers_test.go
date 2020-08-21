package handlers_test

import (
	"fmt"
	"github.com/cobbinma/kanye-west-api/cmd/handlers"
	"github.com/cobbinma/kanye-west-api/cmd/handlers/mock"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
)

var _ = Describe("Handlers", func() {
	var (
		ctrl   *gomock.Controller
		client *mock.MockClient
	)
	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		client = mock.NewMockClient(ctrl)
	})
	When("Kanye Handler is called", func() {
		When("a valid response is received", func() {
			quote := "Perhaps I should have been more like water today"
			body := fmt.Sprintf(`{"quote":"%s"}`, quote)
			resp := &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(strings.NewReader(body)),
			}
			BeforeEach(func() {
				client.EXPECT().Get(gomock.Eq("https://api.kanye.rest/")).Return(resp, nil)
			})

			It("should return HTTP status of OK and correct quote", func() {
				req, _ := http.NewRequest("GET", "/kanye", nil)
				rr := httptest.NewRecorder()
				handler := http.HandlerFunc(handlers.Kanye(client))
				handler.ServeHTTP(rr, req)

				Expect(rr.Code).To(Equal(http.StatusOK))
				Expect(rr.Body.String()).To(Equal(quote))
			})
		})
	})
})
