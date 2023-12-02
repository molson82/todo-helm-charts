bb:
	@echo "Building backend..."
	@docker build -t todo-backend ./apps/backend/

bf:
	@echo "Building frontend..."
	@docker build -t todo-frontend ./apps/frontend/

db:
	@echo "Docker run backend..."
	@docker run --publish 8880:8880 todo-backend
