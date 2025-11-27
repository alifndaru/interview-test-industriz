import { Table, Tag, Button, Switch } from "antd";
import type { Todo } from "../../types/todo";

export default function TodoTable({
  data,
  onEdit,
  onDelete,
  onToggle,
}: {
  data: Todo[];
  onEdit: (t: Todo) => void;
  onDelete: (id: number) => void;
  onToggle: (t: Todo) => void;
}) {
  const columns = [
    {
      title: "Title",
      dataIndex: "title",
    },
    {
      title: "Category",
      render: (_: unknown, row: Todo) => (
        <Tag color={row.category?.color}>{row.category?.name}</Tag>
      ),
    },
    {
      title: "Priority",
      dataIndex: "priority",
    },
    {
      title: "Completed",
      render: (_: unknown, row: Todo) => (
        <Switch
          checked={row.completed}
          onChange={() => onToggle(row)}
          checkedChildren="Done"
          unCheckedChildren="Todo"
        />
      ),
    },
    {
      title: "Actions",
      render: (_: unknown, row: Todo) => (
        <>
          <Button onClick={() => onEdit(row)} type="link">
            Edit
          </Button>
          <Button onClick={() => onDelete(row.id)} danger type="link">
            Delete
          </Button>
        </>
      ),
    },
  ];

  return <Table rowKey="id" columns={columns} dataSource={data} />;
}
