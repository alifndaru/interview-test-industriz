import { useCategories } from "../hooks/useCategories";
import { Button, Table, Modal, Form, Input } from "antd";
import { useState } from "react";
import { CategoryApi } from "../api/categoryApi";

interface CategoryFormData {
  name: string;
  color?: string;
}

export default function CategoryPage() {
  const { categories } = useCategories();
  const [open, setOpen] = useState(false);

  const [form] = Form.useForm();

  const createCategory = (data: CategoryFormData) => {
    CategoryApi.create({ ...data, color: data.color || '' }).then(() => setOpen(false));
  };

  return (
    <>
      <Button type="primary" onClick={() => setOpen(true)}>
        + Add Category
      </Button>

      <Table
        rowKey="id"
        dataSource={categories}
        columns={[
          { title: "Name", dataIndex: "name" },
          { title: "Color", dataIndex: "color" },
        ]}
      />

      <Modal
        open={open}
        onCancel={() => setOpen(false)}
        onOk={() => form.submit()}
        title="Create Category"
      >
        <Form form={form} layout="vertical" onFinish={createCategory}>
          <Form.Item name="name" label="Name" rules={[{ required: true }]}>
            <Input />
          </Form.Item>
          <Form.Item name="color" label="Color">
            <Input />
          </Form.Item>
        </Form>
      </Modal>
    </>
  );
}
