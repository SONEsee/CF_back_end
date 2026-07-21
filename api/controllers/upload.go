package controllers

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/SONEsee/go-echo/api/presenters"
	"github.com/labstack/echo/v4"
)

var allowedImageExt = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
	".webp": true,
	".gif":  true,
}

var safeCategory = regexp.MustCompile(`^[a-z0-9_-]+$`)

// UploadImageController ຮັບ multipart/form-data field "file" ແລະ query param "type"
// (ຕົວຢ່າງ: shop, product) ໃຊ້ຈັດໂຟນເດີ, ບັນທຶກໄວ້ໃນ uploads/<type>/ ແລ້ວສົ່ງ URL ກັບຄືນ
func UploadImageController(c echo.Context) error {
	category := c.QueryParam("type")
	if category == "" {
		category = "misc"
	}
	if !safeCategory.MatchString(category) {
		return c.JSON(http.StatusBadRequest, presenters.ResponseError(
			"ປະເພດ (type) ບໍ່ຖືກຕ້ອງ",
			"ອະນຸຍາດສະເພາະຕົວອັກສອນນ້ອຍ, ຕົວເລກ, - ແລະ _",
		))
	}

	fileHeader, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, presenters.ResponseError(
			"ບໍ່ພົບໄຟລ໌",
			"ກະລຸນາແນບໄຟລ໌ຮູບພາບໃນ field ຊື່ 'file'",
		))
	}

	ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
	if !allowedImageExt[ext] {
		return c.JSON(http.StatusBadRequest, presenters.ResponseError(
			"ຮູບແບບໄຟລ໌ບໍ່ຖືກຕ້ອງ",
			"ອະນຸຍາດສະເພາະ .jpg, .jpeg, .png, .webp, .gif",
		))
	}

	const maxUploadSize = 5 << 20 // 5MB
	if fileHeader.Size > maxUploadSize {
		return c.JSON(http.StatusBadRequest, presenters.ResponseError(
			"ໄຟລ໌ໃຫຍ່ເກີນໄປ",
			"ຂະໜາດສູງສຸດ 5MB",
		))
	}

	src, err := fileHeader.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, presenters.ResponseError(
			"ເກີດຂໍ້ຜິດພາດ", "ບໍ່ສາມາດເປີດໄຟລ໌ໄດ້",
		))
	}
	defer src.Close()

	destDir := filepath.Join("uploads", category)
	if err := os.MkdirAll(destDir, 0755); err != nil {
		return c.JSON(http.StatusInternalServerError, presenters.ResponseError(
			"ເກີດຂໍ້ຜິດພາດ", "ບໍ່ສາມາດສ້າງໂຟນເດີໄດ້",
		))
	}

	randomSuffix := make([]byte, 8)
	if _, err := rand.Read(randomSuffix); err != nil {
		return c.JSON(http.StatusInternalServerError, presenters.ResponseError(
			"ເກີດຂໍ້ຜິດພາດ", "ບໍ່ສາມາດສ້າງຊື່ໄຟລ໌ໄດ້",
		))
	}
	filename := fmt.Sprintf("%d-%s%s", time.Now().UnixNano(), hex.EncodeToString(randomSuffix), ext)
	destPath := filepath.Join(destDir, filename)

	dst, err := os.Create(destPath)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, presenters.ResponseError(
			"ເກີດຂໍ້ຜິດພາດ", "ບໍ່ສາມາດບັນທຶກໄຟລ໌ໄດ້",
		))
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return c.JSON(http.StatusInternalServerError, presenters.ResponseError(
			"ເກີດຂໍ້ຜິດພາດ", "ບໍ່ສາມາດບັນທຶກໄຟລ໌ໄດ້",
		))
	}

	// ສົ່ງເປັນ absolute URL ເພາະ frontend ຢູ່ຄົນລະ origin/port ກັບ API
	url := fmt.Sprintf("%s://%s/uploads/%s/%s", c.Scheme(), c.Request().Host, category, filename)
	return c.JSON(http.StatusOK, presenters.ResponseSuccessWithData("ອັບໂຫຼດສຳເລັດ", map[string]string{
		"url": url,
	}))
}
