package hpdf

/*
#cgo LDFLAGS: -lhpdf -lpng -lz
#include "hpdf.h"
#include "hpdf_types.h"
*/
import "C"
import (
	"fmt"
)

var messageMap = map[C.HPDF_STATUS]string{
	C.HPDF_ARRAY_COUNT_ERR:              "Internal error. Data consistency was lost",
	C.HPDF_ARRAY_ITEM_NOT_FOUND:         "Internal error. Data consistency was lost",
	C.HPDF_ARRAY_ITEM_UNEXPECTED_TYPE:   "Internal error. Data consistency was lost",
	C.HPDF_BINARY_LENGTH_ERR:            "Data length > HPDF_LIMIT_MAX_STRING_LEN",
	C.HPDF_CANNOT_GET_PALLET:            "Cannot get pallet data from PNG image",
	C.HPDF_DICT_COUNT_ERR:               "Dictionary elements > HPDF_LIMIT_MAX_DICT_ELEMENT",
	C.HPDF_DICT_ITEM_NOT_FOUND:          "Internal error. Data consistency was lost",
	C.HPDF_DICT_ITEM_UNEXPECTED_TYPE:    "Internal error. Data consistency was lost",
	C.HPDF_DICT_STREAM_LENGTH_NOT_FOUND: "Internal error. Data consistency was lost",
	C.HPDF_DOC_ENCRYPTDICT_NOT_FOUND:    "HPDF_SetEncryptMode() or HPDF_SetPermission() called before password set",
	C.HPDF_DOC_INVALID_OBJECT:           "Internal error. Data consistency was lost",
	C.HPDF_DUPLICATE_REGISTRATION:       "Tried to re-register a registered font",
	C.HPDF_EXCEED_JWW_CODE_NUM_LIMIT:    "Cannot register a character to the Japanese word wrap characters list",
	C.HPDF_ENCRYPT_INVALID_PASSWORD:     "Tried to set the owner password to NULL or Owner and user password are the same",
	C.HPDF_ERR_UNKNOWN_CLASS:            "Internal error. Data consistency was lost",
	C.HPDF_EXCEED_GSTATE_LIMIT:          "Stack depth > HPDF_LIMIT_MAX_GSTATE",
	C.HPDF_FAILD_TO_ALLOC_MEM:           "Memory allocation failed",
	C.HPDF_FILE_IO_ERROR:                "File processing failed. (Detailed code is set.)",
	C.HPDF_FILE_OPEN_ERROR:              "Cannot open a file. (Detailed code is set.)",
	C.HPDF_FONT_EXISTS:                  "Tried to load a font that has been registered",
	C.HPDF_FONT_INVALID_WIDTHS_TABLE:    "Font-file format is invalid or Internal error. Data consistency was lost",
	C.HPDF_INVALID_AFM_HEADER:           "Cannot recognize header of afm file",
	C.HPDF_INVALID_ANNOTATION:           "Specified annotation handle is invalid",
	C.HPDF_INVALID_BIT_PER_COMPONENT:    "Bit-per-component of a image which was set as mask-image is invalid",
	C.HPDF_INVALID_CHAR_MATRICS_DATA:    "Cannot recognize char-matrics-data of afm file",
	C.HPDF_INVALID_COLOR_SPACE:          "Invalid color_space parameter of HPDF_LoadRawImage, Color-space of a image which was set as mask-image is invalid or Invoked function invalid in present color-space",
	C.HPDF_INVALID_COMPRESSION_MODE:     "Invalid value set when invoking HPDF_SetCommpressionMode()",
	C.HPDF_INVALID_DATE_TIME:            "An invalid date-time value was set",
	C.HPDF_INVALID_DESTINATION:          "An invalid destination handle was set",
	C.HPDF_INVALID_DOCUMENT:             "An invalid document handle was set",
	C.HPDF_INVALID_DOCUMENT_STATE:       "Function invalid in the present state was invoked",
	C.HPDF_INVALID_ENCODER:              "An invalid encoder handle was set",
	C.HPDF_INVALID_ENCODER_TYPE:         "Combination between font and encoder is wrong",
	C.HPDF_INVALID_ENCODING_NAME:        "An Invalid encoding name is specified",
	C.HPDF_INVALID_ENCRYPT_KEY_LEN:      "Encryption key length is invalid",
	C.HPDF_INVALID_FONTDEF_DATA:         "An invalid font handle was set or Unsupported font format",
	C.HPDF_INVALID_FONTDEF_TYPE:         "Internal error. Data consistency was lost",
	C.HPDF_INVALID_FONT_NAME:            "Font with the specified name is not found",
	C.HPDF_INVALID_IMAGE:                "Unsupported image format",
	C.HPDF_INVALID_JPEG_DATA:            "Unsupported image format",
	C.HPDF_INVALID_N_DATA:               "Cannot read a postscript-name from an afm file",
	C.HPDF_INVALID_OBJECT:               "An invalid object is set or Internal error. Data consistency was lost",
	C.HPDF_INVALID_OBJ_ID:               "Internal error. Data consistency was lost",
	C.HPDF_INVALID_OPERATION:            "Invoked HPDF_Image_SetColorMask() against the image-object which was set a mask-image",
	C.HPDF_INVALID_OUTLINE:              "An invalid outline-handle was specified",
	C.HPDF_INVALID_PAGE:                 "An invalid page-handle was specified",
	C.HPDF_INVALID_PAGES:                "An invalid pages-handle was specified. (internal error)",
	C.HPDF_INVALID_PARAMETER:            "An invalid value is set",
	C.HPDF_INVALID_PNG_IMAGE:            "Invalid PNG image format",
	C.HPDF_INVALID_STREAM:               "Internal error. Data consistency was lost",
	C.HPDF_MISSING_FILE_NAME_ENTRY:      "Internal error. \"_FILE_NAME\" entry for delayed loading is missing",
	C.HPDF_INVALID_TTC_FILE:             "Invalid .TTC file format",
	C.HPDF_INVALID_TTC_INDEX:            "Index parameter > number of included fonts",
	C.HPDF_INVALID_WX_DATA:              "Cannot read a width-data from an afm file",
	C.HPDF_ITEM_NOT_FOUND:               "Internal error. Data consistency was lost",
	C.HPDF_LIBPNG_ERROR:                 "Error returned from PNGLIB while loading image",
	C.HPDF_NAME_INVALID_VALUE:           "Internal error. Data consistency was lost",
	C.HPDF_NAME_OUT_OF_RANGE:            "Internal error. Data consistency was lost",
	C.HPDF_PAGES_MISSING_KIDS_ENTRY:     "Internal error. Data consistency was lost",
	C.HPDF_PAGE_CANNOT_FIND_OBJECT:      "Internal error. Data consistency was lost",
	C.HPDF_PAGE_CANNOT_GET_ROOT_PAGES:   "Internal error. Data consistency was lost",
	C.HPDF_PAGE_CANNOT_RESTORE_GSTATE:   "There are no graphics-states to be restored",
	C.HPDF_PAGE_CANNOT_SET_PARENT:       "Internal error. Data consistency was lost",
	C.HPDF_PAGE_FONT_NOT_FOUND:          "The current font is not set",
	C.HPDF_PAGE_INVALID_FONT:            "An invalid font-handle was specified",
	C.HPDF_PAGE_INVALID_FONT_SIZE:       "An invalid font-size was set",
	C.HPDF_PAGE_INVALID_GMODE:           "See Graphics mode",
	C.HPDF_PAGE_INVALID_INDEX:           "Internal error. Data consistency was lost",
	C.HPDF_PAGE_INVALID_ROTATE_VALUE:    "Specified value is not multiple of 90",
	C.HPDF_PAGE_INVALID_SIZE:            "An invalid page-size was set",
	C.HPDF_PAGE_INVALID_XOBJECT:         "An invalid image-handle was set",
	C.HPDF_PAGE_OUT_OF_RANGE:            "The specified value is out of range",
	C.HPDF_REAL_OUT_OF_RANGE:            "The specified value is out of range",
	C.HPDF_STREAM_EOF:                   "Unexpected EOF marker was detected",
	C.HPDF_STREAM_READLN_CONTINUE:       "Internal error. Data consistency was lost",
	C.HPDF_STRING_OUT_OF_RANGE:          "The length of the text is too long",
	C.HPDF_THIS_FUNC_WAS_SKIPPED:        "Function not executed because of other errors",
	C.HPDF_TTF_CANNOT_EMBEDDING_FONT:    "Font cannot be embedded. (license restriction)",
	C.HPDF_TTF_INVALID_CMAP:             "Unsupported ttf format. (cannot find unicode cmap)",
	C.HPDF_TTF_INVALID_FOMAT:            "Unsupported ttf format",
	C.HPDF_TTF_MISSING_TABLE:            "Unsupported ttf format. (cannot find a necessary table)",
	C.HPDF_UNSUPPORTED_FONT_TYPE:        "Internal error. Data consistency was lost",
	C.HPDF_UNSUPPORTED_FUNC:             "Library not configured to use PNGLIB or Internal error. Data consistency was lost",
	C.HPDF_UNSUPPORTED_JPEG_FORMAT:      "Unsupported JPEG format",
	C.HPDF_UNSUPPORTED_TYPE1_FONT:       "Failed to parse .PFB file",
	C.HPDF_XREF_COUNT_ERR:               "Internal error. Data consistency was lost",
	C.HPDF_ZLIB_ERROR:                   "Error while executing ZLIB function",
	C.HPDF_INVALID_PAGE_INDEX:           "An invalid page index was passed",
	C.HPDF_INVALID_URI:                  "An invalid URI was set",
	C.HPDF_ANNOT_INVALID_ICON:           "An invalid icon was set",
	C.HPDF_ANNOT_INVALID_BORDER_STYLE:   "An invalid border-style was set",
	C.HPDF_PAGE_INVALID_DIRECTION:       "An invalid page-direction was set",
	C.HPDF_INVALID_FONT:                 "An invalid font-handle was specified",
}

type Error struct {
	status C.HPDF_STATUS
	detail C.HPDF_STATUS
}

func NewError(status, detail C.HPDF_STATUS) *Error {
	return &Error{status, detail}
}

func (err *Error) Error() string {
	return fmt.Sprintf(
		"HPDF Error: 0x%04X-0x%04X: %s",
		err.status,
		err.detail,
		err.message(),
	)
}

func (err *Error) message() string {
	msg, hit := messageMap[err.status]

	if !hit {
		return "Unknown"
	}

	return msg
}
