package wework_finance_sdk

/*
#cgo CFLAGS: -I${SRCDIR}/C_sdk
#cgo darwin LDFLAGS: -L${SRCDIR}/C_sdk -lWeWorkFinanceSdk_C -Wl,-rpath,${SRCDIR}/C_sdk
#cgo linux,amd64 LDFLAGS: -L${SRCDIR}/C_sdk -lWeWorkFinanceSdk_C -Wl,-rpath,${SRCDIR}/C_sdk
#cgo linux,arm64 LDFLAGS: -L${SRCDIR}/C_sdk -lWeWorkFinanceSdk_C -Wl,-rpath,${SRCDIR}/C_sdk
#include "WeWorkFinanceSdk_C.h"
#include <stdlib.h>
#include <string.h>
static inline char* copyCStr(const char* src) {
    if (src == NULL) return NULL;
    size_t len = strlen(src);
    char* dst = (char*)malloc(len + 1);
    memcpy(dst, src, len);
    dst[len] = '\0';
    return dst;
}
*/
import "C"
import (
    "errors"
    "unsafe"
)

type SDK struct {
    sdk *C.WeWorkFinanceSdk_t
}

func NewSDK(corpID, secret string) (*SDK, error) {
    cCorpID := C.CString(corpID)
    defer C.free(unsafe.Pointer(cCorpID))
    cSecret := C.CString(secret)
    defer C.free(unsafe.Pointer(cSecret))

    instance := C.NewSdk()
    if instance == nil {
        return nil, errors.New("failed to create sdk instance")
    }

    ret := C.Init(instance, cCorpID, cSecret)
    if ret != 0 {
        C.DestroySdk(instance)
        return nil, errors.New("init sdk failed, code:" + itoa(int(ret)))
    }
    return &SDK{sdk: instance}, nil
}

// Close must be called to free underlying C resources
func (s *SDK) Close() {
    if s.sdk != nil {
        C.DestroySdk(s.sdk)
        s.sdk = nil
    }
}

type ChatDataSlice struct {
    cSlice *C.Slice_t
}

func (s *SDK) GetChatData(seq uint64, limit uint, proxy, passwd string, timeout int) ([]byte, error) {
    cProxy := C.CString(proxy)
    defer C.free(unsafe.Pointer(cProxy))
    cPass := C.CString(passwd)
    defer C.free(unsafe.Pointer(cPass))

    slice := C.NewSlice()
    if slice == nil {
        return nil, errors.New("NewSlice failed")
    }
    defer C.FreeSlice(slice)

    ret := C.GetChatData(s.sdk, C.ulonglong(seq), C.uint(limit), cProxy, cPass, C.int(timeout), slice)
    if ret != 0 {
        return nil, errors.New("GetChatData failed code:" + itoa(int(ret)))
    }

    buf := C.GetContentFromSlice(slice)
    length := C.GetSliceLen(slice)
    if buf == nil || length == 0 {
        return nil, nil
    }
    data := C.GoBytes(unsafe.Pointer(buf), C.int(length))
    return data, nil
}

func (s *SDK) DecryptData(encKey, encMsg string) ([]byte, error) {
    cKey := C.CString(encKey)
    defer C.free(unsafe.Pointer(cKey))
    cMsg := C.CString(encMsg)
    defer C.free(unsafe.Pointer(cMsg))

    slice := C.NewSlice()
    if slice == nil {
        return nil, errors.New("NewSlice failed")
    }
    defer C.FreeSlice(slice)

    ret := C.DecryptData(cKey, cMsg, slice)
    if ret != 0 {
        return nil, errors.New("DecryptData failed code:" + itoa(int(ret)))
    }
    buf := C.GetContentFromSlice(slice)
    length := C.GetSliceLen(slice)
    data := C.GoBytes(unsafe.Pointer(buf), C.int(length))
    return data, nil
}

type MediaData struct {
    OutIndex string
    Data     []byte
    Finish   bool
}

func (s *SDK) GetMediaData(indexBuf, sdkFileID, proxy, passwd string, timeout int) (*MediaData, error) {
    cIndex := C.CString(indexBuf)
    defer C.free(unsafe.Pointer(cIndex))
    cFile := C.CString(sdkFileID)
    defer C.free(unsafe.Pointer(cFile))
    cProxy := C.CString(proxy)
    defer C.free(unsafe.Pointer(cProxy))
    cPass := C.CString(passwd)
    defer C.free(unsafe.Pointer(cPass))

    m := C.NewMediaData()
    if m == nil {
        return nil, errors.New("NewMediaData failed")
    }
    defer C.FreeMediaData(m)

    ret := C.GetMediaData(s.sdk, cIndex, cFile, cProxy, cPass, C.int(timeout), m)
    if ret != 0 {
        return nil, errors.New("GetMediaData failed code:" + itoa(int(ret)))
    }

    outIndex := C.GoString(C.GetOutIndexBuf(m))
    dataLen := int(C.GetDataLen(m))
    var data []byte
    if dataLen > 0 {
        data = C.GoBytes(unsafe.Pointer(C.GetData(m)), C.int(dataLen))
    }
    finish := C.IsMediaDataFinish(m) != 0
    return &MediaData{OutIndex: outIndex, Data: data, Finish: finish}, nil
}

// helper itoa without strconv to avoid pulling heavy package in tiny sdk
func itoa(i int) string {
    if i == 0 {
        return "0"
    }
    neg := false
    if i < 0 {
        neg = true
        i = -i
    }
    var buf [20]byte
    pos := len(buf)
    for i > 0 {
        pos--
        buf[pos] = byte('0' + i%10)
        i /= 10
    }
    if neg {
        pos--
        buf[pos] = '-'
    }
    return string(buf[pos:])
}