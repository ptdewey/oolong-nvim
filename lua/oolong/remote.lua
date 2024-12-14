local M = {}

local chan
local bin_path
local plugin_path

local options = {}

--- job runner for oolong-nvim remote binary
---@return integer?
local function ensure_job()
    if chan then
        return chan
    end

    if not bin_path then
        print("Error: oolong-nvim binary not found.")
        return
    end

    chan = vim.fn.jobstart({ bin_path }, {
        rpc = true,
        on_exit = function(_, code, _)
            if code ~= 0 then
                print("Error: oolong-nvim job exited with code " .. code)
                chan = nil
            end
        end,
        on_stderr = function(_, data, _)
            for _, line in ipairs(data) do
                if line ~= "" then
                    print("stderr: " .. line)
                end
            end
        end,
        on_stdout = function(_, data, _)
            for _, line in ipairs(data) do
                if line ~= "" then
                    print("stdout: " .. line)
                end
            end
        end,
    })

    if not chan or chan == 0 then
        error("Failed to start oolong-nvim job")
    end

    return chan
end

--- Create plugin user commands to build binary and show report
local function setup_oolong_commands()
    vim.api.nvim_create_user_command("OolongSearch", function(args)
        chan = ensure_job()
        if not chan or chan == 0 then
            print("Error: Invalid channel")
            return
        end

        local success, result =
            pcall(vim.fn.rpcrequest, chan, "oolong-search", args.args)
        if not success then
            print("RPC request failed: " .. result)
        end
    end, { nargs = "?" })

    vim.api.nvim_create_user_command("OolongRebuild", function()
        print("Rebuilding oolong-nvim binary with Go...")
        local result =
            os.execute("cd " .. plugin_path .. "remote" .. " && go build")
        if result == 0 then
            print("Go binary compiled successfully.")
            if chan then
                vim.fn.jobstop(chan)
                chan = nil
            end
        else
            print("Failed to compile Go binary.")
        end
    end, { nargs = 0 })
end

--- Report generation setup (requires go)
---@param opts table
function M.setup(opts)
    -- Get plugin install path
    plugin_path = debug.getinfo(1).source:sub(2):match("(.*/).*/.*/")

    -- Check os to switch separators and binary extension if necessary
    local uname = vim.loop.os_uname().sysname
    local path_separator = (uname == "Windows_NT") and "\\" or "/"
    bin_path = plugin_path
        .. "remote"
        .. path_separator
        .. "oolong-nvim"
        .. (uname == "Windows_NT" and ".exe" or "")

    setup_oolong_commands()

    -- Check if binary exists
    local uv = vim.loop
    local handle = uv.fs_open(bin_path, "r", 438)
    if handle then
        uv.fs_close(handle)
        return
    end

    -- Compile binary if it doesn't exist
    print(
        "oolong-nvim binary not found at "
            .. bin_path
            .. ", attempting to compile..."
    )

    local result =
        os.execute("cd " .. plugin_path .. "remote" .. " && go build")
    if result == 0 then
        print("Go binary compiled successfully.")
    else
        print("Failed to compile Go binary." .. uv.cwd())
    end
end

return M
