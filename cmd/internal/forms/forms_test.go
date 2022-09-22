package forms

import (
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/test", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid when should have been valid")
	} 
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/test", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("form shows valid when required fields missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "b")
	postedData.Add("c", "a")

	r = httptest.NewRequest("POST", "/test", nil)

	r.PostForm = postedData
	form = New(r.PostForm) 
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("form shows does not have required fields when it does")
	}
}