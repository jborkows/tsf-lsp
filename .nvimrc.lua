print("Haalo")

vim.cmd([[ autocmd BufRead,BufNewFile *.tsf set filetype=tsf ]])
vim.api.nvim_buf_set_option(0, "filetype", "tsf")

local client = vim.lsp.start_client({
	cmd = { "./tmp/main" },
	name = "tsf-lsp",
	on_attach = require("custom.lspFun").attach,
})

if not client then
	print("Failed to start client")
end

-- use lua api for add filetype
--

vim.api.nvim_create_autocmd("FileType", {
	group = vim.api.nvim_create_augroup("kickstart-highlight-yank", { clear = true }),
	pattern = { "tsf" },
	callback = function()
		print("tsf file detected - autocmd")
		local data = {
			buf = vim.fn.expand("<abuf>"),
			fileType = vim.fn.expand("<amatch>"),
			fileName = vim.fn.expand("<afile>"),
		}

		vim.schedule(function()
			print("registering client for tsf file")
			print(vim.inspect(data))
			vim.lsp.buf_attach_client(0, client)
		end)
	end,
})
