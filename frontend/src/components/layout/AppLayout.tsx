import { Layout, Menu } from "antd";
import { Link } from "react-router-dom";
import { type ReactNode } from "react";

const { Header, Content } = Layout;
export default function AppLayout({ children }: { children: ReactNode }) {
  return (
    <Layout>
      <Header>
        <Menu 
          theme="dark" 
          mode="horizontal"
          items={[
            {
              key: '1',
              label: <Link to="/">Todos</Link>,
            },
            {
              key: '2', 
              label: <Link to="/categories">Categories</Link>,
            },
          ]}
        />
      </Header>

      <Content style={{ padding: "20px" }}>{children}</Content>
    </Layout>
  );
}
