package upload

import (
	"bytes"
	"github.com/tencentyun/cos-go-sdk-v5"
	"quick_web_golang/config"

	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type TencentCOS struct {
	client *cos.Client
}

func (t *TencentCOS) GetDomain() string {
	return fmt.Sprintf(
		"https://%s.cos.%s.myqcloud.com",
		config.Get(config.TencentCosBucket),
		config.Get(config.TencentCosRegion),
	)
}

func (t *TencentCOS) GetClient() *cos.Client {
	if t.client == nil {
		u, _ := url.Parse(t.GetDomain())
		baseURL := &cos.BaseURL{BucketURL: u}
		t.client = cos.NewClient(baseURL, &http.Client{
			Transport: &cos.AuthorizationTransport{
				SecretID:  config.Get(config.TencentCosSecretId),
				SecretKey: config.Get(config.TencentCosSecretKey),
			},
		})
	}
	return t.client
}

func (t *TencentCOS) UploadFile(file io.Reader, objectKey string) (string, error) {
	objectKey = strings.TrimPrefix(objectKey, "/")
	_, err := t.GetClient().Object.Put(context.Background(), objectKey, file, nil)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/%s", config.Get(config.TencentCosBasePath), objectKey), nil
}

func (t *TencentCOS) UploadBytes(file []byte, objectKey string) (string, error) {
	f := bytes.NewReader(file)
	return t.UploadFile(f, objectKey)
}

func (t *TencentCOS) DeleteFile(key string) error {
	name := config.Get(config.TencentCosBasePath) + "/" + key
	_, err := t.GetClient().Object.Delete(context.Background(), name)
	if err != nil {
		return err
	}
	return nil
}

func (t *TencentCOS) Copy(source, target string) error {
	source = strings.TrimPrefix(source, "/")
	target = strings.TrimPrefix(target, "/")

	sourceUrl := fmt.Sprintf(
		"%s.cos.%s.myqcloud.com/%s",
		config.Get(config.TencentCosBucket),
		config.Get(config.TencentCosRegion),
		source,
	)
	_, _, err := t.GetClient().Object.Copy(
		context.Background(),
		target,
		sourceUrl,
		nil,
	)

	return err
}
