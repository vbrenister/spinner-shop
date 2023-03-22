package main

import (
	"net/http"
)

func (app *application) VeritualTerminal(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["publishable_key"] = app.config.stripe.key
	if err := app.renderTemplate(w, r, "terminal", &templateData{
		StringMap: stringMap,
	}, "stripe-js"); err != nil {
		app.erroLog.Println(err)
	}
}

func (app *application) PaymentSucceeded(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.erroLog.Println(err)
		return
	}

	cardHolder := r.Form.Get("cardholder_name")
	paymentIntent := r.Form.Get("payment_intent")
	paymentMethod := r.Form.Get("payment_method")
	paymentAmount := r.Form.Get("payment_amount")
	paymentCurrency := r.Form.Get("payment_currency")
	cardHolderEmail := r.Form.Get("cardholder_email")

	data := make(map[string]interface{})

	data["cardholder"] = cardHolder
	data["email"] = cardHolderEmail
	data["pi"] = paymentIntent
	data["pm"] = paymentMethod
	data["pa"] = paymentAmount
	data["pc"] = paymentCurrency

	if err := app.renderTemplate(w, r, "succeeded", &templateData{
		Data: data,
	}); err != nil {
		app.erroLog.Println(err)
	}

}

func (app *application) ChargeOnce(w http.ResponseWriter, r *http.Request) {
	if err := app.renderTemplate(w, r, "buy-once", nil, "stripe-js"); err != nil {
		app.erroLog.Println(err)
	}
}
