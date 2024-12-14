# Oolong-nvim

[Oolong](https://github.com/oolong-sh/oolongd) integration for Neovim


## Installation

Oolong-nvim requires you have Go installed.

With lazy.nvim:
```lua
{
    "ptdewey/oolong-nvim",
    config = function()
        require("oolong").setup({})
    end,
}
```

## Usage

Open your Oolong graph in your browser: `:OolongGraph`

Oolong-nvim also provides a search command that takes in a keyword or note path.
Keyword searches show global weight and count, and a list of all documents the keyword occurs in, along with the number of occurrence and the in-document weight.
Note searches take in a file path and shows all keywords in the document, along with their weights.

Keyword search:
`:OolongSearch keyword <search_term>`

Note search:
`:OolongSearch note </path/to/your/note>`


To rebuild the Oolong-nvim binary (recommended each time the plugin updates): `:OolongRebuild`

## Configuration

TODO


## Disclaimer

This plugin is very early in development and features are missing/subject to change.

