// credentials
package model

func init() {

}

type PayliveCredentials struct {
	merchantEmail string
	merchantKey   string
}

func (c PayliveCredentials) MerchantEmail() string {
	return c.merchantEmail
}

func (c *PayliveCredentials) SetMerchantEmail(merchantEmail string) {
	c.merchantEmail = merchantEmail
}

func (c PayliveCredentials) MerchantKey() string {
	return c.merchantKey
}

func (c *PayliveCredentials) SetMerchantKey(merchantKey string) {
	c.merchantKey = merchantKey
}
