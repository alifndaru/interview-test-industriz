import api from "./axios";
import type { CreateTodoDto, UpdateTodoDto } from "../types/todo";

export const TodoApi = {
  list: (page: number, search: string) =>
    api.get("/todos", { params: { page, search } }),

  get: (id: number) => api.get(`/todos/${id}`),

  create: (data: CreateTodoDto) => api.post("/todos", data),

  update: (id: number, data: UpdateTodoDto) => api.put(`/todos/${id}`, data),

  delete: (id: number) => api.delete(`/todos/${id}`),

  toggle: (id: number, completed: boolean) =>
    api.patch(`/todos/${id}/complete`, { completed }),
};
