import { useContext, useEffect } from "react";
import { TodoContext } from "../context/TodoContext";
import { TodoApi } from "../api/todoApi";

export const useTodos = (page: number, search: string) => {
  const { todos, setTodos } = useContext(TodoContext);

  useEffect(() => {
    TodoApi.list(page, search).then((res) => {
      setTodos(res.data.data.items || []);
    }).catch((error) => {
      console.error('Failed to fetch todos:', error);
      setTodos([]);
    });
  }, [page, search, setTodos]);

  return { todos };
};
