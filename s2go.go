package main

/*
#include <stdlib.h>
*/
import "C"

import (
	"strconv"
	"fmt"
	"github.com/golang/geo/s2"
	"unsafe"
	"encoding/json"
	"math"
)


const EARTH_RADIUS = 6371.01

func Radius2height (radius float64) float64 {
	return 1 - math.Sqrt(1 - math.Pow((radius / EARTH_RADIUS), 2))
}

//export S2CellID
func S2CellID(lat, lng C.double) C.longlong {
	latFloat := float64(lat)
	lngFloat := float64(lng)
	ll := s2.LatLngFromDegrees(latFloat, lngFloat).Normalized()
	cellId := s2.CellFromLatLng(ll).ID()
	ret, _ := strconv.ParseInt(fmt.Sprintf("%d", cellId), 10, 64)
	return C.longlong(ret)
}

//export S2GetCoving
func S2GetCoving(lat, lng, radius C.double, maxcell C.int) unsafe.Pointer {
	cellids := make([][2]int64, 0)
	latFloat := float64(lat)
	lngFloat := float64(lng)
	radiusFloat := float64(radius)
	maxcellInt := int(maxcell)
	ll := s2.LatLngFromDegrees(latFloat, lngFloat).Normalized()
	p := s2.PointFromLatLng(ll)
	cap := s2.CapFromCenterHeight(p, Radius2height(radiusFloat/1000.0))

	rc := s2.RegionCoverer{
		MinLevel: 0,
		MaxLevel: 30,
		LevelMod: 1,
		MaxCells: maxcellInt,
	}

	cells := rc.Covering(cap)
	for _, cell := range cells {
		b := fmt.Sprintf("%d", cell)
		e := fmt.Sprintf("%d", cell.Next())
		bInt64, _ := strconv.ParseInt(b, 10, 64)
		eInt64, _ := strconv.ParseInt(e, 10, 64)
		cellids = append(cellids, [2]int64{bInt64, eInt64})
	}

	//返回CString指针
	jsonByte, _ := json.Marshal(cellids)
	jsonCString := C.CString(string(jsonByte))

	defer C.free(unsafe.Pointer(jsonCString))
	return unsafe.Pointer(jsonCString)
}

func main() {
	
}