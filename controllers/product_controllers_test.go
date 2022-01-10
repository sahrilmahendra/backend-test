package controllers

import (
	"bytes"
	"encoding/json"

	"erajaya/config"
	"erajaya/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// inisialisasi echo
func InitEchoAPI() *echo.Echo {
	config.InitDBTest()
	e := echo.New()

	return e
}

type ResponseAllData struct {
	Message string
	Data    []models.Product
}

type ResponseData struct {
	Message string
	Data    models.Product
}

type ResponseNonData struct {
	Message string
}

var (
	mock_data_product = models.Product{
		Product_Name:        "Dummy",
		Product_Price:       1000,
		Product_Description: "lorem",
		Product_Qty:         2,
	}
	mock_data_product_name_empty = models.Product{
		Product_Name:        "",
		Product_Price:       1000,
		Product_Description: "lorem",
		Product_Qty:         2,
	}
	mock_data_product_price_empty = models.Product{
		Product_Name:        "Dummy",
		Product_Description: "lorem",
		Product_Qty:         2,
	}
	mock_data_product_price_zero = models.Product{
		Product_Name:        "Dummy",
		Product_Price:       -2,
		Product_Description: "lorem",
		Product_Qty:         2,
	}
	mock_data_product_description_empty = models.Product{
		Product_Name:        "Dummy",
		Product_Price:       1000,
		Product_Description: "",
		Product_Qty:         2,
	}
	mock_data_product_qty_empty = models.Product{
		Product_Name:        "Dummy",
		Product_Price:       1000,
		Product_Description: "lorem",
	}
	mock_data_product_qty_zero = models.Product{
		Product_Name:        "Dummy",
		Product_Price:       1000,
		Product_Description: "lorem",
		Product_Qty:         -4,
	}
)

func TestGetAllProductSuccess(t *testing.T) {
	e := InitEchoAPI()
	config.DB.Save(&mock_data_product)
	req := httptest.NewRequest(http.MethodGet, "/products", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	contex := e.NewContext(req, res)

	if assert.NoError(t, GetAllProductsController(contex)) {
		var product ResponseAllData
		body := res.Body.String()
		if err := json.Unmarshal([]byte(body), &product); err != nil {
			assert.Error(t, err, "error")
		}
		assert.Equal(t, "Success Operation", product.Message)
		assert.Equal(t, "Dummy", product.Data[0].Product_Name)
	}
}

func TestGetAllProductFailure(t *testing.T) {
	e := InitEchoAPI()
	t.Run("data_not_found", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/products", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)

		if assert.NoError(t, GetAllProductsController(context)) {
			body := res.Body.String()
			var product ResponseNonData
			err := json.Unmarshal([]byte(body), &product)

			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, "Data Not Found", product.Message)
		}
	})
	t.Run("bad_request", func(t *testing.T) {
		config.DB.Migrator().DropTable(models.Product{})
		config.DB.Save(&mock_data_product)
		req := httptest.NewRequest(http.MethodGet, "/products", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)

		if assert.NoError(t, GetAllProductsController(context)) {
			body := res.Body.String()
			var product ResponseNonData
			err := json.Unmarshal([]byte(body), &product)

			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, "Bad Request", product.Message)
		}
	})
}

func TestGetAllByCreatedDescProductSuccess(t *testing.T) {
	e := InitEchoAPI()
	config.DB.Save(&mock_data_product)
	req := httptest.NewRequest(http.MethodGet, "/products/newest", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	contex := e.NewContext(req, res)

	if assert.NoError(t, GetAllProductsByCreatedDescController(contex)) {
		var product ResponseAllData
		body := res.Body.String()
		if err := json.Unmarshal([]byte(body), &product); err != nil {
			assert.Error(t, err, "error")
		}
		assert.Equal(t, "Success Operation", product.Message)
		assert.Equal(t, "Dummy", product.Data[0].Product_Name)
	}
}

func TestGetAllProductByCreatedDescFailure(t *testing.T) {
	e := InitEchoAPI()
	t.Run("data_not_found", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/products/newest", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)

		if assert.NoError(t, GetAllProductsByCreatedDescController(context)) {
			body := res.Body.String()
			var product ResponseNonData
			err := json.Unmarshal([]byte(body), &product)

			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, "Data Not Found", product.Message)
		}
	})
	t.Run("bad_request", func(t *testing.T) {
		config.DB.Migrator().DropTable(models.Product{})
		config.DB.Save(&mock_data_product)
		req := httptest.NewRequest(http.MethodGet, "/products/newest", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)

		if assert.NoError(t, GetAllProductsByCreatedDescController(context)) {
			body := res.Body.String()
			var product ResponseNonData
			err := json.Unmarshal([]byte(body), &product)

			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, "Bad Request", product.Message)
		}
	})
}

func TestGetAllProductsByNameAscControllerSuccess(t *testing.T) {
	e := InitEchoAPI()
	config.DB.Save(&mock_data_product)
	req := httptest.NewRequest(http.MethodGet, "/products/name/asc", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	contex := e.NewContext(req, res)

	if assert.NoError(t, GetAllProductsByNameAscController(contex)) {
		var product ResponseAllData
		body := res.Body.String()
		if err := json.Unmarshal([]byte(body), &product); err != nil {
			assert.Error(t, err, "error")
		}
		assert.Equal(t, "Success Operation", product.Message)
		assert.Equal(t, "Dummy", product.Data[0].Product_Name)
	}
}

func TestGetAllProductsByNameAscControllerFailure(t *testing.T) {
	e := InitEchoAPI()
	t.Run("data_not_found", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/products/name/asc", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)

		if assert.NoError(t, GetAllProductsByNameAscController(context)) {
			body := res.Body.String()
			var product ResponseNonData
			err := json.Unmarshal([]byte(body), &product)

			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, "Data Not Found", product.Message)
		}
	})
	t.Run("bad_request", func(t *testing.T) {
		config.DB.Migrator().DropTable(models.Product{})
		config.DB.Save(&mock_data_product)
		req := httptest.NewRequest(http.MethodGet, "/products/name/asc", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)

		if assert.NoError(t, GetAllProductsByNameAscController(context)) {
			body := res.Body.String()
			var product ResponseNonData
			err := json.Unmarshal([]byte(body), &product)

			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, "Bad Request", product.Message)
		}
	})
}

func TestGetAllProductsByNameDescControllerSuccess(t *testing.T) {
	e := InitEchoAPI()
	config.DB.Save(&mock_data_product)
	req := httptest.NewRequest(http.MethodGet, "/products/name/desc", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	contex := e.NewContext(req, res)

	if assert.NoError(t, GetAllProductsByNameDescController(contex)) {
		var product ResponseAllData
		body := res.Body.String()
		if err := json.Unmarshal([]byte(body), &product); err != nil {
			assert.Error(t, err, "error")
		}
		assert.Equal(t, "Success Operation", product.Message)
		assert.Equal(t, "Dummy", product.Data[0].Product_Name)
	}
}

func TestGetAllProductsByNameDescControllerFailure(t *testing.T) {
	e := InitEchoAPI()
	t.Run("data_not_found", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/products/name/desc", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)

		if assert.NoError(t, GetAllProductsByNameDescController(context)) {
			body := res.Body.String()
			var product ResponseNonData
			err := json.Unmarshal([]byte(body), &product)

			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, "Data Not Found", product.Message)
		}
	})
	t.Run("bad_request", func(t *testing.T) {
		config.DB.Migrator().DropTable(models.Product{})
		config.DB.Save(&mock_data_product)
		req := httptest.NewRequest(http.MethodGet, "/products/name/desc", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)

		if assert.NoError(t, GetAllProductsByNameDescController(context)) {
			body := res.Body.String()
			var product ResponseNonData
			err := json.Unmarshal([]byte(body), &product)

			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, "Bad Request", product.Message)
		}
	})
}

func TestGetAllProductsByPriceAscControllerSuccess(t *testing.T) {
	e := InitEchoAPI()
	config.DB.Save(&mock_data_product)
	req := httptest.NewRequest(http.MethodGet, "/products/price/low", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	contex := e.NewContext(req, res)

	if assert.NoError(t, GetAllProductsByPriceAscController(contex)) {
		var product ResponseAllData
		body := res.Body.String()
		if err := json.Unmarshal([]byte(body), &product); err != nil {
			assert.Error(t, err, "error")
		}
		assert.Equal(t, "Success Operation", product.Message)
		assert.Equal(t, "Dummy", product.Data[0].Product_Name)
	}
}

func TestGetAllProductsByPriceAscControllerFailure(t *testing.T) {
	e := InitEchoAPI()
	t.Run("data_not_found", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/products/price/low", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)

		if assert.NoError(t, GetAllProductsByPriceAscController(context)) {
			body := res.Body.String()
			var product ResponseNonData
			err := json.Unmarshal([]byte(body), &product)

			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, "Data Not Found", product.Message)
		}
	})
	t.Run("bad_request", func(t *testing.T) {
		config.DB.Migrator().DropTable(models.Product{})
		config.DB.Save(&mock_data_product)
		req := httptest.NewRequest(http.MethodGet, "/products/price/low", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)

		if assert.NoError(t, GetAllProductsByPriceAscController(context)) {
			body := res.Body.String()
			var product ResponseNonData
			err := json.Unmarshal([]byte(body), &product)

			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, "Bad Request", product.Message)
		}
	})
}

func TestGetAllProductsByPriceDescControllerSuccess(t *testing.T) {
	e := InitEchoAPI()
	config.DB.Save(&mock_data_product)
	req := httptest.NewRequest(http.MethodGet, "/products/price/expensive", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	contex := e.NewContext(req, res)

	if assert.NoError(t, GetAllProductsByPriceDescController(contex)) {
		var product ResponseAllData
		body := res.Body.String()
		if err := json.Unmarshal([]byte(body), &product); err != nil {
			assert.Error(t, err, "error")
		}
		assert.Equal(t, "Success Operation", product.Message)
		assert.Equal(t, "Dummy", product.Data[0].Product_Name)
	}
}

func TestGetAllProductsByPriceDescControllerFailure(t *testing.T) {
	e := InitEchoAPI()
	t.Run("data_not_found", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/products/price/expensive", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)

		if assert.NoError(t, GetAllProductsByPriceDescController(context)) {
			body := res.Body.String()
			var product ResponseNonData
			err := json.Unmarshal([]byte(body), &product)

			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, "Data Not Found", product.Message)
		}
	})
	t.Run("bad_request", func(t *testing.T) {
		config.DB.Migrator().DropTable(models.Product{})
		config.DB.Save(&mock_data_product)
		req := httptest.NewRequest(http.MethodGet, "/products/price/expensive", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)

		if assert.NoError(t, GetAllProductsByPriceDescController(context)) {
			body := res.Body.String()
			var product ResponseNonData
			err := json.Unmarshal([]byte(body), &product)

			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, "Bad Request", product.Message)
		}
	})
}

func TestAddProductControllersSuccess(t *testing.T) {
	e := InitEchoAPI()

	body, err := json.Marshal(mock_data_product)
	if err != nil {
		t.Error(t, err, "error")
	}

	req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	if assert.NoError(t, AddProductController(context)) {
		body_response := res.Body.String()
		var product ResponseNonData
		json.Unmarshal([]byte(body_response), &product)
		assert.Equal(t, "Success Operation", product.Message)
	}
}

func TestAddProductFailure(t *testing.T) {
	e := InitEchoAPI()
	t.Run("product name empty", func(t *testing.T) {
		body, err := json.Marshal(mock_data_product_name_empty)
		if err != nil {
			t.Error(t, err, "error")
		}
		req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)

		if assert.NoError(t, AddProductController(context)) {
			body := res.Body.String()
			var product ResponseNonData
			err := json.Unmarshal([]byte(body), &product)

			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, "Product's name can't be empty", product.Message)
		}
	})
	t.Run("product price empty", func(t *testing.T) {
		body, err := json.Marshal(mock_data_product_price_empty)
		if err != nil {
			t.Error(t, err, "error")
		}
		req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)

		if assert.NoError(t, AddProductController(context)) {
			body := res.Body.String()
			var product ResponseNonData
			err := json.Unmarshal([]byte(body), &product)

			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, "Product's price can't be empty", product.Message)
		}
	})
	t.Run("product price greater than 0", func(t *testing.T) {
		body, err := json.Marshal(mock_data_product_price_zero)
		if err != nil {
			t.Error(t, err, "error")
		}
		req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)

		if assert.NoError(t, AddProductController(context)) {
			body := res.Body.String()
			var product ResponseNonData
			err := json.Unmarshal([]byte(body), &product)

			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, "Product's price must be greater than 0", product.Message)
		}
	})
	t.Run("product description empty", func(t *testing.T) {
		body, err := json.Marshal(mock_data_product_description_empty)
		if err != nil {
			t.Error(t, err, "error")
		}
		req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)

		if assert.NoError(t, AddProductController(context)) {
			body := res.Body.String()
			var product ResponseNonData
			err := json.Unmarshal([]byte(body), &product)

			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, "Product's description can't be empty", product.Message)
		}
	})
	t.Run("product qty empty", func(t *testing.T) {
		body, err := json.Marshal(mock_data_product_qty_empty)
		if err != nil {
			t.Error(t, err, "error")
		}
		req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)

		if assert.NoError(t, AddProductController(context)) {
			body := res.Body.String()
			var product ResponseNonData
			err := json.Unmarshal([]byte(body), &product)

			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, "Product's quantity can't be empty", product.Message)
		}
	})
	t.Run("product qty greater than 0", func(t *testing.T) {
		body, err := json.Marshal(mock_data_product_qty_zero)
		if err != nil {
			t.Error(t, err, "error")
		}
		req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)

		if assert.NoError(t, AddProductController(context)) {
			body := res.Body.String()
			var product ResponseNonData
			err := json.Unmarshal([]byte(body), &product)

			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, "Product's quantity must be greater than 0", product.Message)
		}
	})
	t.Run("internal server error", func(t *testing.T) {
		body, err := json.Marshal(mock_data_product)
		if err != nil {
			t.Error(t, err, "error")
		}
		req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		config.DB.Migrator().DropTable(&models.Product{})

		if assert.NoError(t, AddProductController(context)) {
			body := res.Body.String()
			var product ResponseNonData
			err := json.Unmarshal([]byte(body), &product)

			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, "Internal Server Error", product.Message)
		}
	})
}
