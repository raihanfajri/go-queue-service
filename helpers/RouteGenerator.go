package helpers

func GenerateUrls() []string {
	requestUrls := getAllFilesPath("../routes", `.yml`, "")

	return requestUrls
}
