// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "MeasurementInfoItem.h"
import "C"

import (
	"encoding/binary"
	"fmt"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-ies"
	"unsafe"
)

func xerEncodeMeasurementInfoItem(measurementInfoItem *e2sm_kpm_v2.MeasurementInfoItem) ([]byte, error) {
	measurementInfoItemCP, err := newMeasurementInfoItem(measurementInfoItem)
	if err != nil {
		return nil, fmt.Errorf("xerEncodeMeasurementInfoItem() %s", err.Error())
	}

	bytes, err := encodeXer(&C.asn_DEF_MeasurementInfoItem, unsafe.Pointer(measurementInfoItemCP))
	if err != nil {
		return nil, fmt.Errorf("xerEncodeMeasurementInfoItem() %s", err.Error())
	}
	return bytes, nil
}

func perEncodeMeasurementInfoItem(measurementInfoItem *e2sm_kpm_v2.MeasurementInfoItem) ([]byte, error) {
	measurementInfoItemCP, err := newMeasurementInfoItem(measurementInfoItem)
	if err != nil {
		return nil, fmt.Errorf("perEncodeMeasurementInfoItem() %s", err.Error())
	}

	bytes, err := encodePerBuffer(&C.asn_DEF_MeasurementInfoItem, unsafe.Pointer(measurementInfoItemCP))
	if err != nil {
		return nil, fmt.Errorf("perEncodeMeasurementInfoItem() %s", err.Error())
	}
	return bytes, nil
}

func xerDecodeMeasurementInfoItem(bytes []byte) (*e2sm_kpm_v2.MeasurementInfoItem, error) {
	unsafePtr, err := decodeXer(bytes, &C.asn_DEF_MeasurementInfoItem)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from XER is nil")
	}
	return decodeMeasurementInfoItem((*C.MeasurementInfoItem_t)(unsafePtr))
}

func perDecodeMeasurementInfoItem(bytes []byte) (*e2sm_kpm_v2.MeasurementInfoItem, error) {
	unsafePtr, err := decodePer(bytes, len(bytes), &C.asn_DEF_MeasurementInfoItem)
	if err != nil {
		return nil, err
	}
	if unsafePtr == nil {
		return nil, fmt.Errorf("pointer decoded from PER is nil")
	}
	return decodeMeasurementInfoItem((*C.MeasurementInfoItem_t)(unsafePtr))
}

func newMeasurementInfoItem(measurementInfoItem *e2sm_kpm_v2.MeasurementInfoItem) (*C.MeasurementInfoItem_t, error) {

	measTypeC, err := newMeasurementType(measurementInfoItem.MeasType)
	if err != nil {
		return nil, fmt.Errorf("newMeasurementType() %s", err.Error())
	}

	labelInfoListC, err := newLabelInfoList(measurementInfoItem.LabelInfoList)
	if err != nil {
		return nil, fmt.Errorf("newLabelInfoList() %s", err.Error())
	}

	measurementInfoItemC := C.MeasurementInfoItem_t{
		measType:      *measTypeC,
		labelInfoList: labelInfoListC,
	}

	return &measurementInfoItemC, nil
}

func decodeMeasurementInfoItem(measurementInfoItemC *C.MeasurementInfoItem_t) (*e2sm_kpm_v2.MeasurementInfoItem, error) {

	measType, err := decodeMeasurementType(&measurementInfoItemC.measType)
	if err != nil {
		return nil, fmt.Errorf("decodeMeasurementType() %s", err.Error())
	}

	labelInfoList, err := decodeLabelInfoList(measurementInfoItemC.labelInfoList)
	if err != nil {
		return nil, fmt.Errorf("decodeLabelInfoList() %s", err.Error())
	}

	measurementInfoItem := e2sm_kpm_v2.MeasurementInfoItem{
		//ToDo - check whether pointers passed correctly with regard to Protobuf's definition
		MeasType:      measType,
		LabelInfoList: labelInfoList,
	}

	return &measurementInfoItem, nil
}

func decodeMeasurementInfoItemBytes(array [8]byte) (*e2sm_kpm_v2.MeasurementInfoItem, error) { //ToDo - Check addressing correct structure in Protobuf
	measurementInfoItemC := (*C.MeasurementInfoItem_t)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(array[0:8]))))

	return decodeMeasurementInfoItem(measurementInfoItemC)
}