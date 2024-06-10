package forms

import (
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	isValid := form.Valid()

	if !isValid {
		t.Error("got invalid when should have been valid")
	}
}

func TestForm_Required(t *testing.T) {
	postedData := url.Values{}
	postedData.Add("g", "g")

	form := New(postedData)

	form.Required("a", "b", "c")

	if form.Valid() {
		t.Error("form shows valid when required fields missing")
	}

	postedData = url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "a")
	postedData.Add("c", "a")

	form = New(postedData)

	form.Required("a", "b", "c")

	if !form.Valid() {
		t.Error("shows that it does not have required fields when it does")
	}
}

func TestForm_IsEmail(t *testing.T) {
	postedData := url.Values{}
	postedData.Add("tralalalalal", "test@test.si")
	form := New(postedData)

	form.IsEmail("x")

	if form.Valid() {
		t.Error("This form is valid even though it is non existant")
	}

	postedData = url.Values{}
	postedData.Add("email", "test@test.si")
	form = New(postedData)

	form.IsEmail("email")

	if !form.Valid() {
		t.Error("form is not valid but the email is valid")
	}

	postedData = url.Values{}
	postedData.Add("email", "testtest.si")
	form = New(postedData)

	form.IsEmail("email")

	if form.Valid() {
		t.Error("form is valid but the email is not valid")
	}
}

func TestForm_MinLength(t *testing.T) {
	postedData := url.Values{}
	postedData.Add("some_field", "some value")

	form := New(postedData)

	form.MinLength("test", 5)

	if form.Valid() {
		t.Error("form shows minlen for non exsistant field")
	}

	isError := form.Errors.Get("test")

	if isError == "" {
		t.Error("should have an error but did not get one")
	}

	postedData = url.Values{}
	postedData.Add("some_field", "some value")

	form = New(postedData)

	form.MinLength("some_field", 5)

	if !form.Valid() {
		t.Error("form field does not have enough characters")
	}

	isError = form.Errors.Get("some_field")

	if isError != "" {
		t.Error("should not have an error but got one")
	}

	postedData = url.Values{}
	postedData.Add("my_new_field", "some nice value")

	form = New(postedData)

	form.MinLength("my_new_field", 20)

	if form.Valid() {
		t.Error("form field should have enough characters but it doesn't")
	}

}

func TestForm_Has(t *testing.T) {
	postedData := url.Values{}
	postedData.Add("a", "a")
	form := New(postedData)
	isValid := form.Has("first")

	if isValid {
		t.Error("form is valid when is should not be")
	}

	postedData = url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "a")

	form = New(postedData)

	isValid = form.Has("a")

	if !isValid {
		t.Error("form is invalid when it should be valid")
	}

}
