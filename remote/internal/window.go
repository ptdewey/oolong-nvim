package internal

import "github.com/neovim/go-client/nvim"

func createPopupWindow(v *nvim.Nvim, buf nvim.Buffer) error {
	var screen_size [2]int
	err := v.Eval("[&columns, &lines]", &screen_size)
	if err != nil {
		return err
	}

	popupWidth := int(0.85 * float64(screen_size[0]))
	popupHeight := int(0.85 * float64(screen_size[1]))
	_, err = v.OpenWindow(buf, true, &nvim.WindowConfig{
		Relative: "editor",
		Row:      float64((screen_size[1])-popupHeight)/2 - 1,
		Col:      float64((screen_size[0])-popupWidth) / 2,
		Width:    popupWidth,
		Height:   popupHeight,
		Style:    "minimal",
		Border:   "rounded",
		ZIndex:   50,
	})
	if err != nil {
		return err
	}

	return nil
}
