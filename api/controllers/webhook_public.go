package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/SONEsee/go-echo/api/services"
	webhookverify "github.com/SONEsee/go-echo/pkg/webhook-verify"
	"github.com/labstack/echo/v4"
)

// facebookEnvelope ດຶງພຽງ field ທີ່ຈຳເປັນ (page id) ອອກຈາກ payload ຈິງຂອງ Facebook — ບໍ່ parse ຄົບໂຄງສ້າງ
type facebookEnvelope struct {
	Object string `json:"object"`
	Entry  []struct {
		ID string `json:"id"`
	} `json:"entry"`
}

// lineEnvelope ດຶງພຽງ destination (channel id) ອອກຈາກ payload ຈິງຂອງ Line
type lineEnvelope struct {
	Destination string `json:"destination"`
}

// FacebookWebhookVerifyController ຮັບ handshake ຕອນຕັ້ງຄ່າ webhook ຄັ້ງທຳອິດໃນ Facebook Developer Console
// (GET /webhook/facebook?hub.mode=subscribe&hub.verify_token=...&hub.challenge=...)
func FacebookWebhookVerifyController(c echo.Context) error {
	mode := c.QueryParam("hub.mode")
	token := c.QueryParam("hub.verify_token")
	challenge := c.QueryParam("hub.challenge")

	expectedToken := os.Getenv("FB_WEBHOOK_VERIFY_TOKEN")
	if mode == "subscribe" && expectedToken != "" && token == expectedToken {
		return c.String(http.StatusOK, challenge)
	}
	return c.String(http.StatusForbidden, "verification failed")
}

// FacebookWebhookReceiveController ຮັບ event ຈິງ — ກວດ X-Hub-Signature-256 ກ່ອນສະເໝີ
func FacebookWebhookReceiveController(c echo.Context) error {
	rawBody, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	appSecret := os.Getenv("FB_APP_SECRET")
	signature := c.Request().Header.Get("X-Hub-Signature-256")
	if !webhookverify.VerifyFacebookSignature(appSecret, rawBody, signature) {
		log.Println("❌ Facebook webhook: invalid signature")
		return c.NoContent(http.StatusUnauthorized)
	}

	var envelope facebookEnvelope
	pageID := ""
	if err := json.Unmarshal(rawBody, &envelope); err == nil && len(envelope.Entry) > 0 {
		pageID = envelope.Entry[0].ID
	}

	if pageID == "" {
		log.Println("⚠️ Facebook webhook: cannot resolve page id from payload")
		return c.NoContent(http.StatusOK) // ຕອບ 200 ສະເໝີເພື່ອບໍ່ໃຫ້ Facebook retry ຊ້ຳໆ
	}

	if err := services.IngestWebhookEventServices(c.Request().Context(), "FACEBOOK_PAGE", pageID, envelope.Object, string(rawBody)); err != nil {
		log.Printf("⚠️ Facebook webhook ingest error (page=%s): %v", pageID, err)
	}

	return c.NoContent(http.StatusOK)
}

// LineWebhookReceiveController ຮັບ event ຈິງ — ກວດ X-Line-Signature ກ່ອນສະເໝີ
func LineWebhookReceiveController(c echo.Context) error {
	rawBody, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	channelSecret := os.Getenv("LINE_CHANNEL_SECRET")
	signature := c.Request().Header.Get("X-Line-Signature")
	if !webhookverify.VerifyLineSignature(channelSecret, rawBody, signature) {
		log.Println("❌ Line webhook: invalid signature")
		return c.NoContent(http.StatusUnauthorized)
	}

	var envelope lineEnvelope
	if err := json.Unmarshal(rawBody, &envelope); err != nil || envelope.Destination == "" {
		log.Println("⚠️ Line webhook: cannot resolve destination from payload")
		return c.NoContent(http.StatusOK)
	}

	if err := services.IngestWebhookEventServices(c.Request().Context(), "LINE_OA", envelope.Destination, "line_event", string(rawBody)); err != nil {
		log.Printf("⚠️ Line webhook ingest error (destination=%s): %v", envelope.Destination, err)
	}

	return c.NoContent(http.StatusOK)
}
