package types

type MediaType string

// All media types extensions which are supported
// by facebook whatsapp business, please follow the
// link https://developers.facebook.com/docs/whatsapp/api/media/
// for reference on supported media types
const (
	MediaTypeImageJpeg    MediaType = "image/jpeg"
	MediaTypeImagePng     MediaType = "image/png"
	MediaTypeImageWebp    MediaType = "image/webp"
	MediaTypeAudioAcc     MediaType = "audio/aac"
	MediaTypeAudioAmr     MediaType = "audio/amr"
	MediaTypeAudioMpeg    MediaType = "audio/mpeg"
	MediaTypeAudioOgg     MediaType = "audio/ogg"
	MediaTypeAudioOpus    MediaType = "audio/opus"
	MediaTypeVideoMp4     MediaType = "video/mp4"
	MediaTypeVideo3gpp    MediaType = "video/3gpp"
	MediaTypeDocumentDoc  MediaType = "application/msword"
	MediaTypeDocumentDocx MediaType = "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
	MediaTypeDocumentPdf  MediaType = "application/pdf"
	MediaTypeDocumentPpt  MediaType = "application/vnd.ms-powerpoint"
	MediaTypeDocumentPptx MediaType = "application/vnd.openxmlformats-officedocument.presentationml.presentation"
	MediaTypeDocumentXls  MediaType = "application/vnd.ms-excel"
	MediaTypeDocumentXlsx MediaType = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
	MediaTypeDocumentZip  MediaType = "application/zip"
	MediaTypeDocumentRar  MediaType = "application/vnd.rar"
)

var (
	// MediaTypeMapping provides available extension to mime
	// type mapping which can be uploaded to whatsapp
	MediaTypeMapping = map[string]MediaType{
		".jpeg": MediaTypeImageJpeg,
		".png":  MediaTypeImagePng,
		".webp": MediaTypeImageWebp,
		".acc" : MediaTypeAudioAcc,
		".mp3":  MediaTypeAudioMpeg,
		".amr":  MediaTypeAudioAmr,
		".oga":  MediaTypeAudioOgg,
		".opus": MediaTypeAudioOpus,
		".mp4":  MediaTypeVideoMp4,
		".3gp":  MediaTypeVideo3gpp,
		".doc":  MediaTypeDocumentDoc,
		".docx": MediaTypeDocumentDocx,
		".pdf":  MediaTypeDocumentPdf,
		".ppt":  MediaTypeDocumentPpt,
		".pptx": MediaTypeDocumentPptx,
		".xls":  MediaTypeDocumentXls,
		".xlsx": MediaTypeDocumentXlsx,
		".zip":  MediaTypeDocumentZip,
		".rar":  MediaTypeDocumentRar,
	}
)