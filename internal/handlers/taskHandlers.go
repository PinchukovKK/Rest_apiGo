package handlers

import (
	"context"

	"main.go/internal/taskService"
	"main.go/internal/web/tasks"
)

type taskHandler struct {
	Service *taskService.TaskService
}

func NewTaskHandler(service *taskService.TaskService) *taskHandler {
	return &taskHandler{Service: service}
}

func (h taskHandler) GetTasksByUserId(_ context.Context, _ tasks.GetTasksByUserIdRequestObject) (tasks.GetTasksByUserIdResponseObject, error) {
	allTasks, err := h.Service.GetTaskByUserId()
	if err != nil {
		return nil, err
	}

	response := tasks.GetTasksByUserId200JSONResponse{}
	
	for _, tsk := range allTasks {
		task := tasks.Task{
			UserId: &tsk.UserId,
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
		}
		response = append(response, task)
	}

	return response, nil
}

func (h taskHandler) PostTasks(ctx context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	taskRequest := request.Body

	taskToCreate := taskService.Task{
		UserId: uint(*taskRequest.UserId),
		Task:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
	}
	createdTask, err := h.Service.CreateTask(taskToCreate, taskToCreate.UserId)

	if err != nil {
		return nil, err
	}

	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Task,
		IsDone: &createdTask.IsDone,
		UserId: &createdTask.UserId,
	}

	return response, nil
}

func (h taskHandler) PatchTasksId(ctx context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	taskID := request.Id
	taskRequestUpdate := request.Body

	taskToUpdate := taskService.Task{
		Task:   *taskRequestUpdate.Task,
		IsDone: *taskRequestUpdate.IsDone,
	}

	updatedTask, err := h.Service.UpdateTask(taskToUpdate, uint(taskID))
	if err != nil {
		return nil, err
	}

	response := tasks.PatchTasksId200JSONResponse{
		Id:     &updatedTask.ID,
		Task:   &updatedTask.Task,
		IsDone: &updatedTask.IsDone,
	}

	return response, nil
}

func (h taskHandler) DeleteTasksId(ctx context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	taskID := request.Id

	err := h.Service.DeleteTask(uint(taskID))
	if err != nil {
		return nil, err
	}

	return tasks.DeleteTasksId204Response{}, nil
}


