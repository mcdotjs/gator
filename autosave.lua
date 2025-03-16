local bufferNum = 31

vim.api.nvim_create_autocmd("BufWritePost", {
	group = vim.api.nvim_create_augroup("MiroSavingGo", { clear = true }),
	pattern = "*.go",
	callback = function()
		vim.fn.jobstart({ "go", "run", ".", "login", "bootssdsdsd" }, {
			stdout_buffered = true,
			on_stdout = function(_, data)
				if data then
					vim.api.nvim_buf_set_lines(bufferNum, -1, -1, false, data)
				end
			end,
			on_stderr = function(_, data)
				if data then
					vim.api.nvim_buf_set_lines(bufferNum, -1, -1, false, data)
				end
			end,
		})
	end,
})
