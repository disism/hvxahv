package activity

//func (h *Handler) Undo(data []byte) (*pb.ActivityResponse, error) {
//	var (
//		notok []string
//		ok    []string
//		id    = fmt.Sprintf("%s/%s", h.aAddr, uuid.NewString())
//	)
//
//	var b Object
//
//	if err := json.Unmarshal(data, &b); err != nil {
//		return nil, err
//	}
//	body := &activitypub.Reject{
//		Context: "https://www.w3.org/ns/activitystreams",
//		Id:      id,
//		Type:    activitypub.UndoType,
//		Actor:   h.aAddr,
//		Object: struct {
//			Id     string `json:"id"`
//			Type   string `json:"type"`
//			Actor  string `json:"actor"`
//			Object string `json:"object"`
//		}{
//			Id:     b.Id,
//			Type:   b.Type,
//			Actor:  b.Actor,
//			Object: b.Object,
//		},
//	}
//	marshal, err := json.Marshal(body)
//	if err != nil {
//		return nil, err
//	}
//	// DELIVERY ...
//	do, err := NewDelivery(marshal, h.aAddr, h.privateKey).Do(fmt.Sprintf("%s/inbox", h.inbox))
//	if err != nil {
//		return nil, err
//	}
//	if do.StatusCode != 202 {
//		notok = append(notok, h.inbox)
//		return nil, nil
//	}
//	ok = append(ok, h.inbox)
//
//	if err := outbox.NewOutboxesDeleteByActivityId(h.actorId, b.Id).Delete(); err != nil {
//		return nil, err
//	}
//
//	return response(notok, ok)
//}
