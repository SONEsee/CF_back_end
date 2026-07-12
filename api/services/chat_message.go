package services

import (
	"context"
	"fmt"

	"github.com/SONEsee/go-echo/api/schema/requestbody"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
	dbinserts "github.com/SONEsee/go-echo/pkg/db-pkg/db-inserts"
	dbquery "github.com/SONEsee/go-echo/pkg/db-pkg/db-query"
	dbschema "github.com/SONEsee/go-echo/pkg/db-pkg/db-schema"
	dbupdate "github.com/SONEsee/go-echo/pkg/db-pkg/db-update"
	"github.com/SONEsee/go-echo/pkg/pagination"
)

const messagePreviewMaxLen = 100

// SendChatMessageServices — find-or-create conversation (social_account_id, customer_id), ບັນທຶກ message,
// ອັບເດດ last_message_preview/updated_at ສະເໝີ, unread_count+1 ສະເພາະຕອນ sender_type=CUSTOMER — ທັງໝົດໃນ transaction ດຽວ
func SendChatMessageServices(ctx context.Context, req requestbody.SendChatMessageRequest) error {
	tx := dbpkg.GetTransactionManager()
	return tx.WithTransaction(ctx, func(ctx context.Context) error {
		db := dbpkg.GetDBFromContext(ctx)

		conversation, err := dbquery.FindActiveConversationForUpdate(ctx, db, req.SocialAccountID, req.CustomerID)
		if err != nil {
			return err
		}
		var conversationID int64
		if conversation == nil {
			newID, err := dbinserts.CreateChatConversation(ctx, db, req.SocialAccountID, req.CustomerID)
			if err != nil {
				return err
			}
			conversationID = newID
		} else {
			conversationID = int64(conversation.ID)
		}

		messageType := req.MessageType
		if messageType == "" {
			messageType = "TEXT"
		}
		if err := dbinserts.CreateChatMessage(ctx, db, int(conversationID), req.SenderType, messageType, req.MessageBody, req.AttachmentURL); err != nil {
			return err
		}

		preview := req.MessageBody
		if len(preview) > messagePreviewMaxLen {
			preview = preview[:messagePreviewMaxLen]
		}
		incrementUnread := req.SenderType == "CUSTOMER"
		return dbupdate.UpdateConversationAfterMessage(ctx, db, conversationID, preview, incrementUnread)
	})
}

func GetDataChatMessageServices(ctx context.Context, id *int64, conversationID *int, page, pageSize int) ([]dbschema.ChatMessageDBSchema, *pagination.PaginationResult, error) {
	var paginationParam *pagination.PaginationParams
	if page > 0 || pageSize > 0 {
		paginationParam = pagination.NewPaginationParams(page, pageSize)
	}
	items, result, err := dbquery.GetChatMessageDataQuery(ctx, id, conversationID, paginationParam)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get data: %w", err)
	}
	return items, result, nil
}
