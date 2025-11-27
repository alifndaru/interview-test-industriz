import api from "./axios";
import type { Category } from "../types/category";

export const CategoryApi = {
  list: () => api.get("/categories"),
  create: (data: { name: string; color: string }) =>
    api.post("/categories", data),
  update: (id: number, data: Category) => api.put(`/categories/${id}`, data),
  delete: (id: number) => api.delete(`/categories/${id}`),
};
