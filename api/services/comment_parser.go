package services

import (
	"context"
	"fmt"
	"strings"

	"github.com/SONEsee/go-echo/api/schema/requestbody"
	commentparser "github.com/SONEsee/go-echo/pkg/comment-parser"
	dbquery "github.com/SONEsee/go-echo/pkg/db-pkg/db-query"
)

// ParseAndCreateCommentIntentServices — ອ່ານ comments_raw.comment_message, ພະຍາຍາມຈັບຮູບແບບ "CF <code> x<qty>",
// ຈັບຄູ່ code ກັບ live_session_products.cf_code_override ຂອງ live session ນັ້ນ, ຄົ້ນຫາ customer ດ້ວຍ fb_user_id
// (ຖ້າມີ) ແລ້ວເອີ້ນ CreateCommentIntentServices ຕໍ່ (reuse transaction+reserve-stock logic ທັງໝົດ)
func ParseAndCreateCommentIntentServices(ctx context.Context, commentRawID int64) error {
	comments, _, err := dbquery.GetCommentRawDataQuery(ctx, &commentRawID, nil, nil, nil)
	if err != nil {
		return err
	}
	comment := comments[0]

	req := requestbody.CommentIntentRequestBody{CommentRawID: commentRawID}

	if comment.CommentMessage != nil {
		code, qty, matched := commentparser.ParseCFComment(*comment.CommentMessage)
		if matched {
			sessions, _, err := dbquery.GetLiveSessionDataQuery(ctx, &comment.LiveSessionID, nil)
			if err != nil {
				return err
			}
			session := sessions[0]

			socialAccounts, _, err := dbquery.GetSocialAccountDataQuery(ctx, &session.SocialAccountID, nil)
			if err != nil {
				return err
			}
			shopID := socialAccounts[0].ShopID

			products, _, err := dbquery.GetLiveSessionProductDataQuery(ctx, nil, &comment.LiveSessionID, nil)
			if err != nil {
				return err
			}
			for _, p := range products {
				if p.CfCodeOverride != nil && strings.EqualFold(*p.CfCodeOverride, code) {
					variantID := p.ProductVariantID
					req.MatchedProductVariantID = &variantID
					break
				}
			}
			parsedQty := qty
			req.ParsedQty = &parsedQty

			if comment.FbUserID != nil {
				customer, err := dbquery.GetCustomerBySocialPlatformID(ctx, shopID, *comment.FbUserID)
				if err != nil {
					return err
				}
				if customer != nil {
					customerID := customer.ID
					req.CustomerID = &customerID
				}
			}
		}
	}

	if err := CreateCommentIntentServices(ctx, req); err != nil {
		return fmt.Errorf("parse comment intent: %w", err)
	}
	return nil
}
