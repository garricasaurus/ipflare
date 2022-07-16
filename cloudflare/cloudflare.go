package cloudflare

// CloudFlare implements Updater and uses the CloudFlare API to modify
// a DNS record identified by its zone name and record name.
type CloudFlare struct {
	client client
	zone   string
	record string
}

// NewCloudFlare creates a new CloudFlare instance.
// The auth token is obtained from the os environment.
func NewCloudFlare(zone, record, authToken string) *CloudFlare {
	return &CloudFlare{
		zone:   zone,
		record: record,
		client: &httpClient{
			authToken: authToken,
		},
	}
}

// Update the content of the associated record.
func (u *CloudFlare) Update(content string) error {
	z, err := u.client.GetZone(u.zone)
	if err != nil {
		return err
	}
	r, err := u.client.GetRecord(z.Id, u.record)
	if err != nil {
		return err
	}
	r.Content = content
	_, err = u.client.UpdateRecord(z.Id, r)
	return err
}
