package github

import (
	"context"
	"fmt"
)

type PublicKey struct {
	KeyID *string `json:"key_id"`
	Key   *string `json:"key"`
}

func (s *ActionsService) GetRepoPublicKey(ctx context.Context, owner, repo string) (*PublicKey, *Response, error) {
	u := fmt.Sprintf("repos/%v/%v/actions/secrets/public-key", owner, repo)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	pubKey := new(PublicKey)
	resp, err := s.client.Do(ctx, req, pubKey)
	if err != nil {
		return nil, resp, err
	}

	return pubKey, resp, nil
}

func (s *ActionsService) GetOrgPublicKey(ctx context.Context, org string) (*PublicKey, *Response, error) {
	u := fmt.Sprintf("orgs/%v/actions/secrets/public-key", org)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	pubKey := new(PublicKey)
	resp, err := s.client.Do(ctx, req, pubKey)
	if err != nil {
		return nil, resp, err
	}

	return pubKey, resp, nil
}

type Secret struct {
	Name                    string    `json:"name"`
	CreatedAt               Timestamp `json:"created_at"`
	UpdatedAt               Timestamp `json:"updated_at"`
	Visibility              string    `json:"visibility,omitempty"`
	SelectedRepositoriesURL string    `json:"selected_repositories_url,omitempty"`
}

type Secrets struct {
	TotalCount int       `json:"total_count"`
	Secrets    []*Secret `json:"secrets"`
}

func (s *ActionsService) ListRepoSecrets(ctx context.Context, owner, repo string, opts *ListOptions) (*Secrets, *Response, error) {
	u := fmt.Sprintf("repos/%v/%v/actions/secrets", owner, repo)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	secrets := new(Secrets)
	resp, err := s.client.Do(ctx, req, &secrets)
	if err != nil {
		return nil, resp, err
	}

	return secrets, resp, nil
}

func (s *ActionsService) GetRepoSecret(ctx context.Context, owner, repo, name string) (*Secret, *Response, error) {
	u := fmt.Sprintf("repos/%v/%v/actions/secrets/%v", owner, repo, name)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	secret := new(Secret)
	resp, err := s.client.Do(ctx, req, secret)
	if err != nil {
		return nil, resp, err
	}

	return secret, resp, nil
}


type SelectedRepoIDs []int64
type EncryptedSecret struct {
	Name                  string          `json:"-"`
	KeyID                 string          `json:"key_id"`
	EncryptedValue        string          `json:"encrypted_value"`
	Visibility            string          `json:"visibility,omitempty"`
	SelectedRepositoryIDs SelectedRepoIDs `json:"selected_repository_ids,omitempty"`
}

func (s *ActionsService) CreateOrUpdateRepoSecret(ctx context.Context, owner, repo string, eSecret *EncryptedSecret) (*Response, error) {
	u := fmt.Sprintf("repos/%v/%v/actions/secrets/%v", owner, repo, eSecret.Name)

	req, err := s.client.NewRequest("PUT", u, eSecret)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

func (s *ActionsService) DeleteRepoSecret(ctx context.Context, owner, repo, name string) (*Response, error) {
	u := fmt.Sprintf("repos/%v/%v/actions/secrets/%v", owner, repo, name)

	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

func (s *ActionsService) ListOrgSecrets(ctx context.Context, org string, opts *ListOptions) (*Secrets, *Response, error) {
	u := fmt.Sprintf("orgs/%v/actions/secrets", org)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	secrets := new(Secrets)
	resp, err := s.client.Do(ctx, req, &secrets)
	if err != nil {
		return nil, resp, err
	}

	return secrets, resp, nil
}

func (s *ActionsService) GetOrgSecret(ctx context.Context, org, name string) (*Secret, *Response, error) {
	u := fmt.Sprintf("orgs/%v/actions/secrets/%v", org, name)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	secret := new(Secret)
	resp, err := s.client.Do(ctx, req, secret)
	if err != nil {
		return nil, resp, err
	}

	return secret, resp, nil
}

func (s *ActionsService) CreateOrUpdateOrgSecret(ctx context.Context, org string, eSecret *EncryptedSecret) (*Response, error) {
	u := fmt.Sprintf("orgs/%v/actions/secrets/%v", org, eSecret.Name)

	req, err := s.client.NewRequest("PUT", u, eSecret)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

type SelectedReposList struct {
	TotalCount   *int          `json:"total_count,omitempty"`
	Repositories []*Repository `json:"repositories,omitempty"`
}

func (s *ActionsService) ListSelectedReposForOrgSecret(ctx context.Context, org, name string) (*SelectedReposList, *Response, error) {
	u := fmt.Sprintf("orgs/%v/actions/secrets/%v/repositories", org, name)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	result := new(SelectedReposList)
	resp, err := s.client.Do(ctx, req, result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

func (s *ActionsService) SetSelectedReposForOrgSecret(ctx context.Context, org, name string, ids SelectedRepoIDs) (*Response, error) {
	u := fmt.Sprintf("orgs/%v/actions/secrets/%v/repositories", org, name)

	type repoIDs struct {
		SelectedIDs SelectedRepoIDs `json:"selected_repository_ids,omitempty"`
	}

	req, err := s.client.NewRequest("PUT", u, repoIDs{SelectedIDs: ids})
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

func (s *ActionsService) AddSelectedRepoToOrgSecret(ctx context.Context, org, name string, repo *Repository) (*Response, error) {
	u := fmt.Sprintf("orgs/%v/actions/secrets/%v/repositories/%v", org, name, *repo.ID)
	req, err := s.client.NewRequest("PUT", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

func (s *ActionsService) RemoveSelectedRepoFromOrgSecret(ctx context.Context, org, name string, repo *Repository) (*Response, error) {
	u := fmt.Sprintf("orgs/%v/actions/secrets/%v/repositories/%v", org, name, *repo.ID)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

func (s *ActionsService) DeleteOrgSecret(ctx context.Context, org, name string) (*Response, error) {
	u := fmt.Sprintf("orgs/%v/actions/secrets/%v", org, name)

	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}
