package wework_finance_sdk

import (
	"testing"
)

func TestNewSDK(t *testing.T) {
	// Test with invalid credentials
	_, err := NewSDK("invalid_corp_id", "invalid_secret")
	if err != nil {
		t.Error("NewSDK error: ", err)
	}
}

func TestGetChatData(t *testing.T) {
	// Skip if no valid SDK instance
	sdk, err := NewSDK("test_corp_id", "test_secret")
	if err != nil {
		t.Error("NewSDK error: ", err)
	}
	defer sdk.Close()

	if sdk == nil {
		t.Skip()
	}

	// Test GetChatData
	data, err := sdk.GetChatData(0, 1, "", "", 30)
	if err != nil {
		t.Error("GetChatData error: ", err)
	}
	if data != nil {
		t.Log("GetChatData data: ", data)
	}
}

func TestDecryptData(t *testing.T) {
	// Skip if no valid SDK instance
	sdk, err := NewSDK("test_corp_id", "test_secret")
	if err != nil {
		t.Error("NewSDK error: ", err)
	}
	defer sdk.Close()

	if sdk == nil {
		t.Skip()
	}

	// Test with invalid input
	data, err := sdk.DecryptData("invalid_key", "invalid_message")
	if err != nil {
		t.Error("DecryptData error: ", err)
	}
	if data != nil {
		t.Log("DecryptData data: ", data)
	}
}

func TestGetMediaData(t *testing.T) {
	// Skip if no valid SDK instance
	sdk, err := NewSDK("test_corp_id", "test_secret")
	if err != nil {
		t.Error("NewSDK error: ", err)
	}
	defer sdk.Close()

	if sdk == nil {
		t.Skip()
	}

	// Test with invalid file ID
	media, err := sdk.GetMediaData("test_index", "invalid_file_id", "", "", 0)
	if err != nil {
		t.Error("GetMediaData error: ", err)
	}
	if media != nil {
		t.Log("GetMediaData media: ", media)
	}
}

func TestSDKClose(t *testing.T) {
	sdk, err := NewSDK("test_corp_id", "test_secret")
	if err != nil {
		t.Error("NewSDK error: ", err)
	}
	defer sdk.Close()

	if sdk == nil {
		t.Skip()
	}

	// Test Close method
	sdk.Close()
	if sdk.sdk != nil {
		t.Error("SDK instance not properly closed")
	}

	// Test double close
	sdk.Close() // Should not panic
}
