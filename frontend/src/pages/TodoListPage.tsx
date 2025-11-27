import { useState } from "react";
import { Button, Input } from "antd";

import { useTodos } from "../hooks/useTodos";
import TodoTable from "../components/todo/TodoTable";
import TodoForm from "../components/todo/TodoForm";
import { TodoApi } from "../api/todoApi";

export default function TodoListPage() {
  const [page] = useState(1);
  const [search, setSearch] = useState("");
  const [openForm, setOpenForm] = useState(false);
  const [editData, setEditData] = useState(null);

  const { todos } = useTodos(page, search);

  const handleCreate = (data: any) => {
    TodoApi.create(data).then(() => {
      setOpenForm(false);
    });
  };

  const handleEdit = (data: any) => {
    setEditData(data);
    setOpenForm(true);
  };

  return (
    <>
      <div
        style={{
          display: "flex",
          justifyContent: "space-between",
          marginBottom: 15,
        }}
      >
        <h2>Todo List</h2>
        <Button onClick={() => setOpenForm(true)} type="primary">
          + Add Todo
        </Button>
      </div>

      <div style={{ marginBottom: 15 }}>
        <Input.Search
          placeholder="Search todos..."
          value={search}
          onChange={(e) => setSearch(e.target.value)}
          style={{ maxWidth: 300 }}
        />
      </div>

      <TodoTable
        data={todos}
        onEdit={handleEdit}
        onDelete={(id) => TodoApi.delete(id)}
        onToggle={(t) =>
          TodoApi.toggle(t.id, !t.completed)
        }
      />

      <TodoForm
        open={openForm}
        onClose={() => setOpenForm(false)}
        onSubmit={handleCreate}
        initialValues={editData || undefined}
      />
    </>
  );
}
