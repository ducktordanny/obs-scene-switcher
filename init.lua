local conf = require("telescope.config").values
local pickers = require("telescope.pickers")
local finders = require("telescope.finders")
local actions = require("telescope.actions")
local action_state = require("telescope.actions.state")
local themes = require("telescope.themes")

local M = {}

M.get_obs_scenes = function()
	return vim.fn.systemlist("go run .")
end

M.select_obs_scene = function(scene_name)
	-- TODO: Should handle obs not opened state
	vim.fn.system("go run . -n " .. scene_name)
end

M.obs_scene_selector_ui = function(opts)
	opts = opts or {}
	local obs_scenes = M.get_obs_scenes()

	pickers
		.new(opts, {
			prompt_title = "OBS Scenes",
			finder = finders.new_table({
				results = obs_scenes,
			}),
			sorter = conf.generic_sorter(opts),
			attach_mappings = function(prompt_bufnr, _)
				actions.select_default:replace(function()
					actions.close(prompt_bufnr)
					local selection = action_state.get_selected_entry()
					M.select_obs_scene(obs_scenes[selection.index])
				end)
				return true
			end,
		})
		:find()
end

vim.keymap.set("n", "<leader>os", function()
	M.obs_scene_selector_ui(themes.get_dropdown({}))
end, { desc = "[O]BS [S]cene selector" })

return M
