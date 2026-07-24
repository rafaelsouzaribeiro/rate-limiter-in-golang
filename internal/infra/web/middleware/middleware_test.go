package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/rafaelsouzaribeiro/rate-limiter-in-golang/internal/usecase"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUsecase struct {
	mock.Mock
}

func (m *MockUsecase) IsBlocked(key string) (bool, error) {
	args := m.Called(key)
	return args.Bool(0), args.Error(1)
}

func (m *MockUsecase) IncreaseRequest(ip string, duration time.Duration) (int64, error) {
	args := m.Called(ip, duration)
	return int64(args.Int(0)), args.Error(1)
}

func (m *MockUsecase) Block(ip string, duration time.Duration) error {
	args := m.Called(ip, duration)
	return args.Error(0)
}

func setupViper() {
	viper.Set("MAX_REQUEST", 5)
	viper.Set("BLOCK_TIME", "1m")
	viper.Set("IP_LIMIT", "1s")
	viper.Set("TOKEN_LIMIT", "1s")

}

func TestRateLimiter_AllowsRequestWithinLimit(t *testing.T) {
	setupViper()
	mockUC := new(MockUsecase)
	mockUC.On("IsBlocked", mock.Anything).Return(false, nil)
	mockUC.On("IncreaseRequest", mock.Anything, mock.Anything).Return(1, nil)
	usecease := usecase.NewUsecase(mockUC)

	h := &Middleware{usecase: usecease}

	nextCalled := false
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nextCalled = true
		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.RemoteAddr = "127.0.0.1:1234"
	rec := httptest.NewRecorder()

	h.RateLimiter(next).ServeHTTP(rec, req)

	assert.True(t, nextCalled)
	assert.Equal(t, http.StatusOK, rec.Code)
	mockUC.AssertExpectations(t)
}

func TestRateLimiter_BlocksWhenIpIsBlocked(t *testing.T) {
	setupViper()
	mockUC := new(MockUsecase)
	mockUC.On("IsBlocked", mock.Anything).Return(true, nil)

	usecease := usecase.NewUsecase(mockUC)
	h := &Middleware{usecase: usecease}

	nextCalled := false
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nextCalled = true
	})

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.RemoteAddr = "127.0.0.1:1234"
	rec := httptest.NewRecorder()

	h.RateLimiter(next).ServeHTTP(rec, req)

	assert.False(t, nextCalled)
	assert.Equal(t, http.StatusTooManyRequests, rec.Code)
}

func TestRateLimiter_BlocksAfterExceedingMaxRequests(t *testing.T) {
	setupViper()
	viper.Set("MAX_REQUEST", 1)

	mockUC := new(MockUsecase)
	mockUC.On("IsBlocked", mock.Anything).Return(false, nil)
	mockUC.On("IncreaseRequest", mock.Anything, mock.Anything).Return(2, nil)
	mockUC.On("Block", mock.Anything, mock.Anything).Return(nil)

	usecease := usecase.NewUsecase(mockUC)
	h := &Middleware{usecase: usecease}

	nextCalled := false
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nextCalled = true
	})

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.RemoteAddr = "127.0.0.1:1234"
	rec := httptest.NewRecorder()

	h.RateLimiter(next).ServeHTTP(rec, req)

	assert.False(t, nextCalled)
	assert.Equal(t, http.StatusTooManyRequests, rec.Code)
	mockUC.AssertCalled(t, "Block", mock.Anything, mock.Anything)
}

func TestRateLimiter_UsesApiKeyWhenPresent(t *testing.T) {
	setupViper()
	mockUC := new(MockUsecase)
	mockUC.On("IsBlocked", "my-token").Return(false, nil)
	mockUC.On("IncreaseRequest", "my-token", mock.Anything).Return(1, nil)

	usecease := usecase.NewUsecase(mockUC)
	h := &Middleware{usecase: usecease}

	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("API_KEY", "my-token")
	rec := httptest.NewRecorder()

	h.RateLimiter(next).ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUC.AssertCalled(t, "IsBlocked", "my-token")
}
