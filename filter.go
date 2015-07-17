package gmf

/*

#cgo pkg-config: libavfilter

#include <stdlib.h>
#include <unistd.h>
#include "libavcodec/avcodec.h"
#include "libavformat/avformat.h"
#include "libavfilter/avfiltergraph.h"
#include "libavfilter/avcodec.h"
#include "libavfilter/buffersink.h"
#include "libavfilter/buffersrc.h"
#include "libavutil/opt.h"
*/
import "C"
import (
	"errors"
	"fmt"
)

// var (
// AVMEDIA_TYPE_AUDIO int32 = C.AVMEDIA_TYPE_AUDIO
// AVMEDIA_TYPE_VIDEO int32 = C.AVMEDIA_TYPE_VIDEO
//
// AV_PIX_FMT_YUV420P      int32 = C.AV_PIX_FMT_YUV420P
// AV_PIX_FMT_YUVJ420P     int32 = C.AV_PIX_FMT_YUVJ420P
// AV_PIX_FMT_RGB24        int32 = C.AV_PIX_FMT_RGB24
// AV_PIX_FMT_NONE         int32 = C.AV_PIX_FMT_NONE
// FF_PROFILE_MPEG4_SIMPLE int   = C.FF_PROFILE_MPEG4_SIMPLE
// AV_NOPTS_VALUE          int   = C.AV_NOPTS_VALUE
// )

type AVFilter struct {
	avFilter *C.struct_AVFilter
}

type AVFilterInOut struct {
	avFilterInOut *C.struct_AVFilterInOut
}

func init() {
	println("filter init")
	C.avfilter_register_all()
}

// const AVFilter* avfilter_get_by_name	(	const char * 	name	)
func GetFilter(name string) (*AVFilter, error) {
	f := C.avfilter_get_by_name(C.CString(name))
	if f == nil {
		return nil, errors.New(fmt.Sprintf("Unable to GetFilter with name %s", name))
	}
	return &AVFilter{avFilter: C.avfilter_get_by_name(C.CString(name))}, nil
}

// AVFilterInOut* avfilter_inout_alloc	(	void 		)
func NewFilterInOut() (*AVFilterInOut, error) {
	f := C.avfilter_inout_alloc()
	if f == nil {
		return nil, errors.New("Unable to create a new FilterInOut")
	}
	return &AVFilterInOut{avFilterInOut: f}, nil
}

// void avfilter_inout_free	(	AVFilterInOut ** 	inout	)
func (this *AVFilterInOut) Free() {
	C.avfilter_inout_free(&this.avFilterInOut)
}
