import { BrowserRouter, Routes, Route } from "react-router-dom";
import AppLayout from "./components/layout/AppLayout";
import TodoListPage from "./pages/TodoListPage";
import CategoryPage from "./pages/CategoryPage";

import { TodoProvider } from "./context/TodoContext";
import { CategoryProvider } from "./context/CategoryContext";

export default function App() {
  return (
    <BrowserRouter>
      <TodoProvider>
        <CategoryProvider>
          <AppLayout>
            <Routes>
              <Route path="/" element={<TodoListPage />} />
              <Route path="/categories" element={<CategoryPage />} />
            </Routes>
          </AppLayout>
        </CategoryProvider>
      </TodoProvider>
    </BrowserRouter>
  );
}
