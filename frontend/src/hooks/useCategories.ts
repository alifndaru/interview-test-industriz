import { useContext, useEffect } from "react";
import { CategoryApi } from "../api/categoryApi";
import { CategoryContext } from "../context/CategoryContext";

export const useCategories = () => {
  const { categories, setCategories } = useContext(CategoryContext);

  useEffect(() => {
    CategoryApi.list().then((res) => {
      setCategories(res.data.data.categories || []);
    }).catch((error) => {
      console.error('Failed to fetch categories:', error);
      setCategories([]);
    });
  }, [setCategories]);

  return { categories };
};
