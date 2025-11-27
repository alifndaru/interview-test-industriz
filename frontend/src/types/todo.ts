export interface Todo {
  id: number;
  title: string;
  description: string;
  completed: boolean;
  category_id: number;
  priority: string;
  due_date?: string;
  created_at: string;
  updated_at: string;
  category?: {
    id: number;
    name: string;
    color: string;
  };
}

export interface CreateTodoDto {
  title: string;
  description: string;
  category_id: number;
  priority: string;
  due_date?: string;
}

export interface UpdateTodoDto extends CreateTodoDto {
  completed: boolean;
}
