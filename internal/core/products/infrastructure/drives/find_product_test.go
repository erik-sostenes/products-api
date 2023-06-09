package drives

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/erik-sostenes/products-api/internal/core/products/business/domain"
	"github.com/erik-sostenes/products-api/internal/core/products/business/services"
	"github.com/erik-sostenes/products-api/internal/core/products/infrastructure/driven/db"
	"github.com/erik-sostenes/products-api/internal/shared/domain/bus/query"
	"github.com/labstack/echo/v4"
)

func TestProductHandler_FindById(t *testing.T) {
	type HandlerFunc func() (echo.HandlerFunc, error)

	tsc := map[string]struct {
		HandlerFunc
		*http.Request
		expectedStatusCode int
	}{
		"Given a valid http request, a code 200 status with data is expected": {
			HandlerFunc: func() (echo.HandlerFunc, error) {
				mock := db.NewMockProductStorer()
				product := domain.Product{
					Id:        "4448b491-153b-4161-92a9-ecd12f541a28",
					Title:     "Celular Samsung Galaxy A54",
					Price:     5499,
					ImageURL:  "https://http2.mlstatic.com/D_NQ_NP_987811-MLM54517443931_032023-O.webp",
					Offer:     true,
					Available: true,
					Rating:    4.8,
					Sales: domain.Sales{
						Days:   60,
						Amount: 30000,
					},
				}

				if err := mock.Save(context.Background(), product.Id, product); err != nil {
					return nil, err
				}

				finderProductQueryHandler := services.NewFinderProductQueryHandler(mock)

				queryBus := make(query.QueryBus[services.FindProductQuery, services.ProductResponse])
				if err := queryBus.Record(services.FindProductQuery{}, finderProductQueryHandler); err != nil {
					return nil, err
				}

				return FindProduct(&queryBus), nil
			},
			Request:            httptest.NewRequest(http.MethodGet, "/api/v1/products/get/?id=4448b491-153b-4161-92a9-ecd12f541a28", http.NoBody),
			expectedStatusCode: http.StatusOK,
		},
		"Given a valid http request, a status code 404 is expected": {
			HandlerFunc: func() (echo.HandlerFunc, error) {
				finderProductQueryHandler := services.NewFinderProductQueryHandler(db.NewMockProductStorer())

				queryBus := make(query.QueryBus[services.FindProductQuery, services.ProductResponse])
				if err := queryBus.Record(services.FindProductQuery{}, finderProductQueryHandler); err != nil {
					return nil, err
				}

				return FindProduct(&queryBus), nil
			},
			Request:            httptest.NewRequest(http.MethodGet, "/api/v1/products/get/?id=4448b491-153b-4161-92a9-ecd12f541a28", http.NoBody),
			expectedStatusCode: http.StatusNotFound,
		},
	}

	for name, ts := range tsc {
		t.Run(name, func(t *testing.T) {
			e := echo.New()

			handlerFunc, err := ts.HandlerFunc()
			if err != nil {
				t.Skip(err)
			}

			e.GET("/api/v1/products/get/", handlerFunc)

			resp := httptest.NewRecorder()
			req := ts.Request
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			e.ServeHTTP(resp, req)

			if resp.Code != ts.expectedStatusCode {
				t.Log(resp.Body.String())
				t.Errorf("status code was expected %d, but it was obtained %d", ts.expectedStatusCode, resp.Code)
			}
		})
	}
}
