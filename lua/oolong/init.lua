local M = {}

local remote = require("oolong.remote")

local default_opts = {
    open_command = nil,
    graph_url = "http://localhost:11975",
}

local open_command = function()
    local uname = vim.loop.os_uname().sysname
    if uname == "Linux" then
        return "xdg-open"
    elseif uname == "Darwin" then
        return "open"
    elseif uname == "Windows_NT" then
        return "start"
    else
        error("Open command for uname '" .. uname("' not found."))
    end
end

---@param opts table?
function M.setup(opts)
    opts = vim.tbl_deep_extend("force", default_opts, opts or {})

    vim.api.nvim_create_user_command("OolongGraph", function()
        vim.fn.jobstart(
            { opts.open_command or open_command(), opts.graph_url },
            { detach = true }
        )
    end, {})

    remote.setup(opts)
end

return M
