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

func TestForm_Has(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)

	has := form.Has("whatever")
	if has {
		t.Error("form shows has field when it does not")
	}
	
	postedData = url.Values{} 
	postedData.Add("a", "a")
	form = New(postedData)

	has = form.Has("a")
	if !has {
		t.Error("form show does not have field when it does")
	}
}

func TestForm_MinLength(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)

	form.MinLength("x", 10)
	if form.Valid() {
		t.Error("form shows min length for non existent field")
	}

	isError := form.Errors.Get("x")
	if isError == "" {
		t.Error("should have an error, but did not get one")
	}

	postedData = url.Values{}
	postedData.Add("a", "some_long_value")

	form = New(postedData)

	form.MinLength("a", 100)
	if form.Valid() {
		t.Error("form shows min length of 100 met when the data is shorter")
	}

	postedData = url.Values{}
	postedData.Add("a", "another_long_field")

	form = New(postedData)
	form.MinLength("a", 1)
	if !form.Valid() {
		t.Error("form shows min length of 1 is not met when it is")
	}

	isError = form.Errors.Get("a")
	if isError != "" {
		t.Error("should not have an error, but did get one")
	}
}

func TestForm_IsEmail(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)

	form.IsEmail("x")
	if form.Valid() {
		t.Error("form shows valid email for not-existent field")
	}

	postedData = url.Values{}
	postedData.Add("email", "me@here.com")
	form = New(postedData)
	form.IsEmail("email")
	if !form.Valid() {
		t.Error("got an invalid email when we should have")
	}

	postedData = url.Values{}
	postedData.Add("email", "x")
	form = New(postedData)
	form.IsEmail("email")
	if form.Valid() {
		t.Error("got valid email when we it should not have")
	}
}