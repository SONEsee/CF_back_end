package requestbody

type ChatConversationPatchRequest struct {
	AssignedStaffID *int    `json:"assigned_staff_id,omitempty" validate:"omitempty,gt=0"`
	Status          *string `json:"status,omitempty" validate:"omitempty,oneof=OPEN PENDING CLOSED"`
}
