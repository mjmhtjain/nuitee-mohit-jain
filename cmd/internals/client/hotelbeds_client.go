package client

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/mjmhtjain/nuitee-mohit-jain/cmd/internals/util"
)

type HotelBedsClient interface {
	SearchHotels(request []byte) ([]byte, error)
}

type HotelBedsClientImpl struct {
	baseURL    string
	httpClient *http.Client
	apiKey     string
	apiSecret  string
}

func NewHotelBedsClient() HotelBedsClient {
	return &HotelBedsClientImpl{
		baseURL:   os.Getenv("HOTEL_BEDS_BASE_URL"),
		apiKey:    os.Getenv("HOTEL_BEDS_API_KEY"),
		apiSecret: os.Getenv("HOTEL_BEDS_SECRET"),
		httpClient: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}

func (c *HotelBedsClientImpl) SearchHotels(request []byte) (response []byte, err error) {

	// Create the request URL with the base URL
	url := fmt.Sprintf("%s/hotel-api/1.0/hotels", c.baseURL)

	// Create new POST request with the JSON body
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(request))
	if err != nil {
		return response, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	err = c.setHeaders(req)
	if err != nil {
		return response, fmt.Errorf("failed to create request: %w", err)
	}

	// Make the request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return response, fmt.Errorf("failed to make API request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return response, fmt.Errorf("API returned non-200 status code: %d", resp.StatusCode)
	}

	var reader io.ReadCloser
	if resp.Header.Get("Content-Encoding") == "gzip" {
		reader, err = gzip.NewReader(resp.Body)
		if err != nil {
			return response, fmt.Errorf("failed to create gzip reader: %w", err)
		}
		defer reader.Close()
	} else {
		reader = resp.Body
	}

	// Read the response body into a string
	response, err = io.ReadAll(reader)
	if err != nil {
		return response, fmt.Errorf("failed to read response body: %w", err)
	}

	return response, nil
}

func (c *HotelBedsClientImpl) setHeaders(req *http.Request) error {
	signature, err := c.generateSignature()
	if err != nil {
		return err
	}

	req.Header.Set(util.HeaderContentType, util.ValueApplicationJSON)
	req.Header.Set(util.HeaderApiKey, c.apiKey)
	req.Header.Set(util.HeaderSignature, signature)
	req.Header.Set(util.HeaderAccept, util.ValueApplicationJSON)
	req.Header.Set(util.HeaderAcceptEncoding, util.ValueGzip)

	return nil
}

func (c *HotelBedsClientImpl) generateSignature() (string, error) {
	// Check if environment variables are set
	if c.apiKey == "" || c.apiSecret == "" {
		return "", fmt.Errorf("environment variables 'HOTEL_BEDS_API_KEY' and 'HOTEL_BEDS_SECRET' are required")
	}

	// Generate the current UTC timestamp in seconds
	timestamp := time.Now().Unix()

	// Assemble the string to hash
	assemble := fmt.Sprintf("%s%s%d", c.apiKey, c.apiSecret, timestamp)

	// Perform SHA-256 encryption
	hash := sha256.New()
	hash.Write([]byte(assemble))
	encryption := hex.EncodeToString(hash.Sum(nil))

	return encryption, nil
}
