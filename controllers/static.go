package controllers

import (
	"net/http"

	"github.com/rabocse/lenslocked2/views"
)

func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}

func FAQ(tpl views.Template) http.HandlerFunc {

	questions := []struct {
		Question string
		Answer   string
	}{
		{
			Question: "Is there a free version?",
			Answer:   "Yes, there is a free version of Lenslocked. You can sign up for a free account and access a limited set of features.",
		},
		{
			Question: "What payment methods do you accept?",
			Answer:   "We accept all major credit cards, as well as PayPal.",
		},
		{
			Question: "Can I cancel my subscription?",
			Answer:   "Yes, you can cancel your subscription at any time. Just go to your account settings and click on the 'Cancel Subscription' button.",
		},
	}

	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, questions)
	}

}
