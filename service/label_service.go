package service

import "google.golang.org/api/gmail/v1"

func GetOrCreateLabel(srv *gmail.Service, user string, labelName string) (string, error) {
	labels, err := srv.Users.Labels.List(user).Do()
	if err != nil {
		return "", err
	}

	for _, l := range labels.Labels {
		if l.Name == labelName {
			return l.Id, nil
		}
	}

	newLabel := &gmail.Label{
		Name:                  labelName,
		LabelListVisibility:   "labelShow",
		MessageListVisibility: "show",
	}

	created, err := srv.Users.Labels.Create(user, newLabel).Do()
	if err != nil {
		return "", err
	}

	return created.Id, nil
}

func MarkEmailProcessed(srv *gmail.Service, user, msgID, labelID string) error {
	req := &gmail.ModifyMessageRequest{
		AddLabelIds: []string{labelID},
	}

	_, err := srv.Users.Messages.Modify(user, msgID, req).Do()
	return err
}
