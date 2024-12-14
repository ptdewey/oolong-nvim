package internal

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/neovim/go-client/nvim"
)

// TODO: allow custom urls (intake from main.go via lua)
var baseURL = "http://localhost:11975"

func SearchHandler(v *nvim.Nvim, s string) error {
	args := strings.Split(s, " ")
	if len(args) < 2 {
		return errors.New("Not enough arguments.")
	}

	switch args[0] {
	case "keyword":
		res, err := keywordSearch(args[1:]...)
		if err != nil {
			return err
		}

		data := formatKeywordSearchResp(res)
		if err := createPopup(v, data); err != nil {
			return err
		}

		return nil
	case "note":
		res, err := noteSearch(args[1:]...)
		if err != nil {
			return err
		}

		data := formatNoteSearchResp(res)
		if err := createPopup(v, data); err != nil {
			return err
		}

		return nil
	default:
		return errors.New("Invalid search type: " + args[0])
	}
}

func createPopup(v *nvim.Nvim, data [][]byte) error {
	buf, err := CreateBuffer(v, data)
	if err != nil {
		return err
	}

	if err := createPopupWindow(v, buf); err != nil {
		return err
	}

	return nil
}

func keywordSearch(kw ...string) (KeywordResp, error) {
	data := KeywordResp{}

	qp := strings.Join(kw, "%20")
	resp, err := http.Get(baseURL + "/search/keyword?keyword=" + qp)
	if err != nil {
		return data, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return data, err
	}

	return data, nil
}

func noteSearch(note ...string) (NoteResp, error) {
	data := NoteResp{}

	qp := strings.Join(note, "%20")
	resp, err := http.Get(baseURL + "/search/note?path=" + qp)
	if err != nil {
		return data, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return data, err
	}

	return data, nil
}
