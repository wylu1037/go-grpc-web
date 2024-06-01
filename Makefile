install:
	@echo "🐌 Installing deps..."

gen_proto:
	@echo "Generate proto"
	@buf generate proto
	@echo "🏁 Generate completed."

gen_db:
	@echo "Generate database"
	@cd app/database && sqlc generate
	@echo "🏁 Generate completed."

buf_mod_update:
	@echo "buf mod update"
	@cd proto && buf mod update
	@echo "🏁 Buf mod update completed."