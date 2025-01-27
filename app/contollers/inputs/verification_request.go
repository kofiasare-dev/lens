package inputs

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/kofiasare/lens/app/services/db"
)

type FaceMatchInput struct {
	Reference      string                `form:"reference" validate:"required,len=32"`
	CallbackURL    string                `form:"callbackUrl" validate:"required,url"`
	ReferenceImage *multipart.FileHeader `form:"-"`
	TargetImage    *multipart.FileHeader `form:"-"`
}

func (fmi *FaceMatchInput) Bind(c fiber.Ctx) (err error) {
	if err = c.Bind().Form(fmi); err != nil {
		return
	}

	return
}

func (fmi *FaceMatchInput) Validate(c fiber.Ctx) (err error) {
	if err = c.Bind().Form(fmi); err != nil {
		return
	}

	if fmi.TargetImage, _ = c.FormFile("targetImage"); fmi.TargetImage == nil {
		return errors.New("targetImage is required")
	}

	if err = fmi.validateFile(fmi.TargetImage); err != nil {
		return errors.New("invalid targetImage: " + err.Error())
	}

	if fmi.ReferenceImage, _ = c.FormFile("referenceImage"); fmi.ReferenceImage == nil {
		return errors.New("referenceImage is required")
	}

	if err = fmi.validateFile(fmi.ReferenceImage); err != nil {
		return errors.New("invalid referenceImage: " + err.Error())
	}

	return
}

func (fmi *FaceMatchInput) UploadImages() (err error) {
	bc := db.GetBadgerClient()

	files := map[string]*multipart.FileHeader{
		"referenceImage": fmi.ReferenceImage,
		"targetImage":    fmi.TargetImage,
	}

	for id, file := range files {
		f, err := file.Open()
		if err != nil {
			return err
		}

		defer f.Close()

		data, err := io.ReadAll(f)
		if err != nil {
			return err
		}

		err = bc.Set(fmt.Sprintf("%s:%s", id, fmi.Reference), data)
		if err != nil {
			return err
		}

	}

	return nil
}

func (fmi *FaceMatchInput) validateFile(file *multipart.FileHeader) error {
	if file.Size > 5*1024*1024 {
		return errors.New("file size exceeds 5MB")
	}

	if !strings.EqualFold(file.Header.Get("Content-Type"), "image/jpeg") {
		return errors.New("file must be of type image/jpeg")
	}

	return nil
}
