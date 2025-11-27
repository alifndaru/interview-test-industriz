import { createContext, useState } from "react";
import type { ReactNode } from "react";
import type { Category } from "../types/category";

export const CategoryContext = createContext<{
  categories: Category[];
  setCategories: (categories: Category[]) => void;
}>({
  categories: [],
  setCategories: () => {},
});

export const CategoryProvider = ({ children }: { children: ReactNode }) => {
  const [categories, setCategories] = useState<Category[]>([]);

  return (
    <CategoryContext.Provider value={{ categories, setCategories }}>
      {children}
    </CategoryContext.Provider>
  );
};
