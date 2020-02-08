package s3

import (
	"fmt"
	"github.com/lbernardo/aws-local/internal/adapters/secondary/storage"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type S3StorageLocal struct {
	params ParamsS3
}

func NewS3StorageLocal(params ParamsS3) *S3StorageLocal{
	return &S3StorageLocal{
		params:params,
	}
}

func (s *S3StorageLocal) getPathAndFile(uri string) (string,string) {
	ss := strings.Split(uri,"/")
	filename := ss[len(ss)-1]
	p := strings.Join(ss[:len(ss)-1],"/")
	return p,filename
}

func  (s *S3StorageLocal)  putFile(uri string, body io.ReadCloser) error {
	content,_ := ioutil.ReadAll(body)
	p, filename := s.getPathAndFile(uri)

	if err := storage.CreateVolume(fmt.Sprintf("%v/%v",s.params.Volume,p)) ; err != nil {
		return err
	}
	if err := storage.StorageContent(fmt.Sprintf("%v/%v/%v",s.params.Volume,p,filename), string(content)) ; err != nil {
		return err
	}
	return nil
}

func (s *S3StorageLocal) getFile(uri string) ([]byte, string, error) {
	p, filename := s.getPathAndFile(uri)
	return storage.ReadFile(fmt.Sprintf("%v/%v/%v",s.params.Volume,p,filename))
}

func  (s *S3StorageLocal)  StartS3Storage() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "PUT" {
			if err := s.putFile(r.RequestURI, r.Body) ;  err != nil {
				fmt.Println(err)
				w.WriteHeader(http.StatusBadGateway)
				return
			}
			w.WriteHeader(http.StatusCreated)
		}

		if r.Method == "GET" {
			content, contentType, err := s.getFile(r.RequestURI)
			if err != nil {
				fmt.Println(err)
				w.WriteHeader(http.StatusBadGateway)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Header().Add("Content-type",contentType)
			w.Write(content)
		}


	});

	fmt.Printf("Start S3 server %v:%v\n",s.params.Host, s.params.Port)
	http.ListenAndServe(fmt.Sprintf("%v:%v",s.params.Host, s.params.Port),nil)
}