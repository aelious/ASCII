package ASCII

func ConvertToAscii(imgUrl, textToWrite string) {
	GetImage(imgUrl)
	CreatePicText(textToWrite)
	ConvertToGrayScale()
	GetAndConvertColors()

}
