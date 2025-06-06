package handler

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/miraklik/TODO-list/db"
)

type TaskHandler struct {
	TaskService *db.TaskService
}

type ErrorResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

type CreateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type SuccessResponse struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func NewTaskHandler() (*TaskHandler, error) {
	taskService, err := db.NewTodoService()
	if err != nil {
		log.Println("Failed to create new todo service: %v", err)
	}

	return &TaskHandler{
		TaskService: taskService,
	}, nil
}

func (h *TaskHandler) CreateTask(c *fiber.Ctx) error {
	var req CreateTaskRequest

	if err := c.BodyParser(&req); err != nil {
		log.Printf("Failed to parse request: %s", err)
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error:   true,
			Message: "Invalid request body",
		})
	}

	task, err := h.TaskService.CreateTodo(req.Title, req.Description)
	if err != nil {
		log.Printf("Failed to create task: %v", err)

		if err.Error() == "task not found" {
			return c.Status(fiber.StatusNotFound).JSON(ErrorResponse{
				Error:   true,
				Message: "Task not found",
			})
		}

		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error:   true,
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(SuccessResponse{
		Error:   false,
		Message: "Task created",
		Data:    task,
	})
}

func (h *TaskHandler) GetAllTasks(c *fiber.Ctx) error {
	tasks, err := h.TaskService.GetAllTasks()
	if err != nil {
		log.Printf("Failed to get all task: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
			Error:   true,
			Message: "Failed to retrieve tasks",
		})
	}

	return c.Status(fiber.StatusOK).JSON(SuccessResponse{
		Error:   false,
		Message: "Tasks retrieved",
		Data:    tasks,
	})
}

func (h *TaskHandler) UpdateTaks(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("Failed to get task id: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error:   true,
			Message: "Invalid task id",
		})
	}

	var req CreateTaskRequest

	if err := c.BodyParser(&req); err != nil {
		log.Printf("Failed to parse request body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error:   true,
			Message: "Invalid request body",
		})
	}

	task, err := h.TaskService.Update(id, req.Title, req.Description)
	if err != nil {
		log.Printf("Failed to update task: %v", err)

		if err.Error() == "task not found" {
			return c.Status(fiber.StatusNotFound).JSON(ErrorResponse{
				Error:   true,
				Message: "Task not found",
			})
		}

		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error:   true,
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(SuccessResponse{
		Error:   false,
		Message: "task updated",
		Data:    task,
	})
}

func (h *TaskHandler) DeleteTask(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("Failed to get task id: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error:   true,
			Message: "Invalid task id",
		})
	}

	if err := h.TaskService.Delete(id); err != nil {
		log.Printf("Failed to delete tasks: %v", err)

		if err.Error() == "task not found" {
			return c.Status(fiber.StatusNotFound).JSON(ErrorResponse{
				Error:   true,
				Message: "Task not found",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
			Error:   true,
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(SuccessResponse{
		Error:   false,
		Message: "Task deleted",
	})
}
