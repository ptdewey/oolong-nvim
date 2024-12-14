package internal

import "github.com/neovim/go-client/nvim"

func CreateBuffer(v *nvim.Nvim, data [][]byte) (nvim.Buffer, error) {
	buf, err := v.CreateBuffer(false, true)
	if err != nil {
		return buf, err
	}

	if err := v.SetBufferOption(buf, "filetype", "markdown"); err != nil {
		return buf, err
	}

	if err := v.SetBufferLines(buf, 0, -1, false, data); err != nil {
		return buf, err
	}

	kopts := map[string]bool{
		"silent": true,
	}
	if err := v.SetBufferKeyMap(buf, "n", "q", "<cmd>close!<CR>", kopts); err != nil {
		return buf, err
	}

	return buf, err
}
