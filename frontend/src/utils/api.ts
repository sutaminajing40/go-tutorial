import axios from "axios";
import { Todo } from "../types/todo";
import { getApiBaseUrl } from "../config/env";

const API_BASE_URL = getApiBaseUrl();

export const api = axios.create({
  baseURL: API_BASE_URL,
});

export const getTodos = async (): Promise<Todo[]> => {
  try {
    const response = await api.get<Todo[]>("/todos");
    return Array.isArray(response.data) ? response.data : [];
  } catch (error) {
    console.error("Error fetching todos:", error);
    return [];
  }
};

export const createTodo = async (title: string): Promise<Todo> => {
  const response = await api.post<Todo>("/todos", { title });
  return response.data;
};

export const updateTodo = async (
  id: number,
  title: string,
  completed: boolean
): Promise<Todo> => {
  console.log("updateTodo", id, title, completed);
  const response = await api.put<Todo>(`/todos/${id}`, { title, completed });
  return response.data;
};

export const deleteTodo = async (id: number): Promise<void> => {
  await api.delete(`/todos/${id}`);
};
