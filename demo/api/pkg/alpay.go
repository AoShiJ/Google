package pkg

import (
	"github.com/smartwalle/alipay/v3"
	"github.com/spf13/viper"
	"net/url"
)

func getClient() (*alipay.Client, error) {
	viper.SetConfigFile("./etc/api-api.yaml")
	viper.ReadInConfig()
	var client, err = alipay.New(viper.GetString("Alpay.AppID"), viper.GetString("Alpay.PrivateKey"), viper.GetBool("Alpay.IsProduction"))
	if err != nil {
		return nil, err
	}

	err = client.LoadAliPayPublicKey(viper.GetString("Alpay.PublicKey"))
	if err != nil {
		return nil, err
	}
	return client, err
}
func GetWebPayUrl(title, orderNO, amount string) (string, error) {
	viper.SetConfigFile("./etc/api-api.yaml")
	viper.ReadInConfig()
	var p = alipay.TradePagePay{
		Trade: alipay.Trade{
			NotifyURL:   viper.GetString("Alpay.NotifyURL"),
			ReturnURL:   viper.GetString("Alpay.ReturnURL"),
			Subject:     title,
			OutTradeNo:  orderNO,
			TotalAmount: amount,
			ProductCode: "FAST_INSTANT_TRADE_PAY",
		},
	}
	cli, err := getClient()
	if err != nil {
		return "", err
	}

	payUrl, err := cli.TradePagePay(p)
	if err != nil {
		return "", err
	}
	return payUrl.String(), nil
}
func VerifySign(values url.Values) (*alipay.Notification, error) {
	cli, err := getClient()
	if err != nil {
		return nil, err
	}

	return cli.DecodeNotification(values)
}
