package magickwand

/*
#cgo LDFLAGS: -lMagickWand -lMagickCore
#cgo CFLAGS: -fopenmp -I/usr/include/ImageMagick

#include <wand/magick_wand.h>
*/
import "C"

import (
	"fmt"
	"unsafe"
)

const (
	/* FilterTypes */
	BesselFilter   = int(C.BesselFilter)
	BlackmanFilter = int(C.BlackmanFilter)
	BoxFilter      = int(C.BoxFilter)
	CatromFilter   = int(C.CatromFilter)
	GaussianFilter = int(C.GaussianFilter)
	HanningFilter  = int(C.HanningFilter)
	HermiteFilter  = int(C.HermiteFilter)
	LanczosFilter  = int(C.LanczosFilter)
	MitchellFilter = int(C.MitchellFilter)
	SincFilter     = int(C.SincFilter)
	TriangleFilter = int(C.TriangleFilter)

	/* Compression */
	UndefinedCompression    = int(C.UndefinedCompression)
	NoCompression           = int(C.NoCompression)
	BZipCompression         = int(C.BZipCompression)
	DXT1Compression         = int(C.DXT1Compression)
	DXT3Compression         = int(C.DXT3Compression)
	DXT5Compression         = int(C.DXT5Compression)
	FaxCompression          = int(C.FaxCompression)
	Group4Compression       = int(C.Group4Compression)
	JPEGCompression         = int(C.JPEGCompression)
	JPEG2000Compression     = int(C.JPEG2000Compression)
	LosslessJPEGCompression = int(C.LosslessJPEGCompression)
	LZWCompression          = int(C.LZWCompression)
	RLECompression          = int(C.RLECompression)
	ZipCompression          = int(C.ZipCompression)
	ZipSCompression         = int(C.ZipSCompression)
	PizCompression          = int(C.PizCompression)
	Pxr24Compression        = int(C.Pxr24Compression)
	B44Compression          = int(C.B44Compression)
	B44ACompression         = int(C.B44ACompression)
)

type MagickWand struct {
	wand *C.MagickWand
}

/* Initializes the MagickWand environment. */
func Genesis() {
	C.MagickWandGenesis()
}

/* Terminates the MagickWand environment. */
func Terminus() {
	C.MagickWandTerminus()
}

/* 
   Create a wand required for all other methods in the API. A panic() is 
   thrown if there is not enough memory to allocate the wand. Use Destroy() to 
   dispose of the wand when it is no longer needed. 
*/
func New() *MagickWand {
	return &MagickWand{wand: C.NewMagickWand()}
}

/* 
   clears resources associated with the wand, leaving the wand blank, and 
   ready to be used for a new set of images. 
*/
func (w *MagickWand) Clear() {
	C.ClearMagickWand(w.wand)
}

/* Deallocates memory associated with an MagickWand. */
func (w *MagickWand) Destroy() {
	C.DestroyMagickWand(w.wand)
}

func (w *MagickWand) Exception() string {
	var severity C.ExceptionType
	errPtr := C.MagickGetException(w.wand, &severity)
	C.MagickClearException(w.wand)
	err := C.GoString(errPtr)
	C.MagickRelinquishMemory(unsafe.Pointer(errPtr))

	return err
}

/* Reads an image or image sequence from a blob. */
func (w *MagickWand) ReadBlob(blob []byte, length uint) error {
	if C.MagickReadImageBlob(w.wand, unsafe.Pointer(&blob[0]),
		C.size_t(length)) == C.MagickFalse {
		return fmt.Errorf("ReadBlob() failed : %s", w.Exception())
	}

	return nil
}

/* Implements direct to memory image formats. */
func (w *MagickWand) Blob(length *uint) []byte {
	blobPtr := unsafe.Pointer(C.MagickGetImageBlob(w.wand,
		(*C.size_t)(unsafe.Pointer(length))))
	blob := C.GoBytes(blobPtr, C.int(int(*length)))
	C.MagickRelinquishMemory(unsafe.Pointer(blobPtr))
	return blob
}

func (w *MagickWand) Read(fileName string) error {
	if C.MagickReadImage(w.wand, C.CString(fileName)) == C.MagickFalse {
		return fmt.Errorf("Read() failed : %s", w.Exception())
	}

	return nil
}

/* Writes an image to the specified filename. */
func (w *MagickWand) Write(fileName string) error {
	if C.MagickWriteImage(w.wand, C.CString(fileName)) == C.MagickFalse {
		return fmt.Errorf("Write() failed : %s", w.Exception())
	}

	return nil
}

/* Extracts a region of the image. */
func (w *MagickWand) Crop(width, height uint, x, y int) error {
	if C.MagickCropImage(w.wand, C.size_t(width), C.size_t(height),
		C.ssize_t(x), C.ssize_t(y)) == C.MagickFalse {
		return fmt.Errorf("Crop() failed : %s", w.Exception())
	}

	return nil
}

/* Adaptively resize image with data dependent triangulation. */
func (w *MagickWand) AdaptiveResize(columns, rows uint) error {
	if C.MagickAdaptiveResizeImage(w.wand, C.size_t(columns),
		C.size_t(rows)) == C.MagickFalse {
		return fmt.Errorf("AdaptiveResize() failed : %s", w.Exception())
	}

	return nil
}

/* Scales an image to the desired dimensions. */
func (w *MagickWand) Resize(columns, rows uint, filter int,
	blur float64) error {
	if C.MagickResizeImage(w.wand, C.size_t(columns),
		C.size_t(rows), C.FilterTypes(filter), C.double(blur)) ==
		C.MagickFalse {
		return fmt.Errorf("Resize() failed : %s", w.Exception())
	}

	return nil
}

/* Gets the image height. */
func (w *MagickWand) Height() uint {
	return uint(C.MagickGetImageHeight(w.wand))
}

/* Gets the image width. */
func (w *MagickWand) Width() uint {
	return uint(C.MagickGetImageWidth(w.wand))
}

/* Sets the image compression quality. */
func (w *MagickWand) SetQuality(quality uint) error {
	if C.MagickSetImageCompressionQuality(w.wand, C.size_t(quality)) ==
		C.MagickFalse {
		return fmt.Errorf("SetQuality() failed : %s", w.Exception())
	}

	return nil
}

/* Gets the image compression quality. */
func (w *MagickWand) Quality() uint {
	return uint(C.MagickGetImageCompressionQuality(w.wand))
}

/* Sets the image compression. */
func (w *MagickWand) SetCompression(compression int) error {

	if C.MagickSetImageCompression(w.wand, C.CompressionType(compression)) ==
		C.MagickFalse {
		return fmt.Errorf("SetCompression() failed : %s", w.Exception())
	}

	return nil
}

/* Gets the image compression. */
func (w *MagickWand) Compression() int {
	return int(C.MagickGetImageCompression(w.wand))
}

/* Sets the format of a particular image in a sequence */
func (w *MagickWand) SetFormat(format string) error {
	if C.MagickSetImageFormat(w.wand, C.CString(format)) == C.MagickFalse {
		return fmt.Errorf("SetFormat() failed : %s", w.Exception())
	}

	return nil
}

/* Gets the format of a particular image in a sequence. */
func (w *MagickWand) Format() string {
	return C.GoString(C.MagickGetImageFormat(w.wand))
}
