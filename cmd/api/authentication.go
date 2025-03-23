package api

import (
	"net/http"

	"github.com/nimilgp/URLcommentary/internal/dblayer"
	"github.com/nimilgp/URLcommentary/internal/request"
	"github.com/nimilgp/URLcommentary/internal/response"
	"github.com/nimilgp/URLcommentary/internal/validator"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

func (s *APIServer) postSignUp(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Username  string              `json: Username`
		FullName  string              `json: Userid`
		Emailid   string              `json: Emailid`
		Validator validator.Validator `json: -`
	}

	err := request.DecodeJSON(w, r, &input)
	if err != nil {
		s.badRequest(w, r, err)
		return
	}

	input.Validator.CheckField(input.Username != "", "Username", "Username required")
	input.Validator.CheckField(input.FullName != "", "Fullname", "Fullname required")
	input.Validator.CheckField(input.Emailid != "", "Email", "Email required")
	valid := validator.IsEmail(input.Emailid)
	if !valid {
		input.Validator.AddError("Not a valid email")
	}
	if input.Validator.HasErrors() {
		s.failedValidation(w, r, input.Validator)
		return
	}
	exist, err := s.querier.DoesUserExist(s.ctx, input.Emailid)
	if err != nil {
		s.logger.Warn("user existance could not be verified")
		s.serverError(w, r, err)
	}

	if exist {
		http.Redirect(w, r, "http://localhost:3333/signin", http.StatusConflict)
	} else {
		key, err := totp.Generate(totp.GenerateOpts{
			Issuer:      "URLcommentary",
			AccountName: input.Emailid,
			Period:      35,
			Digits:      otp.DigitsSix,
			Algorithm:   otp.AlgorithmSHA256,
		})
		arg := dblayer.CreateUserParams{
			Username:     input.Username,
			Fullname:     input.FullName,
			Emailid:      input.Emailid,
			Passwordhash: key.Secret(),
		}
		err = s.querier.CreateUser(s.ctx, arg)
		if err != nil {
			s.serverError(w, r, err)
		}
		err = response.JSON(w, http.StatusOK, envelope{"generator-qr-code": key.URL()})
		if err != nil {
			s.serverError(w, r, err)
		}
	}

}

func (s *APIServer) postSignIn(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Emailid   string              `json: Emailid`
		Passcode  string              `json: Passcode`
		Validator validator.Validator `json: -`
	}

	err := request.DecodeJSON(w, r, &input)
	if err != nil {
		s.badRequest(w, r, err)
		return
	}

	input.Validator.CheckField(input.Emailid != "", "Email", "Email required")
	input.Validator.CheckField(input.Passcode != "", "Passcode", "Passcode required")
	valid := validator.IsEmail(input.Emailid)
	if !valid {
		input.Validator.AddError("Not a valid email")
	}
	if input.Validator.HasErrors() {
		s.failedValidation(w, r, input.Validator)
		return
	}
	exist, err := s.querier.DoesUserExist(s.ctx, input.Emailid)
	if err != nil {
		s.logger.Warn("user existance could not be verified")
		s.serverError(w, r, err)
	}

	if !exist {
		http.Redirect(w, r, "http://localhost:3333/signup", http.StatusConflict)
	} else {
		secret, err := s.querier.RetrivePasswordHash(s.ctx, input.Emailid)
		if err != nil {
			s.serverError(w, r, err)
		}

		valid := totp.Validate(input.Passcode, secret)

		if valid {
			err = response.JSON(w, http.StatusOK, envelope{"login": "success"})
			if err != nil {
				s.serverError(w, r, err)
			}
		} else {
			err = response.JSON(w, http.StatusOK, envelope{"login": "failed"})
			if err != nil {
				s.serverError(w, r, err)
			}
		}

	}

}
