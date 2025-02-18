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

func (h taskHandler) GetTasks(_ context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := h.Service.GetAllTask()
	if err != nil {
		return nil, err
	}

	response := tasks.GetTasks200JSONResponse{}

	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
		}
		response = append(response, task)
	}

	return response, nil
}

func (h taskHandler) PostTasksPost(ctx context.Context, request tasks.PostTasksPostRequestObject) (tasks.PostTasksPostResponseObject, error) {
	taskRequest := request.Body

	taskToCreate := taskService.Task{
		Task:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
	}
	createdTask, err := h.Service.CreateTask(taskToCreate)

	if err != nil {
		return nil, err
	}

	response := tasks.PostTasksPost201JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Task,
		IsDone: &createdTask.IsDone,
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

func (h taskHandler) PatchTasksPatchId(ctx context.Context, request tasks.PatchTasksPatchIdRequestObject) (tasks.PatchTasksPatchIdResponseObject, error) {
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

	response := tasks.PatchTasksPatchId200JSONResponse{
		Id:     &updatedTask.ID,
		Task:   &updatedTask.Task,
		IsDone: &updatedTask.IsDone,
	}

	return response, nil
}
