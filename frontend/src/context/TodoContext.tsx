import { createContext, useState } from "react";
import type { ReactNode } from "react";
import type { Todo } from "../types/todo";

interface TodoContextType {
  todos: Todo[];
  setTodos: (t: Todo[]) => void;
}

export const TodoContext = createContext<TodoContextType>({
  todos: [],
  setTodos: () => {},
});

export const TodoProvider = ({ children }: { children: ReactNode }) => {
  const [todos, setTodos] = useState<Todo[]>([]);

  return (
    <TodoContext.Provider value={{ todos, setTodos }}>
      {children}
    </TodoContext.Provider>
  );
};
