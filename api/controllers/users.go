package controllers

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/SONEsee/go-echo/api/presenters"
	"github.com/SONEsee/go-echo/api/schema/requestbody"
	"github.com/SONEsee/go-echo/api/services"
	"github.com/SONEsee/go-echo/api/validators"

	"github.com/labstack/echo/v4"
)

const userProfileUploadDir = "./uploads/profile"

func GetUserController(c echo.Context) error {
	ctx := c.Request().Context()
	var id *int
	if idParam := c.QueryParam("id"); idParam != "" {
		parsedID, err := strconv.Atoi(idParam)
		if err != nil {
			return c.JSON(http.StatusBadRequest, presenters.ResponseError("ຮູບແບບ ID ບໍ່ຖືກຕ້ອງ", "ກະລຸນາປ້ອນ ID ເປັນຕົວເລກ"))
		}
		id = &parsedID
	}
	page, _ := strconv.Atoi(c.QueryParam("page"))
	pageSize, _ := strconv.Atoi(c.QueryParam("limit"))
	q := c.QueryParam("q")

	result, paginationResult, err := services.GetUserService(ctx, id, page, pageSize, q)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.JSON(http.StatusNotFound, presenters.ResponseError("ບໍ່ພົບຂໍ້ມູນ", err.Error()))
		}
		log.Printf("get user error: %v", err)
		return c.JSON(http.StatusInternalServerError, presenters.ResponseError("ເກີດຂໍ້ຜິດພາດ", "ບໍ່ສາມາດດຶງຂໍ້ມູນໄດ້"))
	}
	if paginationResult != nil {
		return c.JSON(http.StatusOK, presenters.ResponseSuccessListData(
			result,
			paginationResult.CurrentPage,
			paginationResult.CurrentPageTotalItem,
			paginationResult.TotalItems,
			paginationResult.TotalPage,
		))
	}
	return c.JSON(http.StatusOK, presenters.ResponseSuccessWithData("ດືງຂໍ້ມູນສຳເລັດ", result))
}

func CreateUserController(c echo.Context) error {
	var req requestbody.UserRequestBody

	// ---- ຮັບ multipart/form-data ----
	req.Username = c.FormValue("username")
	req.Password = c.FormValue("password")
	req.FullName = c.FormValue("full_name")
	req.Email = c.FormValue("email")
	req.Phone = c.FormValue("phone")

	if roleIDStr := c.FormValue("role_id"); roleIDStr != "" {
		roleID, err := strconv.Atoi(roleIDStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, presenters.ResponseError("ຂໍ້ມູນບໍ່ຖືກຕ້ອງ", "Role ID ຕ້ອງເປັນຕົວເລກ"))
		}
		req.RoleID = roleID
	}
	if shopIDStr := c.FormValue("shop_id"); shopIDStr != "" {
		shopID, err := strconv.Atoi(shopIDStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, presenters.ResponseError("ຂໍ້ມູນບໍ່ຖືກຕ້ອງ", "Shop ID ຕ້ອງເປັນຕົວເລກ"))
		}
		req.ShopID = &shopID
	}

	// ---- Validate ຟິວທີ່ບັງຄັບ ----
	if req.Username == "" || req.Password == "" || req.FullName == "" || req.RoleID == 0 {
		return c.JSON(http.StatusBadRequest, presenters.ResponseError(
			"ຂໍ້ມູນບໍ່ຖືກຕ້ອງ",
			"ກະລຸນາປ້ອນ username, password, full_name ແລະ role_id ໃຫ້ຄົບ",
		))
	}

	// ---- ຮັບໄຟລ໌ຮູບໂປຣໄຟລ໌ (ບໍ່ບັງຄັບ) ----
	fileHeader, err := c.FormFile("profile_image")
	if err == nil && fileHeader != nil {
		savedPath, saveErr := saveProfileImage(fileHeader)
		if saveErr != nil {
			return c.JSON(http.StatusBadRequest, presenters.ResponseError("ອັບໂຫລດຮູບບໍ່ສຳເລັດ", saveErr.Error()))
		}
		req.ProfileImage = savedPath
	}

	if err := services.CreateUserService(c.Request().Context(), req); err != nil {
		log.Printf("create user error: %v", err)
		return c.JSON(http.StatusBadRequest, presenters.ResponseError("ບໍ່ສາມາດສ້າງຜູ້ໃຊ້ໄດ້", err.Error()))
	}
	return c.JSON(http.StatusOK, presenters.ResponseSuccess("SUCCESS"))
}

// generateRandomHex ສ້າງ random string ດ້ວຍ crypto/rand (standard library, ບໍ່ຕ້ອງເພີ່ມ dependency ໃໝ່)
func generateRandomHex(n int) (string, error) {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

// saveProfileImage validate + save ໄຟລ໌ຮູບ, return path ທີ່ຈະເກັບໃນ DB (ເຊັ່ນ /uploads/profile/xxx.jpg)
func saveProfileImage(fileHeader *multipart.FileHeader) (string, error) {
	const maxFileSize = 2 << 20 // 2MB
	allowedExt := map[string]bool{".jpg": true, ".jpeg": true, ".png": true, ".webp": true}

	if fileHeader.Size > maxFileSize {
		return "", fmt.Errorf("ຂະໜາດໄຟລ໌ຕ້ອງບໍ່ເກີນ 2MB")
	}

	ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
	if !allowedExt[ext] {
		return "", fmt.Errorf("ອະນຸຍາດສະເພາະໄຟລ໌ຮູບພາບ (jpg, jpeg, png, webp)")
	}

	if err := os.MkdirAll(userProfileUploadDir, 0o755); err != nil {
		return "", fmt.Errorf("ບໍ່ສາມາດສ້າງໂຟນເດີອັບໂຫລດໄດ້: %w", err)
	}

	src, err := fileHeader.Open()
	if err != nil {
		return "", fmt.Errorf("ບໍ່ສາມາດເປີດໄຟລ໌ຮູບໄດ້: %w", err)
	}
	defer src.Close()

	randomPart, err := generateRandomHex(8)
	if err != nil {
		return "", fmt.Errorf("ບໍ່ສາມາດສ້າງຊື່ໄຟລ໌ໄດ້: %w", err)
	}
	newFileName := fmt.Sprintf("%d_%s%s", time.Now().Unix(), randomPart, ext)
	dstPath := filepath.Join(userProfileUploadDir, newFileName)

	dst, err := os.Create(dstPath)
	if err != nil {
		return "", fmt.Errorf("ບໍ່ສາມາດບັນທຶກໄຟລ໌ຮູບໄດ້: %w", err)
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		os.Remove(dstPath)
		return "", fmt.Errorf("ບໍ່ສາມາດຄັດລອກໄຟລ໌ຮູບໄດ້: %w", err)
	}

	return "/uploads/profile/" + newFileName, nil
}

// DeactivateUserController ໃຊ້ແທນການລົບ (users ບໍ່ມີ deleted_at) — set is_active = false
func DeactivateUserController(c echo.Context) error {
	idParam, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, presenters.ResponseError("ຮູບແບບ ID ບໍ່ຖືກຕ້ອງ", "ກະລຸນາປ້ອນ ID ເປັນຕົວເລກ"))
	}
	if err := services.DeactivateUserServices(c.Request().Context(), idParam); err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "not found") {
			return c.JSON(http.StatusNotFound, presenters.ResponseError("ບໍ່ພົບຂໍ້ມູນ", err.Error()))
		}
		log.Printf("deactivate user error: %v", err)
		return c.JSON(http.StatusInternalServerError, presenters.ResponseError("ເກີດຂໍ້ຜິດພາດ", "ບໍ່ສາມາດປິດການໃຊ້ງານໄດ້"))
	}
	return c.JSON(http.StatusOK, presenters.ResponseSuccess("ປິດການໃຊ້ງານສຳເລັດ"))
}

func UpdateUserController(c echo.Context) error {
	idParam, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, presenters.ResponseError(
			"ຮູບແບບ ID ບໍ່ຖືກຕ້ອງ",
			"ກະລຸນາປ້ອນ ID ເປັນຕົວເລກ",
		))
	}

	var req requestbody.UserRequestBodyPacth
	if err := validators.ParseAndValidateBody(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, presenters.ResponseError(
			"ຂໍ້ມູນບໍ່ຖືກຕ້ອງ",
			err.Error(),
		))
	}

	ctx := c.Request().Context()
	err = services.UpdateUserServices(ctx, idParam, req)
	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "not found") {
			return c.JSON(http.StatusNotFound, presenters.ResponseError(
				"ບໍ່ພົບຂໍ້ມູນ",
				fmt.Sprintf("ບໍ່ພົບຜູ້ໃຊ້ ID: %d", idParam),
			))
		}
		if strings.Contains(strings.ToLower(err.Error()), "already exists") {
			return c.JSON(http.StatusConflict, presenters.ResponseError(
				"Username ຊ້ຳກັນ",
				err.Error(),
			))
		}
		log.Printf("Update user error (ID: %d): %v", idParam, err)
		return c.JSON(http.StatusInternalServerError, presenters.ResponseError(
			"ເກີດຂໍ້ຜິດພາດ",
			"ບໍ່ສາມາດອັບເດດຂໍ້ມູນໄດ້",
		))
	}

	return c.JSON(http.StatusOK, presenters.ResponseSuccess(
		"ອັບເດດຂໍ້ມູນສຳເລັດ",
	))
}

func UserAuthController(c echo.Context) error {
	claims := map[string]interface{}{
		"user_id":   c.Get("user_id"),
		"user_name": c.Get("user_name"),
		"role_id":   c.Get("role_id"),
	}
	return c.JSON(http.StatusOK, presenters.ResponseSuccess(claims))
}