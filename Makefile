install:
	@echo "🐌 Installing deps..."

gen_proto:
	@echo "Generate proto"
	@buf generate proto
	@echo "🏁 Generate completed."

gen_db:
	@echo "Generate database"
	@sqlc generate
	@echo "🏁 Generate completed."