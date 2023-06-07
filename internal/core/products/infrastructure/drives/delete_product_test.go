package drives

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/erik-sostenes/products-api/internal/core/products/business/domain"
	"github.com/erik-sostenes/products-api/internal/core/products/business/services"
	"github.com/erik-sostenes/products-api/internal/core/products/infrastructure/driven/db"
	"github.com/erik-sostenes/products-api/internal/shared/domain/bus/command"
	"github.com/labstack/echo/v4"
)

func TestProductHandler_Delete(t *testing.T) {
	type HandlerFunc func() (echo.HandlerFunc, error)

	tsc := map[string]struct {
		HandlerFunc
		*http.Request
		expectedStatusCode int
	}{
		"Given an existing valid record, it is delete, a status code 201 is expected": {
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

				deleteProductCommandHandler := services.NewDeleteProductCommandHandler(mock)
				deleteProductCmdBus := make(command.CommandBus[services.DeleteProductCommand])
				if err := deleteProductCmdBus.Record(services.DeleteProductCommand{}, deleteProductCommandHandler); err != nil {
					return nil, err
				}
				return DeleteProduct(deleteProductCmdBus), nil
			},
			Request:            httptest.NewRequest(http.MethodDelete, "/api/v1/products/delete/4448b491-153b-4161-92a9-ecd12f541a28", http.NoBody),
			expectedStatusCode: http.StatusOK,
		},
		"Given a non-existent record, will not be found, a status code 404 is expected": {
			HandlerFunc: func() (echo.HandlerFunc, error) {
				deleteProductCommandHandler := services.NewDeleteProductCommandHandler(db.NewMockProductStorer())
				deleteProductCmdBus := make(command.CommandBus[services.DeleteProductCommand])
				if err := deleteProductCmdBus.Record(services.DeleteProductCommand{}, deleteProductCommandHandler); err != nil {
					return nil, err
				}
				return DeleteProduct(deleteProductCmdBus), nil
			},
			Request:            httptest.NewRequest(http.MethodDelete, "/api/v1/products/delete/4448b491-153b-4161-92a9-ecd12f541a28", http.NoBody),
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

			e.DELETE("/api/v1/products/delete/:id", handlerFunc)

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
