package base

import "os"

type OtomatConfig struct {
}

func (config OtomatConfig) BaseUrl1() string {
	return os.Getenv("OTOMAT_BASE_URL_1")
}

func (config OtomatConfig) ApiID1() string {
	return os.Getenv("OTOMAT_API_ID_1")
}

func (config OtomatConfig) ApiKey1() string {
	return os.Getenv("OTOMAT_API_KEY_1")
}

func (config OtomatConfig) BaseUrl2() string {
	return os.Getenv("OTOMAT_BASE_URL_2")
}

func (config OtomatConfig) ApiID2() string {
	return os.Getenv("OTOMAT_API_ID_2")
}

func (config OtomatConfig) ApiKey2() string {
	return os.Getenv("OTOMAT_API_KEY_2")
}
