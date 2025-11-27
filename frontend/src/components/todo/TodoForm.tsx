import { Form, Input, Modal, Select, DatePicker } from "antd";
import type { CreateTodoDto } from "../../types/todo";
import { useCategories } from "../../hooks/useCategories";

export default function TodoForm({
  open,
  onClose,
  onSubmit,
  initialValues,
}: {
  open: boolean;
  onClose: () => void;
  onSubmit: (data: CreateTodoDto) => void;
  initialValues?: Partial<CreateTodoDto>;
}) {
  const [form] = Form.useForm();
  const { categories } = useCategories();

  return (
    <Modal
      open={open}
      title="Todo Form"
      onCancel={onClose}
      onOk={() => form.submit()}
    >
      <Form
        form={form}
        layout="vertical"
        initialValues={initialValues}
        onFinish={onSubmit}
      >
        <Form.Item name="title" label="Title" rules={[{ required: true }]}>
          <Input />
        </Form.Item>

        <Form.Item name="description" label="Description">
          <Input.TextArea rows={3} />
        </Form.Item>

        <Form.Item name="category_id" label="Category">
          <Select>
            {(categories || []).map((c) => (
              <Select.Option key={c.id} value={c.id}>
                {c.name}
              </Select.Option>
            ))}
          </Select>
        </Form.Item>

        <Form.Item name="priority" label="Priority">
          <Select>
            <Select.Option value="high">High</Select.Option>
            <Select.Option value="medium">Medium</Select.Option>
            <Select.Option value="low">Low</Select.Option>
          </Select>
        </Form.Item>

        <Form.Item name="due_date" label="Due Date">
          <DatePicker style={{ width: "100%" }} />
        </Form.Item>
      </Form>
    </Modal>
  );
}
