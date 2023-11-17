package repository

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/morkid/paginate"
	"github.com/pranavparaswar/taskmanager/database/migrations"
	"github.com/pranavparaswar/taskmanager/database/models"
	"gopkg.in/go-playground/validator.v9"
)

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

var validate = validator.New()

func ValidateStruct(task models.Task) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(task)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func (r *Repository) CreateTask(context *fiber.Ctx) error {
	task := models.Task{}
	err := context.BodyParser(&task)
	fmt.Printf("Request Body: %+v\n", task)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "Request failed"})

		return err
	}

	errors := ValidateStruct((task))

	if errors != nil {
		return context.Status(fiber.StatusBadRequest).JSON(errors)
	}

	if err := r.DB.Create(&task).Error; err != nil {
		return context.Status(http.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Coudlnt create user", "data": err})
	}
	context.Status(http.StatusCreated).JSON(&fiber.Map{"message": "User has been added", "data": task})
	return nil

}

func (r *Repository) EditTask(context *fiber.Ctx) error {
	task := models.Task{}
	err := context.BodyParser(&task)
	fmt.Printf("Received Task Data: %+v\n", task)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "Request Failed"})

		return err
	}

	errors := ValidateStruct(task)
	if errors != nil {
		return context.Status(fiber.StatusBadRequest).JSON(errors)
	}

	fmt.Println("reached here")
	db := r.DB
	taskIDStr := context.Params("id")
	fmt.Println("this is task_id", taskIDStr)

	// Parse the task_id as an integer
	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": "Invalid task_id"})
		return nil
	}

	if taskID == 0 {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": "task_id Cannot be nil"})
		return nil
	}

	// Find the task by task_id
	existingTask := models.Task{}
	if err := db.Where("task_id = ?", taskID).First(&existingTask).Error; err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "Could not get task with given task_id"})
		return err
	}

	// Update only the specified fields
	if err := db.Model(&existingTask).Where("task_id = ?", taskID).Updates(models.Task{
		TaskName: task.TaskName,
		Note:     task.Note,
		Deadline: task.Deadline,
	}).Error; err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "Could not update task", "data": err})
		return err
	}

	return context.JSON(fiber.Map{"status": "success", "message": "Task Successfully updated"})
}

func (r *Repository) DeleteTask(context *fiber.Ctx) error {
	taskModel := migrations.Task{}

	id := context.Params("id")

	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": "ID Cannot be empty"})
		return nil
	}

	err := r.DB.Delete(taskModel, id)

	if err.Error != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "Could not be deleted"})
		return err.Error
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "user deleted succesfully"})
	return nil

}

func (r *Repository) GetTask(context *fiber.Ctx) error {
	db := r.DB
	model := db.Model(&migrations.Task{})

	pg := paginate.New(&paginate.Config{
		DefaultSize:        5,
		CustomParamEnabled: true,
	})

	page := pg.With(model).Request(context.Request()).Response(&[]migrations.Task{})

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"data": page,
	})

	return nil
}

func (r *Repository) GetTaskByID(context *fiber.Ctx) error {
	id := context.Params("id")
	taskModel := &migrations.Task{}
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": "ID cannot be empty"})
		return nil
	}

	err := r.DB.Where("id = ?", id).First(taskModel).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "Could not get the task"})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "Task id fetched successfully", "data": taskModel})
	return nil
}
